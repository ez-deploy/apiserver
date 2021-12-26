package configurable

import (
	"crypto/tls"
	"net/http"

	"github.com/ez-deploy/apiserver/restapi/operations"
)

type ConfigurableImpl struct{}

func (c *ConfigurableImpl) ConfigureFlags(api *operations.EzDeployApiserverAPI) {
	// do nothing.
}

func (c *ConfigurableImpl) ConfigureTLS(tlsConfig *tls.Config) {
	// do nothing.
}

func (c *ConfigurableImpl) ConfigureServer(s *http.Server, scheme, addr string) {
	// do nothing.
}

func (c *ConfigurableImpl) CustomConfigure(api *operations.EzDeployApiserverAPI) {
	// do nothing.
}

func (c *ConfigurableImpl) SetupMiddlewares(handler http.Handler) http.Handler {
	// do nothing.

	return handler
}

func (c *ConfigurableImpl) SetupGlobalMiddleware(handler http.Handler) http.Handler {
	// do nothing.

	return handler
}
