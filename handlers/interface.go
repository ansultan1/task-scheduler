package handlers

import (
	"github.com/go-openapi/loads"

	runtime "github.com/ansultan1/task-scheduler"
	"github.com/ansultan1/task-scheduler/gen/restapi/operations"
)

// Handler replaces swagger handler
type Handler *operations.TaskSchedulerAPI

// NewHandler overrides swagger api handlers
func NewHandler(rt *runtime.Runtime, spec *loads.Document) Handler {
	handler := operations.NewTaskSchedulerAPI(spec)

	// task handlers
	handler.AddTaskHandler = NewAddTask(rt)
	handler.GetTaskByIDHandler = NewGetTask(rt)
	handler.ListTasksHandler = NewListTasks(rt)
	handler.DeleteTaskHandler = NewDeleteTask(rt)
	handler.UpdateTaskHandler = NewUpdateTask(rt)

	return handler
}
