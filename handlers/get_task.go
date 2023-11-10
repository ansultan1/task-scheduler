package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	runtime "github.com/ansultan1/task-scheduler"
	domainErr "github.com/ansultan1/task-scheduler/errors"
	"github.com/ansultan1/task-scheduler/gen/models"
	"github.com/ansultan1/task-scheduler/gen/restapi/operations"
)

// NewGetTask handles a request for retrieving task
func NewGetTask(rt *runtime.Runtime) operations.GetTaskByIDHandler {
	return &getTask{rt: rt}
}

type getTask struct {
	rt *runtime.Runtime
}

// Handle the get task request
func (d *getTask) Handle(params operations.GetTaskByIDParams) middleware.Responder {
	task, err := d.rt.Service().GetTask(params.ID)
	if err != nil {
		switch apiErr := err.(*domainErr.APIError); {
		case apiErr.IsError(domainErr.NotFound):
			return operations.NewGetTaskByIDNotFound()
		default:
			return operations.NewAddTaskInternalServerError()
		}
	}

	return operations.NewGetTaskByIDOK().WithPayload(&models.Task{
		Name:          task.Name,
		ID:            task.ID,
		Command:       task.Command,
		ScheduledTime: strfmt.DateTime(task.ScheduledTime),
		Recurring:     task.Recurring,
		TimeZone:      task.TimeZone,
	})
}
