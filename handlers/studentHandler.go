package handlers

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jle02/ShyftLabs-Takehome/db/repositories"
	swaggerModels "github.com/jle02/ShyftLabs-Takehome/gen/models"
	"github.com/jle02/ShyftLabs-Takehome/gen/restapi/operations"
	studentOperations "github.com/jle02/ShyftLabs-Takehome/gen/restapi/operations/students"
	"gorm.io/gorm"
)

func SetStudentAPIHandlers(api *operations.ShyftlabsIoAPI, db *gorm.DB, studentRepository *repositories.StudentRepository, resultRepository *repositories.ResultRepository) {
	// CREATE STUDENT
	api.StudentsCreateStudentsHandler = studentOperations.CreateStudentsHandlerFunc(
		func(params studentOperations.CreateStudentsParams) middleware.Responder {
			err := studentRepository.CreateStudent(*params.Body.FirstName, *params.Body.FamilyName, *params.Body.DateOfBirth, *params.Body.EmailAddress)
			if err != nil {
				errorMessage := err.Error()
				return studentOperations.NewCreateStudentsDefault(500).WithPayload(&swaggerModels.Error{Code: 500, Message: &errorMessage})
			}
			return studentOperations.NewCreateStudentsCreated()
		})

	//GET STUDENT
	api.StudentsGetStudentsHandler = studentOperations.GetStudentsHandlerFunc(
		func(params studentOperations.GetStudentsParams) middleware.Responder {
			dbStudents, err := studentRepository.GetStudents()
			if err != nil {
				errorMessage := "Unable to get students"
				return studentOperations.NewCreateStudentsDefault(500).WithPayload(&swaggerModels.Error{Code: 500, Message: &errorMessage})
			}

			var swaggerStudents []*swaggerModels.StudentOutput

			for _, dbStudent := range dbStudents {
				swaggerStudent := &swaggerModels.StudentOutput{
					ID:           int64(dbStudent.ID),
					FirstName:    &dbStudent.FirstName,
					FamilyName:   &dbStudent.FamilyName,
					DateOfBirth:  &dbStudent.DateOfBirth,
					EmailAddress: &dbStudent.EmailAddress,
				}
				swaggerStudents = append(swaggerStudents, swaggerStudent)
			}

			return studentOperations.NewGetStudentsOK().WithPayload(swaggerStudents)
		})

	//DELETE STUDENT
	api.StudentsDeleteStudentHandler = studentOperations.DeleteStudentHandlerFunc(
		func(params studentOperations.DeleteStudentParams) middleware.Responder {
			tx := db.Begin()
			err := resultRepository.DeleteResultByStudent(tx, params.ID)
			if err != nil {
				tx.Rollback()
				errorMessage := "Unable to delete student with ID: " + fmt.Sprint(params.ID)
				return studentOperations.NewCreateStudentsDefault(500).WithPayload(&swaggerModels.Error{Code: 500, Message: &errorMessage})
			}
			err = studentRepository.DeleteStudent(tx, params.ID)
			if err != nil {
				tx.Rollback()
				errorMessage := "Unable to delete student with ID: " + fmt.Sprint(params.ID)
				return studentOperations.NewCreateStudentsDefault(500).WithPayload(&swaggerModels.Error{Code: 500, Message: &errorMessage})
			}
			tx.Commit()
			return studentOperations.NewDeleteStudentNoContent()
		})
}
