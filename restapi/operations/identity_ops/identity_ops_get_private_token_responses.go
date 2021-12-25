// Code generated by go-swagger; DO NOT EDIT.

package identity_ops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ez-deploy/apiserver/models"
)

// IdentityOpsGetPrivateTokenOKCode is the HTTP code returned for type IdentityOpsGetPrivateTokenOK
const IdentityOpsGetPrivateTokenOKCode int = 200

/*IdentityOpsGetPrivateTokenOK A successful response.

swagger:response identityOpsGetPrivateTokenOK
*/
type IdentityOpsGetPrivateTokenOK struct {

	/*
	  In: Body
	*/
	Payload *models.ModelToken `json:"body,omitempty"`
}

// NewIdentityOpsGetPrivateTokenOK creates IdentityOpsGetPrivateTokenOK with default headers values
func NewIdentityOpsGetPrivateTokenOK() *IdentityOpsGetPrivateTokenOK {

	return &IdentityOpsGetPrivateTokenOK{}
}

// WithPayload adds the payload to the identity ops get private token o k response
func (o *IdentityOpsGetPrivateTokenOK) WithPayload(payload *models.ModelToken) *IdentityOpsGetPrivateTokenOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the identity ops get private token o k response
func (o *IdentityOpsGetPrivateTokenOK) SetPayload(payload *models.ModelToken) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IdentityOpsGetPrivateTokenOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*IdentityOpsGetPrivateTokenDefault An unexpected error response.

swagger:response identityOpsGetPrivateTokenDefault
*/
type IdentityOpsGetPrivateTokenDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GatewayruntimeError `json:"body,omitempty"`
}

// NewIdentityOpsGetPrivateTokenDefault creates IdentityOpsGetPrivateTokenDefault with default headers values
func NewIdentityOpsGetPrivateTokenDefault(code int) *IdentityOpsGetPrivateTokenDefault {
	if code <= 0 {
		code = 500
	}

	return &IdentityOpsGetPrivateTokenDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the identity ops get private token default response
func (o *IdentityOpsGetPrivateTokenDefault) WithStatusCode(code int) *IdentityOpsGetPrivateTokenDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the identity ops get private token default response
func (o *IdentityOpsGetPrivateTokenDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the identity ops get private token default response
func (o *IdentityOpsGetPrivateTokenDefault) WithPayload(payload *models.GatewayruntimeError) *IdentityOpsGetPrivateTokenDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the identity ops get private token default response
func (o *IdentityOpsGetPrivateTokenDefault) SetPayload(payload *models.GatewayruntimeError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IdentityOpsGetPrivateTokenDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
