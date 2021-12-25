// Code generated by go-swagger; DO NOT EDIT.

package authority_ops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ez-deploy/apiserver/models"
)

// AuthorityOpsListAuthoritiesByResourceOKCode is the HTTP code returned for type AuthorityOpsListAuthoritiesByResourceOK
const AuthorityOpsListAuthoritiesByResourceOKCode int = 200

/*AuthorityOpsListAuthoritiesByResourceOK A successful response.

swagger:response authorityOpsListAuthoritiesByResourceOK
*/
type AuthorityOpsListAuthoritiesByResourceOK struct {

	/*
	  In: Body
	*/
	Payload *models.AuthorityListAuthoritiesResp `json:"body,omitempty"`
}

// NewAuthorityOpsListAuthoritiesByResourceOK creates AuthorityOpsListAuthoritiesByResourceOK with default headers values
func NewAuthorityOpsListAuthoritiesByResourceOK() *AuthorityOpsListAuthoritiesByResourceOK {

	return &AuthorityOpsListAuthoritiesByResourceOK{}
}

// WithPayload adds the payload to the authority ops list authorities by resource o k response
func (o *AuthorityOpsListAuthoritiesByResourceOK) WithPayload(payload *models.AuthorityListAuthoritiesResp) *AuthorityOpsListAuthoritiesByResourceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the authority ops list authorities by resource o k response
func (o *AuthorityOpsListAuthoritiesByResourceOK) SetPayload(payload *models.AuthorityListAuthoritiesResp) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AuthorityOpsListAuthoritiesByResourceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AuthorityOpsListAuthoritiesByResourceDefault An unexpected error response.

swagger:response authorityOpsListAuthoritiesByResourceDefault
*/
type AuthorityOpsListAuthoritiesByResourceDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GatewayruntimeError `json:"body,omitempty"`
}

// NewAuthorityOpsListAuthoritiesByResourceDefault creates AuthorityOpsListAuthoritiesByResourceDefault with default headers values
func NewAuthorityOpsListAuthoritiesByResourceDefault(code int) *AuthorityOpsListAuthoritiesByResourceDefault {
	if code <= 0 {
		code = 500
	}

	return &AuthorityOpsListAuthoritiesByResourceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the authority ops list authorities by resource default response
func (o *AuthorityOpsListAuthoritiesByResourceDefault) WithStatusCode(code int) *AuthorityOpsListAuthoritiesByResourceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the authority ops list authorities by resource default response
func (o *AuthorityOpsListAuthoritiesByResourceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the authority ops list authorities by resource default response
func (o *AuthorityOpsListAuthoritiesByResourceDefault) WithPayload(payload *models.GatewayruntimeError) *AuthorityOpsListAuthoritiesByResourceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the authority ops list authorities by resource default response
func (o *AuthorityOpsListAuthoritiesByResourceDefault) SetPayload(payload *models.GatewayruntimeError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AuthorityOpsListAuthoritiesByResourceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
