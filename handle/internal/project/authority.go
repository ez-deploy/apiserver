package project

import (
	"net/http"

	"github.com/ez-deploy/apiserver/models"
	"github.com/go-openapi/runtime/middleware"
)

func (h *ProjectOpsImpl) haveAuthorityToProject(principal *models.IdentityVerifyResp, projectName string) (bool, error) {
	// TODO: impl and use authoritypb.Verify method.

	return true, nil
}

func newForbiddenRespFromMsg(msg string) middleware.Responder {
	return middleware.Error(http.StatusForbidden, msg)
}
