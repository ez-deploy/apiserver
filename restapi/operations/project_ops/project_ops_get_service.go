// Code generated by go-swagger; DO NOT EDIT.

package project_ops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ProjectOpsGetServiceHandlerFunc turns a function with the right signature into a project ops get service handler
type ProjectOpsGetServiceHandlerFunc func(ProjectOpsGetServiceParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ProjectOpsGetServiceHandlerFunc) Handle(params ProjectOpsGetServiceParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ProjectOpsGetServiceHandler interface for that can handle valid project ops get service params
type ProjectOpsGetServiceHandler interface {
	Handle(ProjectOpsGetServiceParams, interface{}) middleware.Responder
}

// NewProjectOpsGetService creates a new http.Handler for the project ops get service operation
func NewProjectOpsGetService(ctx *middleware.Context, handler ProjectOpsGetServiceHandler) *ProjectOpsGetService {
	return &ProjectOpsGetService{Context: ctx, Handler: handler}
}

/* ProjectOpsGetService swagger:route GET /project/service/get ProjectOps projectOpsGetService

ProjectOpsGetService project ops get service API

*/
type ProjectOpsGetService struct {
	Context *middleware.Context
	Handler ProjectOpsGetServiceHandler
}

func (o *ProjectOpsGetService) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewProjectOpsGetServiceParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}