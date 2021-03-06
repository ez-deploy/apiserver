// Code generated by go-swagger; DO NOT EDIT.

package project_ops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ez-deploy/apiserver/models"
)

// ProjectOpsDeleteServiceOKCode is the HTTP code returned for type ProjectOpsDeleteServiceOK
const ProjectOpsDeleteServiceOKCode int = 200

/*ProjectOpsDeleteServiceOK A successful response.

swagger:response projectOpsDeleteServiceOK
*/
type ProjectOpsDeleteServiceOK struct {

	/*
	  In: Body
	*/
	Payload *models.ModelCommonResp `json:"body,omitempty"`
}

// NewProjectOpsDeleteServiceOK creates ProjectOpsDeleteServiceOK with default headers values
func NewProjectOpsDeleteServiceOK() *ProjectOpsDeleteServiceOK {

	return &ProjectOpsDeleteServiceOK{}
}

// WithPayload adds the payload to the project ops delete service o k response
func (o *ProjectOpsDeleteServiceOK) WithPayload(payload *models.ModelCommonResp) *ProjectOpsDeleteServiceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the project ops delete service o k response
func (o *ProjectOpsDeleteServiceOK) SetPayload(payload *models.ModelCommonResp) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProjectOpsDeleteServiceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ProjectOpsDeleteServiceDefault An unexpected error response.

swagger:response projectOpsDeleteServiceDefault
*/
type ProjectOpsDeleteServiceDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GatewayruntimeError `json:"body,omitempty"`
}

// NewProjectOpsDeleteServiceDefault creates ProjectOpsDeleteServiceDefault with default headers values
func NewProjectOpsDeleteServiceDefault(code int) *ProjectOpsDeleteServiceDefault {
	if code <= 0 {
		code = 500
	}

	return &ProjectOpsDeleteServiceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the project ops delete service default response
func (o *ProjectOpsDeleteServiceDefault) WithStatusCode(code int) *ProjectOpsDeleteServiceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the project ops delete service default response
func (o *ProjectOpsDeleteServiceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the project ops delete service default response
func (o *ProjectOpsDeleteServiceDefault) WithPayload(payload *models.GatewayruntimeError) *ProjectOpsDeleteServiceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the project ops delete service default response
func (o *ProjectOpsDeleteServiceDefault) SetPayload(payload *models.GatewayruntimeError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProjectOpsDeleteServiceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
