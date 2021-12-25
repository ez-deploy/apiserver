// Code generated by go-swagger; DO NOT EDIT.

package project_ops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ProjectOpsCreateProjectHandlerFunc turns a function with the right signature into a project ops create project handler
type ProjectOpsCreateProjectHandlerFunc func(ProjectOpsCreateProjectParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ProjectOpsCreateProjectHandlerFunc) Handle(params ProjectOpsCreateProjectParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ProjectOpsCreateProjectHandler interface for that can handle valid project ops create project params
type ProjectOpsCreateProjectHandler interface {
	Handle(ProjectOpsCreateProjectParams, interface{}) middleware.Responder
}

// NewProjectOpsCreateProject creates a new http.Handler for the project ops create project operation
func NewProjectOpsCreateProject(ctx *middleware.Context, handler ProjectOpsCreateProjectHandler) *ProjectOpsCreateProject {
	return &ProjectOpsCreateProject{Context: ctx, Handler: handler}
}

/* ProjectOpsCreateProject swagger:route POST /project ProjectOps projectOpsCreateProject

ProjectOpsCreateProject project ops create project API

*/
type ProjectOpsCreateProject struct {
	Context *middleware.Context
	Handler ProjectOpsCreateProjectHandler
}

func (o *ProjectOpsCreateProject) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewProjectOpsCreateProjectParams()
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