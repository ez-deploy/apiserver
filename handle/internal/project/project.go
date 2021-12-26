package project

import (
	"github.com/ez-deploy/apiserver/models"
	"github.com/ez-deploy/apiserver/restapi/operations/project_ops"
	"github.com/go-openapi/runtime/middleware"
)

type ProjectOpsImpl struct{}

func (h *ProjectOpsImpl) ProjectOpsCreateProject(params project_ops.ProjectOpsCreateProjectParams, principal *models.IdentityVerifyResp) middleware.Responder {
	return middleware.NotImplemented("not impl")
}

func (h *ProjectOpsImpl) ProjectOpsDeleteProject(params project_ops.ProjectOpsDeleteProjectParams, principal *models.IdentityVerifyResp) middleware.Responder {
	return middleware.NotImplemented("not impl")
}

func (h *ProjectOpsImpl) ProjectOpsList(params project_ops.ProjectOpsListParams, principal *models.IdentityVerifyResp) middleware.Responder {
	return middleware.NotImplemented("not impl")
}
