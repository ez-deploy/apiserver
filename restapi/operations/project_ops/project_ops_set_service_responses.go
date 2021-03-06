// Code generated by go-swagger; DO NOT EDIT.

package project_ops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ez-deploy/apiserver/models"
)

// ProjectOpsSetServiceOKCode is the HTTP code returned for type ProjectOpsSetServiceOK
const ProjectOpsSetServiceOKCode int = 200

/*ProjectOpsSetServiceOK A successful response.

swagger:response projectOpsSetServiceOK
*/
type ProjectOpsSetServiceOK struct {

	/*
	  In: Body
	*/
	Payload *models.ModelCommonResp `json:"body,omitempty"`
}

// NewProjectOpsSetServiceOK creates ProjectOpsSetServiceOK with default headers values
func NewProjectOpsSetServiceOK() *ProjectOpsSetServiceOK {

	return &ProjectOpsSetServiceOK{}
}

// WithPayload adds the payload to the project ops set service o k response
func (o *ProjectOpsSetServiceOK) WithPayload(payload *models.ModelCommonResp) *ProjectOpsSetServiceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the project ops set service o k response
func (o *ProjectOpsSetServiceOK) SetPayload(payload *models.ModelCommonResp) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProjectOpsSetServiceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ProjectOpsSetServiceDefault An unexpected error response.

swagger:response projectOpsSetServiceDefault
*/
type ProjectOpsSetServiceDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GatewayruntimeError `json:"body,omitempty"`
}

// NewProjectOpsSetServiceDefault creates ProjectOpsSetServiceDefault with default headers values
func NewProjectOpsSetServiceDefault(code int) *ProjectOpsSetServiceDefault {
	if code <= 0 {
		code = 500
	}

	return &ProjectOpsSetServiceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the project ops set service default response
func (o *ProjectOpsSetServiceDefault) WithStatusCode(code int) *ProjectOpsSetServiceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the project ops set service default response
func (o *ProjectOpsSetServiceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the project ops set service default response
func (o *ProjectOpsSetServiceDefault) WithPayload(payload *models.GatewayruntimeError) *ProjectOpsSetServiceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the project ops set service default response
func (o *ProjectOpsSetServiceDefault) SetPayload(payload *models.GatewayruntimeError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProjectOpsSetServiceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
