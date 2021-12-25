// Code generated by go-swagger; DO NOT EDIT.

package authority_ops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"github.com/ez-deploy/apiserver/models"
)

// NewAuthorityOpsSetAuthoritiesParams creates a new AuthorityOpsSetAuthoritiesParams object
//
// There are no default values defined in the spec.
func NewAuthorityOpsSetAuthoritiesParams() AuthorityOpsSetAuthoritiesParams {

	return AuthorityOpsSetAuthoritiesParams{}
}

// AuthorityOpsSetAuthoritiesParams contains all the bound params for the authority ops set authorities operation
// typically these are obtained from a http.Request
//
// swagger:parameters AuthorityOps_SetAuthorities
type AuthorityOpsSetAuthoritiesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	Body *models.AuthorityAuthorities
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAuthorityOpsSetAuthoritiesParams() beforehand.
func (o *AuthorityOpsSetAuthoritiesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.AuthorityAuthorities
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body", ""))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
