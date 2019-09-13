/*
 * Copyright (C) 2019 The "MysteriumNetwork/feedback" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package feedback

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"github.com/mysteriumnetwork/feedback/infra"
)

// Endpoint user feedback endpoint
type Endpoint struct {
	reporter    *Reporter
	storage     *Storage
	rateLimiter *infra.RateLimiter
}

// NewEndpoint creates new Endpoint
func NewEndpoint(reporter *Reporter, storage *Storage, rateLimiter *infra.RateLimiter) *Endpoint {
	return &Endpoint{reporter: reporter, storage: storage, rateLimiter: rateLimiter}
}

// CreateGithubIssueRequest create github issue request
// swagger:parameters createGithubIssue
type CreateGithubIssueRequest struct {
	// in: formData
	// required: true
	UserId string `json:"userId"`
	// in: formData
	// required: true
	Description string `json:"description"`
	// in: formData
	Email string `json:"email"`
	// in: formData
	// required: true
	// swagger:file
	File *multipart.FileHeader `json:"file"`
}

// CreateGithubIssueResponse represents a successful github issue creation
// swagger:model
type CreateGithubIssueResponse struct {
	IssueId string `json:"issueId"`
}

// ParseGithubIssueRequest parses CreateGithubIssueRequest from HTTP request
func ParseGithubIssueRequest(c *gin.Context) (form CreateGithubIssueRequest, errors []error) {
	var ok bool
	form.UserId, ok = c.GetPostForm("userId")
	if !ok {
		errors = append(errors, infra.Required("userId"))
	}

	form.Description, ok = c.GetPostForm("description")
	if !ok {
		errors = append(errors, infra.Required("description"))
	}

	form.Email = c.PostForm("email")

	var err error
	form.File, err = c.FormFile("file")
	if err != nil {
		errors = append(errors, infra.Required("file"))
	}

	return form, errors
}

// CreateGithubIssue creates a new Github issue with user report
//
// swagger:operation POST /github createGithubIssue
// ---
// summary: Creates a new Github issue with user report
// description: 1 request per minute is allowed
//
// produces:
// - application/json
// consumes:
// - multipart/form-data
// responses:
//   '200':
//     description: Issue created in Github
//     schema:
//       "$ref": "#/definitions/CreateGithubIssueResponse"
//   '400':
//     description: Bad request
//     schema:
//       "$ref": "#/definitions/ErrorResponse"
//   '429':
//     description: Too many requests
//   '500':
//     description: Internal server error
//     schema:
//       "$ref": "#/definitions/ErrorResponse"
//
func (e *Endpoint) CreateGithubIssue(c *gin.Context) {
	form, requestErrs := ParseGithubIssueRequest(c)
	if len(requestErrs) > 0 {
		c.JSON(http.StatusBadRequest, infra.Multiple(requestErrs))
		return
	}

	tempFile, err := ioutil.TempFile("", form.File.Filename)
	if err != nil {
		err := fmt.Errorf("could not allocate a temporary file: %w", err)
		c.JSON(http.StatusInternalServerError, infra.Single(err))
		return
	}
	defer func() {
		err := os.Remove(tempFile.Name())
		if err != nil {
			_ = log.Warn("Failed to remove temp file", err)
		}
	}()

	err = c.SaveUploadedFile(form.File, tempFile.Name())
	if err != nil {
		err := fmt.Errorf("could not save the uploaded file: %w", err)
		c.JSON(http.StatusInternalServerError, infra.Single(err))
		return
	}

	logURL, err := e.storage.Upload(tempFile.Name())
	if err != nil {
		err := fmt.Errorf("could not upload file to the storage: %w", err)
		c.JSON(http.StatusServiceUnavailable, infra.Single(err))
		return
	}

	issueId, err := e.reporter.ReportIssue(&Report{
		UserId:      form.UserId,
		Description: form.Description,
		Email:       form.Email,
		LogURL:      *logURL,
	})
	if err != nil {
		err := fmt.Errorf("could not report issue to github: %w", err)
		c.JSON(http.StatusServiceUnavailable, infra.Single(err))
		return
	}

	log.Infof("Created github issue %q from request %+v", issueId, form)
	c.JSON(http.StatusOK, &CreateGithubIssueResponse{
		IssueId: issueId,
	})
}

// RegisterRoutes registers feedback API routes
func (e *Endpoint) RegisterRoutes(r gin.IRoutes) {
	r.POST("/github", e.rateLimiter.Handler(), e.CreateGithubIssue)
}
