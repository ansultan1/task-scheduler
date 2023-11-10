package handlers

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	runtime "github.com/ansultan1/task-scheduler"
	domainErr "github.com/ansultan1/task-scheduler/errors"
	"github.com/ansultan1/task-scheduler/gen/models"
	"github.com/ansultan1/task-scheduler/gen/restapi/operations"
)

// NewListTasks handles a request for listing tasks
func NewListTasks(rt *runtime.Runtime) operations.ListTasksHandler {
	return &listTasks{
		rt: rt,
	}
}

type listTasks struct {
	rt *runtime.Runtime
}

// Handle the list tasks request
func (h *listTasks) Handle(params operations.ListTasksParams) middleware.Responder {
	fmt.Println("its good here")
	tasks, err := h.rt.Service().ListTasks()
	if err != nil {
		if apiErr, ok := err.(*domainErr.APIError); ok {
			switch {
			case apiErr.IsError(domainErr.NotFound):

				return operations.NewListTasksNotFound()
			default:

				return operations.NewListTasksInternalServerError()
			}
		}

		// Handle the case where err is not of type *domainErr.APIError
		return operations.NewListTasksInternalServerError()
	}

	var payload []*models.Task
	for _, task := range tasks {
		payload = append(payload, &models.Task{
			Name:          task.Name,
			ID:            task.ID,
			Command:       task.Command,
			ScheduledTime: strfmt.DateTime(task.ScheduledTime),
			Recurring:     task.Recurring,
			TimeZone:      task.TimeZone,
		})
	}

	return operations.NewListTasksOK().WithPayload(payload)
}
