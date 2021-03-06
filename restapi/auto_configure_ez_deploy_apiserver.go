// Code generated by go-swagger; DO NOT EDIT.
// Auto configures api handlers Implementations.

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/ez-deploy/apiserver/handle"
	"github.com/ez-deploy/apiserver/models"
	"github.com/ez-deploy/apiserver/restapi/operations"
	"github.com/ez-deploy/apiserver/restapi/operations/authority_ops"
	"github.com/ez-deploy/apiserver/restapi/operations/identity_ops"
	"github.com/ez-deploy/apiserver/restapi/operations/project_ops"
)

//go:generate swagger generate server --target ../../apiserver --name EzDeployApiserver --spec ../.swagger/swagger.yaml --implementation-package github.com/ez-deploy/apiserver/handle --principal github.com/ez-deploy/apiserver/models.IdentityVerifyResp

// This file auto configures the api backend implementation.
// handle package must already exist.
// handle.New() is implemented by user, and must return an object
// or interface that implements Handler interface defined below.
var Impl Handler = handle.New()

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
	KeyAuth(token string) (*models.IdentityVerifyResp, error)
}

/* AuthorityOpsHandler  */
type AuthorityOpsHandler interface {
	AuthorityOpsDeleteAuthorities(params authority_ops.AuthorityOpsDeleteAuthoritiesParams, principal *models.IdentityVerifyResp) middleware.Responder
	AuthorityOpsListAuthoritiesByIdentity(params authority_ops.AuthorityOpsListAuthoritiesByIdentityParams, principal *models.IdentityVerifyResp) middleware.Responder
	AuthorityOpsListAuthoritiesByResource(params authority_ops.AuthorityOpsListAuthoritiesByResourceParams, principal *models.IdentityVerifyResp) middleware.Responder
	AuthorityOpsSetAuthorities(params authority_ops.AuthorityOpsSetAuthoritiesParams, principal *models.IdentityVerifyResp) middleware.Responder
}

/* IdentityOpsHandler  */
type IdentityOpsHandler interface {
	/* IdentityOpsDeletePublicToken delete public_token. */
	IdentityOpsDeletePublicToken(params identity_ops.IdentityOpsDeletePublicTokenParams, principal *models.IdentityVerifyResp) middleware.Responder
	/* IdentityOpsGeneratePublicToken generate public_token. */
	IdentityOpsGeneratePublicToken(params identity_ops.IdentityOpsGeneratePublicTokenParams, principal *models.IdentityVerifyResp) middleware.Responder
	/* IdentityOpsGetPrivateToken get private_token. */
	IdentityOpsGetPrivateToken(params identity_ops.IdentityOpsGetPrivateTokenParams, principal *models.IdentityVerifyResp) middleware.Responder
	/* IdentityOpsListPublicToken list user's public_tokens. */
	IdentityOpsListPublicToken(params identity_ops.IdentityOpsListPublicTokenParams, principal *models.IdentityVerifyResp) middleware.Responder
	/* IdentityOpsLogin Login by email and password. */
	IdentityOpsLogin(params identity_ops.IdentityOpsLoginParams) middleware.Responder
	/* IdentityOpsReGeneratePrivateToken generate private_token. */
	IdentityOpsReGeneratePrivateToken(params identity_ops.IdentityOpsReGeneratePrivateTokenParams, principal *models.IdentityVerifyResp) middleware.Responder
	/* IdentityOpsRegister Register by email and password. */
	IdentityOpsRegister(params identity_ops.IdentityOpsRegisterParams) middleware.Responder
	/* IdentityOpsVerify Verify by session_token. */
	IdentityOpsVerify(params identity_ops.IdentityOpsVerifyParams, principal *models.IdentityVerifyResp) middleware.Responder
}

/* ProjectOpsHandler  */
type ProjectOpsHandler interface {
	ProjectOpsCreateProject(params project_ops.ProjectOpsCreateProjectParams, principal *models.IdentityVerifyResp) middleware.Responder
	ProjectOpsDeleteProject(params project_ops.ProjectOpsDeleteProjectParams, principal *models.IdentityVerifyResp) middleware.Responder
	ProjectOpsDeleteService(params project_ops.ProjectOpsDeleteServiceParams, principal *models.IdentityVerifyResp) middleware.Responder
	ProjectOpsGetService(params project_ops.ProjectOpsGetServiceParams, principal *models.IdentityVerifyResp) middleware.Responder
	/* ProjectOpsList list all visible projects. */
	ProjectOpsList(params project_ops.ProjectOpsListParams, principal *models.IdentityVerifyResp) middleware.Responder
	ProjectOpsListPods(params project_ops.ProjectOpsListPodsParams, principal *models.IdentityVerifyResp) middleware.Responder
	ProjectOpsListService(params project_ops.ProjectOpsListServiceParams, principal *models.IdentityVerifyResp) middleware.Responder
	ProjectOpsSetService(params project_ops.ProjectOpsSetServiceParams, principal *models.IdentityVerifyResp) middleware.Responder
}

func configureFlags(api *operations.EzDeployApiserverAPI) {
	Impl.ConfigureFlags(api)
}

func configureAPI(api *operations.EzDeployApiserverAPI) http.Handler {

	api.ServeError = errors.ServeError

	api.UseSwaggerUI()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "X-EZDEOPLY-APIKEY" header is set
	api.KeyAuth = func(token string) (*models.IdentityVerifyResp, error) {
		return Impl.KeyAuth(token)
	}

	api.AuthorityOpsAuthorityOpsDeleteAuthoritiesHandler = authority_ops.AuthorityOpsDeleteAuthoritiesHandlerFunc(func(params authority_ops.AuthorityOpsDeleteAuthoritiesParams, principal *models.IdentityVerifyResp) middleware.Responder {
		return Impl.AuthorityOpsDeleteAuthorities(params, principal)
	})
	api.AuthorityOpsAuthorityOpsListAuthoritiesByIdentityHandler = authority_ops.AuthorityOpsListAuthoritiesByIdentityHandlerFunc(func(params authority_ops.AuthorityOpsListAuthoritiesByIdentityParams, principal *models.IdentityVerifyResp) middleware.Responder {
		return Impl.AuthorityOpsListAuthoritiesByIdentity(params, principal)
	})
	api.AuthorityOpsAuthorityOpsListAuthoritiesByResourceHandler = authority_ops.AuthorityOpsListAuthoritiesByResourceHandlerFunc(func(params authority_ops.AuthorityOpsListAuthoritiesByResourceParams, principal *models.IdentityVerifyResp) middleware.Responder {
		return Impl.AuthorityOpsListAuthoritiesByResource(params, principal)
	})
	api.AuthorityOpsAuthorityOpsSetAuthoritiesHandler = authority_ops.AuthorityOpsSetAuthoritiesHandlerFunc(func(params authority_ops.AuthorityOpsSetAuthoritiesParams, principal *models.IdentityVerifyResp) middleware.Responder {
		return Impl.AuthorityOpsSetAuthorities(params, principal)
	})
	api.IdentityOpsIdentityOpsDeletePublicTokenHandler = identity_ops.IdentityOpsDeletePublicTokenHandlerFunc(func(params identity_ops.IdentityOpsDeletePublicTokenParams, principal *models.IdentityVerifyResp) middleware.Responder {
		return Impl.IdentityOpsDeletePublicToken(params, principal)
	})
	api.IdentityOpsIdentityOpsGeneratePublicTokenHandler = identity_ops.IdentityOpsGeneratePublicTokenHandlerFunc(func(params identity_ops.IdentityOpsGeneratePublicTokenParams, principal *models.IdentityVerifyResp) middleware.Responder {
		return Impl.IdentityOpsGeneratePublicToken(params, principal)
	})
	api.IdentityOpsIdentityOpsGetPrivateTokenHandler = identity_ops.IdentityOpsGetPrivateTokenHandlerFunc(func(params identity_ops.IdentityOpsGetPrivateTokenParams, principal *models.IdentityVerifyResp) middleware.Responder {
		return Impl.IdentityOpsGetPrivateToken(params, principal)
	})
	api.IdentityOpsIdentityOpsListPublicTokenHandler = identity_ops.IdentityOpsListPublicTokenHandlerFunc(func(params identity_ops.IdentityOpsListPublicTokenParams, principal *models.IdentityVerifyResp) middleware.Responder {
		return Impl.IdentityOpsListPublicToken(params, principal)
	})
	api.IdentityOpsIdentityOpsLoginHandler = identity_ops.IdentityOpsLoginHandlerFunc(func(params identity_ops.IdentityOpsLoginParams) middleware.Responder {
		return Impl.IdentityOpsLogin(params)
	})
	api.IdentityOpsIdentityOpsReGeneratePrivateTokenHandler = identity_ops.IdentityOpsReGeneratePrivateTokenHandlerFunc(func(params identity_ops.IdentityOpsReGeneratePrivateTokenParams, principal *models.IdentityVerifyResp) middleware.Responder {
		return Impl.IdentityOpsReGeneratePrivateToken(params, principal)
	})
	api.IdentityOpsIdentityOpsRegisterHandler = identity_ops.IdentityOpsRegisterHandlerFunc(func(params identity_ops.IdentityOpsRegisterParams) middleware.Responder {
		return Impl.IdentityOpsRegister(params)
	})
	api.IdentityOpsIdentityOpsVerifyHandler = identity_ops.IdentityOpsVerifyHandlerFunc(func(params identity_ops.IdentityOpsVerifyParams, principal *models.IdentityVerifyResp) middleware.Responder {
		return Impl.IdentityOpsVerify(params, principal)
	})
	api.ProjectOpsProjectOpsCreateProjectHandler = project_ops.ProjectOpsCreateProjectHandlerFunc(func(params project_ops.ProjectOpsCreateProjectParams, principal *models.IdentityVerifyResp) middleware.Responder {
		return Impl.ProjectOpsCreateProject(params, principal)
	})
	api.ProjectOpsProjectOpsDeleteProjectHandler = project_ops.ProjectOpsDeleteProjectHandlerFunc(func(params project_ops.ProjectOpsDeleteProjectParams, principal *models.IdentityVerifyResp) middleware.Responder {
		return Impl.ProjectOpsDeleteProject(params, principal)
	})
	api.ProjectOpsProjectOpsDeleteServiceHandler = project_ops.ProjectOpsDeleteServiceHandlerFunc(func(params project_ops.ProjectOpsDeleteServiceParams, principal *models.IdentityVerifyResp) middleware.Responder {
		return Impl.ProjectOpsDeleteService(params, principal)
	})
	api.ProjectOpsProjectOpsGetServiceHandler = project_ops.ProjectOpsGetServiceHandlerFunc(func(params project_ops.ProjectOpsGetServiceParams, principal *models.IdentityVerifyResp) middleware.Responder {
		return Impl.ProjectOpsGetService(params, principal)
	})
	api.ProjectOpsProjectOpsListHandler = project_ops.ProjectOpsListHandlerFunc(func(params project_ops.ProjectOpsListParams, principal *models.IdentityVerifyResp) middleware.Responder {
		return Impl.ProjectOpsList(params, principal)
	})
	api.ProjectOpsProjectOpsListPodsHandler = project_ops.ProjectOpsListPodsHandlerFunc(func(params project_ops.ProjectOpsListPodsParams, principal *models.IdentityVerifyResp) middleware.Responder {
		return Impl.ProjectOpsListPods(params, principal)
	})
	api.ProjectOpsProjectOpsListServiceHandler = project_ops.ProjectOpsListServiceHandlerFunc(func(params project_ops.ProjectOpsListServiceParams, principal *models.IdentityVerifyResp) middleware.Responder {
		return Impl.ProjectOpsListService(params, principal)
	})
	api.ProjectOpsProjectOpsSetServiceHandler = project_ops.ProjectOpsSetServiceHandlerFunc(func(params project_ops.ProjectOpsSetServiceParams, principal *models.IdentityVerifyResp) middleware.Responder {
		return Impl.ProjectOpsSetService(params, principal)
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	// CustomConfigure can override or add to configurations set above
	Impl.CustomConfigure(api)

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
	Impl.ConfigureTLS(tlsConfig)
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
	Impl.ConfigureServer(s, scheme, addr)
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return Impl.SetupMiddlewares(handler)
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return Impl.SetupGlobalMiddleware(handler)
}
