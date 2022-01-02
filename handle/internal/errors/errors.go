package errors

import (
	"net/http"

	"github.com/ez-deploy/apiserver/models"
	"github.com/ez-deploy/apiserver/restapi/operations/identity_ops"
	"github.com/ez-deploy/protobuf/model"
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

func NewCommonResponseWithErrBody(body *model.Error) middleware.Responder {
	if body == nil {
		return identity_ops.NewIdentityOpsRegisterOK().WithPayload(&models.ModelCommonResp{})
	}

	commonResp := identity_ops.NewIdentityOpsRegisterOK().WithPayload(&models.ModelCommonResp{
		Error: &models.ModelError{
			Error:   body.Error,
			Message: body.Message,
		},
	})

	return commonResp
}
