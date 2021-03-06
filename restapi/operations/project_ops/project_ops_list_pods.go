// Code generated by go-swagger; DO NOT EDIT.

package project_ops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/ez-deploy/apiserver/models"
)

// ProjectOpsListPodsHandlerFunc turns a function with the right signature into a project ops list pods handler
type ProjectOpsListPodsHandlerFunc func(ProjectOpsListPodsParams, *models.IdentityVerifyResp) middleware.Responder

// Handle executing the request and returning a response
func (fn ProjectOpsListPodsHandlerFunc) Handle(params ProjectOpsListPodsParams, principal *models.IdentityVerifyResp) middleware.Responder {
	return fn(params, principal)
}

// ProjectOpsListPodsHandler interface for that can handle valid project ops list pods params
type ProjectOpsListPodsHandler interface {
	Handle(ProjectOpsListPodsParams, *models.IdentityVerifyResp) middleware.Responder
}

// NewProjectOpsListPods creates a new http.Handler for the project ops list pods operation
func NewProjectOpsListPods(ctx *middleware.Context, handler ProjectOpsListPodsHandler) *ProjectOpsListPods {
	return &ProjectOpsListPods{Context: ctx, Handler: handler}
}

/* ProjectOpsListPods swagger:route GET /project/service/pod ProjectOps projectOpsListPods

ProjectOpsListPods project ops list pods API

*/
type ProjectOpsListPods struct {
	Context *middleware.Context
	Handler ProjectOpsListPodsHandler
}

func (o *ProjectOpsListPods) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewProjectOpsListPodsParams()
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
