// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"task-scheduler/gen/models"
)

// GetTaskByIDOKCode is the HTTP code returned for type GetTaskByIDOK
const GetTaskByIDOKCode int = 200

/*
GetTaskByIDOK task response

swagger:response getTaskByIdOK
*/
type GetTaskByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Task `json:"body,omitempty"`
}

// NewGetTaskByIDOK creates GetTaskByIDOK with default headers values
func NewGetTaskByIDOK() *GetTaskByIDOK {

	return &GetTaskByIDOK{}
}

// WithPayload adds the payload to the get task by Id o k response
func (o *GetTaskByIDOK) WithPayload(payload *models.Task) *GetTaskByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get task by Id o k response
func (o *GetTaskByIDOK) SetPayload(payload *models.Task) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTaskByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTaskByIDNotFoundCode is the HTTP code returned for type GetTaskByIDNotFound
const GetTaskByIDNotFoundCode int = 404

/*
GetTaskByIDNotFound task not found

swagger:response getTaskByIdNotFound
*/
type GetTaskByIDNotFound struct {
}

// NewGetTaskByIDNotFound creates GetTaskByIDNotFound with default headers values
func NewGetTaskByIDNotFound() *GetTaskByIDNotFound {

	return &GetTaskByIDNotFound{}
}

// WriteResponse to the client
func (o *GetTaskByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetTaskByIDInternalServerErrorCode is the HTTP code returned for type GetTaskByIDInternalServerError
const GetTaskByIDInternalServerErrorCode int = 500

/*
GetTaskByIDInternalServerError internal server error

swagger:response getTaskByIdInternalServerError
*/
type GetTaskByIDInternalServerError struct {
}

// NewGetTaskByIDInternalServerError creates GetTaskByIDInternalServerError with default headers values
func NewGetTaskByIDInternalServerError() *GetTaskByIDInternalServerError {

	return &GetTaskByIDInternalServerError{}
}

// WriteResponse to the client
func (o *GetTaskByIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
