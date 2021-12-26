package authority

import (
	"github.com/ez-deploy/apiserver/models"
	"github.com/ez-deploy/apiserver/restapi/operations/authority_ops"
	"github.com/go-openapi/runtime/middleware"
)

type AuthorityOpsImpl struct {
}

func (h *AuthorityOpsImpl) AuthorityOpsDeleteAuthorities(params authority_ops.AuthorityOpsDeleteAuthoritiesParams, principal *models.IdentityVerifyResp) middleware.Responder {
	return middleware.NotImplemented("not impl")
}

func (h *AuthorityOpsImpl) AuthorityOpsListAuthoritiesByIdentity(params authority_ops.AuthorityOpsListAuthoritiesByIdentityParams, principal *models.IdentityVerifyResp) middleware.Responder {
	return middleware.NotImplemented("not impl")
}

func (h *AuthorityOpsImpl) AuthorityOpsListAuthoritiesByResource(params authority_ops.AuthorityOpsListAuthoritiesByResourceParams, principal *models.IdentityVerifyResp) middleware.Responder {
	return middleware.NotImplemented("not impl")
}

func (h *AuthorityOpsImpl) AuthorityOpsSetAuthorities(params authority_ops.AuthorityOpsSetAuthoritiesParams, principal *models.IdentityVerifyResp) middleware.Responder {
	return middleware.NotImplemented("not impl")
}
