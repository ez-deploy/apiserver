// Code generated by go-swagger; DO NOT EDIT.

package identity_ops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewIdentityOpsListPublicTokenParams creates a new IdentityOpsListPublicTokenParams object
//
// There are no default values defined in the spec.
func NewIdentityOpsListPublicTokenParams() IdentityOpsListPublicTokenParams {

	return IdentityOpsListPublicTokenParams{}
}

// IdentityOpsListPublicTokenParams contains all the bound params for the identity ops list public token operation
// typically these are obtained from a http.Request
//
// swagger:parameters IdentityOps_ListPublicToken
type IdentityOpsListPublicTokenParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	Email *string
	/*
	  In: query
	*/
	Name *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewIdentityOpsListPublicTokenParams() beforehand.
func (o *IdentityOpsListPublicTokenParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qEmail, qhkEmail, _ := qs.GetOK("email")
	if err := o.bindEmail(qEmail, qhkEmail, route.Formats); err != nil {
		res = append(res, err)
	}

	qName, qhkName, _ := qs.GetOK("name")
	if err := o.bindName(qName, qhkName, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindEmail binds and validates parameter Email from query.
func (o *IdentityOpsListPublicTokenParams) bindEmail(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Email = &raw

	return nil
}

// bindName binds and validates parameter Name from query.
func (o *IdentityOpsListPublicTokenParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Name = &raw

	return nil
}
