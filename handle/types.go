package handle

import (
	"crypto/tls"
	"github.com/ez-deploy/apiserver/restapi/operations"
	"github.com/ez-deploy/apiserver/restapi/operations/authority_ops"
	"github.com/ez-deploy/apiserver/restapi/operations/identity_ops"
	"github.com/ez-deploy/apiserver/restapi/operations/project_ops"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
)

// Handler handles all api server backend configurations and requests
type Handler interface {
	Authable
	Configurable
	AuthorityOpsHandler
	IdentityOpsHandler
	ProjectOpsHandler
}

// Configurable handles all server configurations
type Configurable interface {
	ConfigureFlags(api *operations.EzDeployApiserverAPI)
	ConfigureTLS(tlsConfig *tls.Config)
	ConfigureServer(s *http.Server, scheme, addr string)
	CustomConfigure(api *operations.EzDeployApiserverAPI)
	SetupMiddlewares(handler http.Handler) http.Handler
	SetupGlobalMiddleware(handler http.Handler) http.Handler
}

// Authable handles server authentication
type Authable interface {
	// Applies when the "X-EZDEOPLY-APIKEY" header is set
	KeyAuth(token string) (interface{}, error)
}

/* AuthorityOpsHandler  */
type AuthorityOpsHandler interface {
	AuthorityOpsDeleteAuthorities(params authority_ops.AuthorityOpsDeleteAuthoritiesParams, principal interface{}) middleware.Responder
	AuthorityOpsListAuthoritiesByIdentity(params authority_ops.AuthorityOpsListAuthoritiesByIdentityParams, principal interface{}) middleware.Responder
	AuthorityOpsListAuthoritiesByResource(params authority_ops.AuthorityOpsListAuthoritiesByResourceParams, principal interface{}) middleware.Responder
	AuthorityOpsSetAuthorities(params authority_ops.AuthorityOpsSetAuthoritiesParams, principal interface{}) middleware.Responder
}

/* IdentityOpsHandler  */
type IdentityOpsHandler interface {
	/* IdentityOpsDeletePublicToken delete public_token. */
	IdentityOpsDeletePublicToken(params identity_ops.IdentityOpsDeletePublicTokenParams, principal interface{}) middleware.Responder
	/* IdentityOpsGeneratePublicToken generate public_token. */
	IdentityOpsGeneratePublicToken(params identity_ops.IdentityOpsGeneratePublicTokenParams, principal interface{}) middleware.Responder
	/* IdentityOpsGetPrivateToken get private_token. */
	IdentityOpsGetPrivateToken(params identity_ops.IdentityOpsGetPrivateTokenParams, principal interface{}) middleware.Responder
	/* IdentityOpsListPublicToken list user's public_tokens. */
	IdentityOpsListPublicToken(params identity_ops.IdentityOpsListPublicTokenParams, principal interface{}) middleware.Responder
	/* IdentityOpsLogin Login by email and password. */
	IdentityOpsLogin(params identity_ops.IdentityOpsLoginParams) middleware.Responder
	/* IdentityOpsReGeneratePrivateToken generate private_token. */
	IdentityOpsReGeneratePrivateToken(params identity_ops.IdentityOpsReGeneratePrivateTokenParams, principal interface{}) middleware.Responder
	/* IdentityOpsRegister Register by email and password. */
	IdentityOpsRegister(params identity_ops.IdentityOpsRegisterParams) middleware.Responder
	/* IdentityOpsVerify Verify by session_token. */
	IdentityOpsVerify(params identity_ops.IdentityOpsVerifyParams, principal interface{}) middleware.Responder
}

/* ProjectOpsHandler  */
type ProjectOpsHandler interface {
	ProjectOpsCreateProject(params project_ops.ProjectOpsCreateProjectParams, principal interface{}) middleware.Responder
	ProjectOpsDeleteProject(params project_ops.ProjectOpsDeleteProjectParams, principal interface{}) middleware.Responder
	ProjectOpsDeleteService(params project_ops.ProjectOpsDeleteServiceParams, principal interface{}) middleware.Responder
	ProjectOpsGetService(params project_ops.ProjectOpsGetServiceParams, principal interface{}) middleware.Responder
	ProjectOpsListPods(params project_ops.ProjectOpsListPodsParams, principal interface{}) middleware.Responder
	ProjectOpsListService(params project_ops.ProjectOpsListServiceParams, principal interface{}) middleware.Responder
	ProjectOpsSetService(params project_ops.ProjectOpsSetServiceParams, principal interface{}) middleware.Responder
}

func New() Handler {
	return nil
}

