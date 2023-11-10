package handlers

import (
	"time"

	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/ansultan1/task-scheduler"
	"github.com/ansultan1/task-scheduler/gen/restapi/operations"
	"github.com/ansultan1/task-scheduler/models"
)

// NewAddTask handles request for saving task
func NewAddTask(rt *runtime.Runtime) operations.AddTaskHandler {
	return &addTask{rt: rt}
}

type addTask struct {
	rt *runtime.Runtime
}

// Handle the add task request
func (d *addTask) Handle(params operations.AddTaskParams) middleware.Responder {
	task := models.Task{
		ID:            params.Task.ID,
		Name:          params.Task.Name,
		Command:       params.Task.Command,
		ScheduledTime: time.Time(params.Task.ScheduledTime),
		Recurring:     params.Task.Recurring,
		TimeZone:      params.Task.TimeZone,
	}
	id, err := d.rt.Service().AddOrUpdateTask(&task)
	if err != nil {
		log().Errorf("failed to add task: %s", err)
		return operations.NewAddTaskConflict()
	}

	params.Task.ID = id
	return operations.NewAddTaskCreated().WithPayload(params.Task)
}
