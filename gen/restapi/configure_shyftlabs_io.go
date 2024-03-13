// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/jle02/ShyftLabs-Takehome/gen/restapi/operations"
	"github.com/jle02/ShyftLabs-Takehome/gen/restapi/operations/courses"
	"github.com/jle02/ShyftLabs-Takehome/gen/restapi/operations/results"
	"github.com/jle02/ShyftLabs-Takehome/gen/restapi/operations/students"
)

//go:generate swagger generate server --target ..\..\gen --name ShyftlabsIo --spec ..\..\swagger.yml --principal interface{} --exclude-main

func configureFlags(api *operations.ShyftlabsIoAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ShyftlabsIoAPI) http.Handler {
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

	if api.StudentsCreateStudentsHandler == nil {
		api.StudentsCreateStudentsHandler = students.CreateStudentsHandlerFunc(func(params students.CreateStudentsParams) middleware.Responder {
			return middleware.NotImplemented("operation students.CreateStudents has not yet been implemented")
		})
	}
	if api.StudentsDeleteStudentHandler == nil {
		api.StudentsDeleteStudentHandler = students.DeleteStudentHandlerFunc(func(params students.DeleteStudentParams) middleware.Responder {
			return middleware.NotImplemented("operation students.DeleteStudent has not yet been implemented")
		})
	}
	if api.CoursesGetCoursesHandler == nil {
		api.CoursesGetCoursesHandler = courses.GetCoursesHandlerFunc(func(params courses.GetCoursesParams) middleware.Responder {
			return middleware.NotImplemented("operation courses.GetCourses has not yet been implemented")
		})
	}
	if api.ResultsGetResultsHandler == nil {
		api.ResultsGetResultsHandler = results.GetResultsHandlerFunc(func(params results.GetResultsParams) middleware.Responder {
			return middleware.NotImplemented("operation results.GetResults has not yet been implemented")
		})
	}
	if api.StudentsGetStudentsHandler == nil {
		api.StudentsGetStudentsHandler = students.GetStudentsHandlerFunc(func(params students.GetStudentsParams) middleware.Responder {
			return middleware.NotImplemented("operation students.GetStudents has not yet been implemented")
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
