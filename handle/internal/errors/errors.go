package errors

import (
	"net/http"

	"github.com/ez-deploy/apiserver/models"
	"github.com/ez-deploy/apiserver/restapi/operations/identity_ops"
	"github.com/go-openapi/runtime/middleware"
)

// NewCommonResposeWithErr
func NewCommonResposeWithErr(err error) middleware.Responder {
	commonResp := identity_ops.NewIdentityOpsRegisterOK().WithPayload(&models.ModelCommonResp{
		Error: &models.ModelError{
			Error:   true,
			Message: err.Error(),
		},
	})

	return commonResp
}

// NewInternalServerErrorResponseWithErr
func NewInternalServerErrorResponseWithErr(err error) middleware.Responder {
	return middleware.Error(http.StatusInternalServerError, err.Error())
}
