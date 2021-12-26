package project

import (
	"github.com/ez-deploy/apiserver/models"
	"github.com/ez-deploy/apiserver/restapi/operations/project_ops"
	"github.com/go-openapi/runtime/middleware"
)

func (h *ProjectOpsImpl) ProjectOpsListPods(params project_ops.ProjectOpsListPodsParams, principal *models.IdentityVerifyResp) middleware.Responder {
	return middleware.NotImplemented("not impl")
}
