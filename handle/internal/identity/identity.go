package identity

import (
	"github.com/ez-deploy/apiserver/handle/internal/errors"
	"github.com/ez-deploy/apiserver/models"
	"github.com/ez-deploy/apiserver/restapi/operations/identity_ops"

	"github.com/ez-deploy/protobuf/convert"
	pb "github.com/ez-deploy/protobuf/identity"
	"github.com/go-openapi/runtime/middleware"
)

type IdentityOpsImpl struct {
	*identityHandlerImpl
	*tokenHandlerImpl
}

type identityHandlerImpl struct {
	cli pb.IdentityOpsClient
}

func NewIdentityOpsImpl(identityClient pb.IdentityOpsClient) *IdentityOpsImpl {
	return &IdentityOpsImpl{
		identityHandlerImpl: &identityHandlerImpl{
			cli: identityClient,
		},
	}
}

/* IdentityOpsLogin Login by email and password. */
func (h *identityHandlerImpl) IdentityOpsLogin(params identity_ops.IdentityOpsLoginParams) middleware.Responder {
	pbReq := &pb.LoginReq{}
	if err := convert.WithJSON(params.Body, pbReq); err != nil {
		return errors.NewInternalServerErrorResponseWithErr(err)
	}

	pbResp, err := h.cli.Login(params.HTTPRequest.Context(), pbReq)
	if err != nil {
		return errors.NewInternalServerErrorResponseWithErr(err)
	}

	httpResp := &models.IdentityLoginResp{}
	if err := convert.WithJSON(pbResp, httpResp); err != nil {
		return errors.NewInternalServerErrorResponseWithErr(err)
	}
	tokenType := models.ModelTokenType(pbResp.Token.Type.String())
	httpResp.Token.Type = &tokenType

	rawToken, err := StringifyToken(httpResp.Token)
	if err != nil {
		return errors.NewInternalServerErrorResponseWithErr(err)
	}
	httpResp.Token.Token = rawToken

	return identity_ops.NewIdentityOpsLoginOK().WithPayload(httpResp)
}

/* IdentityOpsRegister Register by email and password. */
func (h *identityHandlerImpl) IdentityOpsRegister(params identity_ops.IdentityOpsRegisterParams) middleware.Responder {
	pbReq := &pb.RegisterReq{}
	if err := convert.WithJSON(params.Body, pbReq); err != nil {
		return errors.NewInternalServerErrorResponseWithErr(err)
	}

	pbResp, err := h.cli.Register(params.HTTPRequest.Context(), pbReq)
	if err != nil {
		return errors.NewInternalServerErrorResponseWithErr(err)
	}

	httpResp := &models.ModelCommonResp{}
	if err := convert.WithJSON(pbResp, httpResp); err != nil {
		return errors.NewInternalServerErrorResponseWithErr(err)
	}

	return identity_ops.NewIdentityOpsRegisterOK().WithPayload(httpResp)
}

/* IdentityOpsVerify Verify by session_token. */
func (h *identityHandlerImpl) IdentityOpsVerify(params identity_ops.IdentityOpsVerifyParams, principal *models.IdentityVerifyResp) middleware.Responder {
	return identity_ops.NewIdentityOpsVerifyOK().WithPayload(principal)
}
