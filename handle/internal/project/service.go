package project

import (
	"github.com/ez-deploy/apiserver/models"
	"github.com/ez-deploy/apiserver/restapi/operations/project_ops"
	"github.com/go-openapi/runtime/middleware"
)

func (h *ProjectOpsImpl) ProjectOpsDeleteService(params project_ops.ProjectOpsDeleteServiceParams, principal *models.IdentityVerifyResp) middleware.Responder {
	return middleware.NotImplemented("not impl")
}
func (h *ProjectOpsImpl) ProjectOpsGetService(params project_ops.ProjectOpsGetServiceParams, principal *models.IdentityVerifyResp) middleware.Responder {
	return middleware.NotImplemented("not impl")
}

func (h *ProjectOpsImpl) ProjectOpsListService(params project_ops.ProjectOpsListServiceParams, principal *models.IdentityVerifyResp) middleware.Responder {
	return middleware.NotImplemented("not impl")
}
func (h *ProjectOpsImpl) ProjectOpsSetService(params project_ops.ProjectOpsSetServiceParams, principal *models.IdentityVerifyResp) middleware.Responder {
	return middleware.NotImplemented("not impl")
}
