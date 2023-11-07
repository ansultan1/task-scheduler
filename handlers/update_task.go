package handlers

import (
	"time"

	"github.com/go-openapi/runtime/middleware"

	runtime "task-scheduler"
	domainErr "task-scheduler/errors"
	"task-scheduler/gen/restapi/operations"
	"task-scheduler/models"
)

// NewUpdateTask handles request for updating task
func NewUpdateTask(rt *runtime.Runtime) operations.UpdateTaskHandler {
	return &updateTask{
		rt: rt,
	}
}

type updateTask struct {
	rt *runtime.Runtime
}

// Handle the update task request
func (d *updateTask) Handle(params operations.UpdateTaskParams) middleware.Responder {
	task := models.Task{
		ID:            params.Task.ID,
		Name:          params.Task.Name,
		Command:       params.Task.Command,
		ScheduledTime: time.Time(params.Task.ScheduledTime),
		Recurring:     params.Task.Recurring,
		TimeZone:      params.Task.TimeZone,
	}
	if _, err := d.rt.Service().AddOrUpdateTask(&task); err != nil {
		switch apiErr := err.(*domainErr.APIError); {
		case apiErr.IsError(domainErr.NotFound):
			return operations.NewUpdateTaskNotFound()
		default:
			return operations.NewUpdateTaskInternalServerError()
		}
	}

	return operations.NewUpdateTaskOK().WithPayload(params.Task)

}
