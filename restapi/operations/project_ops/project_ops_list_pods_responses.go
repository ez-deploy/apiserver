// Code generated by go-swagger; DO NOT EDIT.

package project_ops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ez-deploy/apiserver/models"
)

// ProjectOpsListPodsOKCode is the HTTP code returned for type ProjectOpsListPodsOK
const ProjectOpsListPodsOKCode int = 200

/*ProjectOpsListPodsOK A successful response.

swagger:response projectOpsListPodsOK
*/
type ProjectOpsListPodsOK struct {

	/*
	  In: Body
	*/
	Payload *models.ProjectListPodsResp `json:"body,omitempty"`
}

// NewProjectOpsListPodsOK creates ProjectOpsListPodsOK with default headers values
func NewProjectOpsListPodsOK() *ProjectOpsListPodsOK {

	return &ProjectOpsListPodsOK{}
}

// WithPayload adds the payload to the project ops list pods o k response
func (o *ProjectOpsListPodsOK) WithPayload(payload *models.ProjectListPodsResp) *ProjectOpsListPodsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the project ops list pods o k response
func (o *ProjectOpsListPodsOK) SetPayload(payload *models.ProjectListPodsResp) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProjectOpsListPodsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ProjectOpsListPodsDefault An unexpected error response.

swagger:response projectOpsListPodsDefault
*/
type ProjectOpsListPodsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GatewayruntimeError `json:"body,omitempty"`
}

// NewProjectOpsListPodsDefault creates ProjectOpsListPodsDefault with default headers values
func NewProjectOpsListPodsDefault(code int) *ProjectOpsListPodsDefault {
	if code <= 0 {
		code = 500
	}

	return &ProjectOpsListPodsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the project ops list pods default response
func (o *ProjectOpsListPodsDefault) WithStatusCode(code int) *ProjectOpsListPodsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the project ops list pods default response
func (o *ProjectOpsListPodsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the project ops list pods default response
func (o *ProjectOpsListPodsDefault) WithPayload(payload *models.GatewayruntimeError) *ProjectOpsListPodsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the project ops list pods default response
func (o *ProjectOpsListPodsDefault) SetPayload(payload *models.GatewayruntimeError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProjectOpsListPodsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
