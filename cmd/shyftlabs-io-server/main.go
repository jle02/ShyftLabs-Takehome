// Code generated by go-swagger; DO NOT EDIT.

package main

import (
	"log"
	"os"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	"github.com/jle02/ShyftLabs-Takehome/db"
	"github.com/jle02/ShyftLabs-Takehome/db/repositories"
	"github.com/jle02/ShyftLabs-Takehome/gen/restapi"
	"github.com/jle02/ShyftLabs-Takehome/gen/restapi/operations"
	"github.com/jle02/ShyftLabs-Takehome/handlers"

)

// This file was generated by the swagger tool.
// Make sure not to overwrite this file after you generated it because all your edits would be lost!

func main() {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewShyftlabsIoAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	db, err := db.InitializeDB()
	if err != nil {
		panic("failed to connect database")
	}

	studentRepository := repositories.NewStudentRepository(db)
	courseRepository := repositories.NewCourseRepository(db)
	resultRepository := repositories.NewResultRepository(db, *studentRepository, *courseRepository)


	//set api handler funcs here
	handlers.SetAPIHandlers(api, db, studentRepository, courseRepository, resultRepository)

	server.Port = 8080

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Student Result Management System"
	parser.LongDescription = "Student Result Management System"
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
