// Code generated by go-swagger; DO NOT EDIT.

package authority_ops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewAuthorityOpsListAuthoritiesByResourceParams creates a new AuthorityOpsListAuthoritiesByResourceParams object
//
// There are no default values defined in the spec.
func NewAuthorityOpsListAuthoritiesByResourceParams() AuthorityOpsListAuthoritiesByResourceParams {

	return AuthorityOpsListAuthoritiesByResourceParams{}
}

// AuthorityOpsListAuthoritiesByResourceParams contains all the bound params for the authority ops list authorities by resource operation
// typically these are obtained from a http.Request
//
// swagger:parameters AuthorityOps_ListAuthoritiesByResource
type AuthorityOpsListAuthoritiesByResourceParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	Name *string
	/*
	  In: query
	*/
	Type *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAuthorityOpsListAuthoritiesByResourceParams() beforehand.
func (o *AuthorityOpsListAuthoritiesByResourceParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qName, qhkName, _ := qs.GetOK("name")
	if err := o.bindName(qName, qhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	qType, qhkType, _ := qs.GetOK("type")
	if err := o.bindType(qType, qhkType, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindName binds and validates parameter Name from query.
func (o *AuthorityOpsListAuthoritiesByResourceParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

// bindType binds and validates parameter Type from query.
func (o *AuthorityOpsListAuthoritiesByResourceParams) bindType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Type = &raw

	return nil
}
