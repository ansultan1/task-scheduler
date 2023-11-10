package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/ansultan1/task-scheduler"
	domainErr "github.com/ansultan1/task-scheduler/errors"
	"github.com/ansultan1/task-scheduler/gen/restapi/operations"
)

// NewDeleteTask function will delete the task
func NewDeleteTask(rt *runtime.Runtime) operations.DeleteTaskHandler {
	return &deleteTask{
		rt: rt,
	}
}

type deleteTask struct {
	rt *runtime.Runtime
}

// Handle the delete entry request
func (d *deleteTask) Handle(params operations.DeleteTaskParams) middleware.Responder {
	if err := d.rt.Service().DeleteTask(params.ID); err != nil {
		switch apiErr := err.(*domainErr.APIError); {
		case apiErr.IsError(domainErr.NotFound):
			return operations.NewDeleteTaskNotFound()
		default:
			return operations.NewDeleteTaskInternalServerError()
		}
	}
	return operations.NewDeleteTaskNoContent()
}
