package identity

import (
	"github.com/ez-deploy/apiserver/models"
	"github.com/ez-deploy/apiserver/restapi/operations/identity_ops"
	"github.com/go-openapi/runtime/middleware"
)

type tokenHandlerImpl struct {
}

/* IdentityOpsDeletePublicToken delete public_token. */
func (h *tokenHandlerImpl) IdentityOpsDeletePublicToken(params identity_ops.IdentityOpsDeletePublicTokenParams, principal *models.IdentityVerifyResp) middleware.Responder {

	return middleware.NotImplemented("not impl")
}

/* IdentityOpsGeneratePublicToken generate public_token. */
func (h *tokenHandlerImpl) IdentityOpsGeneratePublicToken(params identity_ops.IdentityOpsGeneratePublicTokenParams, principal *models.IdentityVerifyResp) middleware.Responder {
	return middleware.NotImplemented("not impl")
}

/* IdentityOpsGetPrivateToken get private_token. */
func (h *tokenHandlerImpl) IdentityOpsGetPrivateToken(params identity_ops.IdentityOpsGetPrivateTokenParams, principal *models.IdentityVerifyResp) middleware.Responder {
	return middleware.NotImplemented("not impl")
}

/* IdentityOpsListPublicToken list user's public_tokens. */
func (h *tokenHandlerImpl) IdentityOpsListPublicToken(params identity_ops.IdentityOpsListPublicTokenParams, principal *models.IdentityVerifyResp) middleware.Responder {
	return middleware.NotImplemented("not impl")
}

/* IdentityOpsReGeneratePrivateToken generate private_token. */
func (h *tokenHandlerImpl) IdentityOpsReGeneratePrivateToken(params identity_ops.IdentityOpsReGeneratePrivateTokenParams, principal *models.IdentityVerifyResp) middleware.Responder {
	return middleware.NotImplemented("not impl")
}
