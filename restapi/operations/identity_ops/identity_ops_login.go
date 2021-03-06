// Code generated by go-swagger; DO NOT EDIT.

package identity_ops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// IdentityOpsLoginHandlerFunc turns a function with the right signature into a identity ops login handler
type IdentityOpsLoginHandlerFunc func(IdentityOpsLoginParams) middleware.Responder

// Handle executing the request and returning a response
func (fn IdentityOpsLoginHandlerFunc) Handle(params IdentityOpsLoginParams) middleware.Responder {
	return fn(params)
}

// IdentityOpsLoginHandler interface for that can handle valid identity ops login params
type IdentityOpsLoginHandler interface {
	Handle(IdentityOpsLoginParams) middleware.Responder
}

// NewIdentityOpsLogin creates a new http.Handler for the identity ops login operation
func NewIdentityOpsLogin(ctx *middleware.Context, handler IdentityOpsLoginHandler) *IdentityOpsLogin {
	return &IdentityOpsLogin{Context: ctx, Handler: handler}
}

/* IdentityOpsLogin swagger:route POST /user/login IdentityOps identityOpsLogin

Login by email and password.

*/
type IdentityOpsLogin struct {
	Context *middleware.Context
	Handler IdentityOpsLoginHandler
}

func (o *IdentityOpsLogin) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewIdentityOpsLoginParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
