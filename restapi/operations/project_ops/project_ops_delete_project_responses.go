// Code generated by go-swagger; DO NOT EDIT.

package project_ops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ez-deploy/apiserver/models"
)

// ProjectOpsDeleteProjectOKCode is the HTTP code returned for type ProjectOpsDeleteProjectOK
const ProjectOpsDeleteProjectOKCode int = 200

/*ProjectOpsDeleteProjectOK A successful response.

swagger:response projectOpsDeleteProjectOK
*/
type ProjectOpsDeleteProjectOK struct {

	/*
	  In: Body
	*/
	Payload *models.ModelCommonResp `json:"body,omitempty"`
}

// NewProjectOpsDeleteProjectOK creates ProjectOpsDeleteProjectOK with default headers values
func NewProjectOpsDeleteProjectOK() *ProjectOpsDeleteProjectOK {

	return &ProjectOpsDeleteProjectOK{}
}

// WithPayload adds the payload to the project ops delete project o k response
func (o *ProjectOpsDeleteProjectOK) WithPayload(payload *models.ModelCommonResp) *ProjectOpsDeleteProjectOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the project ops delete project o k response
func (o *ProjectOpsDeleteProjectOK) SetPayload(payload *models.ModelCommonResp) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProjectOpsDeleteProjectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ProjectOpsDeleteProjectDefault An unexpected error response.

swagger:response projectOpsDeleteProjectDefault
*/
type ProjectOpsDeleteProjectDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GatewayruntimeError `json:"body,omitempty"`
}

// NewProjectOpsDeleteProjectDefault creates ProjectOpsDeleteProjectDefault with default headers values
func NewProjectOpsDeleteProjectDefault(code int) *ProjectOpsDeleteProjectDefault {
	if code <= 0 {
		code = 500
	}

	return &ProjectOpsDeleteProjectDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the project ops delete project default response
func (o *ProjectOpsDeleteProjectDefault) WithStatusCode(code int) *ProjectOpsDeleteProjectDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the project ops delete project default response
func (o *ProjectOpsDeleteProjectDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the project ops delete project default response
func (o *ProjectOpsDeleteProjectDefault) WithPayload(payload *models.GatewayruntimeError) *ProjectOpsDeleteProjectDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the project ops delete project default response
func (o *ProjectOpsDeleteProjectDefault) SetPayload(payload *models.GatewayruntimeError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProjectOpsDeleteProjectDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
