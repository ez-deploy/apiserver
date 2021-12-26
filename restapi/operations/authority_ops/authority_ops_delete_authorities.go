// Code generated by go-swagger; DO NOT EDIT.

package authority_ops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/ez-deploy/apiserver/models"
)

// AuthorityOpsDeleteAuthoritiesHandlerFunc turns a function with the right signature into a authority ops delete authorities handler
type AuthorityOpsDeleteAuthoritiesHandlerFunc func(AuthorityOpsDeleteAuthoritiesParams, *models.IdentityVerifyResp) middleware.Responder

// Handle executing the request and returning a response
func (fn AuthorityOpsDeleteAuthoritiesHandlerFunc) Handle(params AuthorityOpsDeleteAuthoritiesParams, principal *models.IdentityVerifyResp) middleware.Responder {
	return fn(params, principal)
}

// AuthorityOpsDeleteAuthoritiesHandler interface for that can handle valid authority ops delete authorities params
type AuthorityOpsDeleteAuthoritiesHandler interface {
	Handle(AuthorityOpsDeleteAuthoritiesParams, *models.IdentityVerifyResp) middleware.Responder
}

// NewAuthorityOpsDeleteAuthorities creates a new http.Handler for the authority ops delete authorities operation
func NewAuthorityOpsDeleteAuthorities(ctx *middleware.Context, handler AuthorityOpsDeleteAuthoritiesHandler) *AuthorityOpsDeleteAuthorities {
	return &AuthorityOpsDeleteAuthorities{Context: ctx, Handler: handler}
}

/* AuthorityOpsDeleteAuthorities swagger:route DELETE /authority AuthorityOps authorityOpsDeleteAuthorities

AuthorityOpsDeleteAuthorities authority ops delete authorities API

*/
type AuthorityOpsDeleteAuthorities struct {
	Context *middleware.Context
	Handler AuthorityOpsDeleteAuthoritiesHandler
}

func (o *AuthorityOpsDeleteAuthorities) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAuthorityOpsDeleteAuthoritiesParams()
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
