// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"github.com/ansultan1/task-scheduler/gen/models"
)

// NewUpdateTaskParams creates a new UpdateTaskParams object
//
// There are no default values defined in the spec.
func NewUpdateTaskParams() UpdateTaskParams {

	return UpdateTaskParams{}
}

// UpdateTaskParams contains all the bound params for the update task operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateTask
type UpdateTaskParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*task details
	  Required: true
	  In: body
	*/
	Task *models.Task
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateTaskParams() beforehand.
func (o *UpdateTaskParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Task
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("task", "body", ""))
			} else {
				res = append(res, errors.NewParseError("task", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Task = &body
			}
		}
	} else {
		res = append(res, errors.Required("task", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
