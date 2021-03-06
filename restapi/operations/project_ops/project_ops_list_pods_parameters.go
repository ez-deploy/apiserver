// Code generated by go-swagger; DO NOT EDIT.

package project_ops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewProjectOpsListPodsParams creates a new ProjectOpsListPodsParams object
//
// There are no default values defined in the spec.
func NewProjectOpsListPodsParams() ProjectOpsListPodsParams {

	return ProjectOpsListPodsParams{}
}

// ProjectOpsListPodsParams contains all the bound params for the project ops list pods operation
// typically these are obtained from a http.Request
//
// swagger:parameters ProjectOps_ListPods
type ProjectOpsListPodsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	ProjectName *string
	/*
	  In: query
	*/
	ServiceName *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewProjectOpsListPodsParams() beforehand.
func (o *ProjectOpsListPodsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qProjectName, qhkProjectName, _ := qs.GetOK("project_name")
	if err := o.bindProjectName(qProjectName, qhkProjectName, route.Formats); err != nil {
		res = append(res, err)
	}

	qServiceName, qhkServiceName, _ := qs.GetOK("service_name")
	if err := o.bindServiceName(qServiceName, qhkServiceName, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindProjectName binds and validates parameter ProjectName from query.
func (o *ProjectOpsListPodsParams) bindProjectName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.ProjectName = &raw

	return nil
}

// bindServiceName binds and validates parameter ServiceName from query.
func (o *ProjectOpsListPodsParams) bindServiceName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.ServiceName = &raw

	return nil
}
