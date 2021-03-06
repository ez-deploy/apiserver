// Code generated by go-swagger; DO NOT EDIT.

package identity_ops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/ez-deploy/apiserver/models"
)

// IdentityOpsDeletePublicTokenHandlerFunc turns a function with the right signature into a identity ops delete public token handler
type IdentityOpsDeletePublicTokenHandlerFunc func(IdentityOpsDeletePublicTokenParams, *models.IdentityVerifyResp) middleware.Responder

// Handle executing the request and returning a response
func (fn IdentityOpsDeletePublicTokenHandlerFunc) Handle(params IdentityOpsDeletePublicTokenParams, principal *models.IdentityVerifyResp) middleware.Responder {
	return fn(params, principal)
}

// IdentityOpsDeletePublicTokenHandler interface for that can handle valid identity ops delete public token params
type IdentityOpsDeletePublicTokenHandler interface {
	Handle(IdentityOpsDeletePublicTokenParams, *models.IdentityVerifyResp) middleware.Responder
}

// NewIdentityOpsDeletePublicToken creates a new http.Handler for the identity ops delete public token operation
func NewIdentityOpsDeletePublicToken(ctx *middleware.Context, handler IdentityOpsDeletePublicTokenHandler) *IdentityOpsDeletePublicToken {
	return &IdentityOpsDeletePublicToken{Context: ctx, Handler: handler}
}

/* IdentityOpsDeletePublicToken swagger:route DELETE /user/token/public IdentityOps identityOpsDeletePublicToken

delete public_token.

*/
type IdentityOpsDeletePublicToken struct {
	Context *middleware.Context
	Handler IdentityOpsDeletePublicTokenHandler
}

func (o *IdentityOpsDeletePublicToken) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewIdentityOpsDeletePublicTokenParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.IdentityVerifyResp
	if uprinc != nil {
		principal = uprinc.(*models.IdentityVerifyResp) // this is really a models.IdentityVerifyResp, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
