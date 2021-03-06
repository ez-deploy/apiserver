// Code generated by go-swagger; DO NOT EDIT.

package identity_ops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/ez-deploy/apiserver/models"
)

// IdentityOpsGeneratePublicTokenHandlerFunc turns a function with the right signature into a identity ops generate public token handler
type IdentityOpsGeneratePublicTokenHandlerFunc func(IdentityOpsGeneratePublicTokenParams, *models.IdentityVerifyResp) middleware.Responder

// Handle executing the request and returning a response
func (fn IdentityOpsGeneratePublicTokenHandlerFunc) Handle(params IdentityOpsGeneratePublicTokenParams, principal *models.IdentityVerifyResp) middleware.Responder {
	return fn(params, principal)
}

// IdentityOpsGeneratePublicTokenHandler interface for that can handle valid identity ops generate public token params
type IdentityOpsGeneratePublicTokenHandler interface {
	Handle(IdentityOpsGeneratePublicTokenParams, *models.IdentityVerifyResp) middleware.Responder
}

// NewIdentityOpsGeneratePublicToken creates a new http.Handler for the identity ops generate public token operation
func NewIdentityOpsGeneratePublicToken(ctx *middleware.Context, handler IdentityOpsGeneratePublicTokenHandler) *IdentityOpsGeneratePublicToken {
	return &IdentityOpsGeneratePublicToken{Context: ctx, Handler: handler}
}

/* IdentityOpsGeneratePublicToken swagger:route POST /user/token/public IdentityOps identityOpsGeneratePublicToken

generate public_token.

*/
type IdentityOpsGeneratePublicToken struct {
	Context *middleware.Context
	Handler IdentityOpsGeneratePublicTokenHandler
}

func (o *IdentityOpsGeneratePublicToken) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewIdentityOpsGeneratePublicTokenParams()
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
