// Code generated by go-swagger; DO NOT EDIT.

package identity_ops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ez-deploy/apiserver/models"
)

// IdentityOpsListPublicTokenOKCode is the HTTP code returned for type IdentityOpsListPublicTokenOK
const IdentityOpsListPublicTokenOKCode int = 200

/*IdentityOpsListPublicTokenOK A successful response.

swagger:response identityOpsListPublicTokenOK
*/
type IdentityOpsListPublicTokenOK struct {

	/*
	  In: Body
	*/
	Payload *models.IdentityListPublicTokenResp `json:"body,omitempty"`
}

// NewIdentityOpsListPublicTokenOK creates IdentityOpsListPublicTokenOK with default headers values
func NewIdentityOpsListPublicTokenOK() *IdentityOpsListPublicTokenOK {

	return &IdentityOpsListPublicTokenOK{}
}

// WithPayload adds the payload to the identity ops list public token o k response
func (o *IdentityOpsListPublicTokenOK) WithPayload(payload *models.IdentityListPublicTokenResp) *IdentityOpsListPublicTokenOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the identity ops list public token o k response
func (o *IdentityOpsListPublicTokenOK) SetPayload(payload *models.IdentityListPublicTokenResp) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IdentityOpsListPublicTokenOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*IdentityOpsListPublicTokenDefault An unexpected error response.

swagger:response identityOpsListPublicTokenDefault
*/
type IdentityOpsListPublicTokenDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GatewayruntimeError `json:"body,omitempty"`
}

// NewIdentityOpsListPublicTokenDefault creates IdentityOpsListPublicTokenDefault with default headers values
func NewIdentityOpsListPublicTokenDefault(code int) *IdentityOpsListPublicTokenDefault {
	if code <= 0 {
		code = 500
	}

	return &IdentityOpsListPublicTokenDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the identity ops list public token default response
func (o *IdentityOpsListPublicTokenDefault) WithStatusCode(code int) *IdentityOpsListPublicTokenDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the identity ops list public token default response
func (o *IdentityOpsListPublicTokenDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the identity ops list public token default response
func (o *IdentityOpsListPublicTokenDefault) WithPayload(payload *models.GatewayruntimeError) *IdentityOpsListPublicTokenDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the identity ops list public token default response
func (o *IdentityOpsListPublicTokenDefault) SetPayload(payload *models.GatewayruntimeError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IdentityOpsListPublicTokenDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}