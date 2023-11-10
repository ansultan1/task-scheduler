// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/ansultan1/task-scheduler/gen/restapi/operations"
)

//go:generate swagger generate server --target ../../gen --name TaskScheduler --spec ../../swagger.yaml --principal interface{} --exclude-main

func configureFlags(api *operations.TaskSchedulerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TaskSchedulerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.AddTaskHandler == nil {
		api.AddTaskHandler = operations.AddTaskHandlerFunc(func(params operations.AddTaskParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.AddTask has not yet been implemented")
		})
	}
	if api.DeleteTaskHandler == nil {
		api.DeleteTaskHandler = operations.DeleteTaskHandlerFunc(func(params operations.DeleteTaskParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteTask has not yet been implemented")
		})
	}
	if api.GetTaskByIDHandler == nil {
		api.GetTaskByIDHandler = operations.GetTaskByIDHandlerFunc(func(params operations.GetTaskByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetTaskByID has not yet been implemented")
		})
	}
	if api.ListTasksHandler == nil {
		api.ListTasksHandler = operations.ListTasksHandlerFunc(func(params operations.ListTasksParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.ListTasks has not yet been implemented")
		})
	}
	if api.UpdateTaskHandler == nil {
		api.UpdateTaskHandler = operations.UpdateTaskHandlerFunc(func(params operations.UpdateTaskParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.UpdateTask has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
