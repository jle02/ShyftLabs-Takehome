package handlers

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jle02/ShyftLabs-Takehome/db/repositories"
	swaggerModels "github.com/jle02/ShyftLabs-Takehome/gen/models"
	"github.com/jle02/ShyftLabs-Takehome/gen/restapi/operations"
	courseOperations "github.com/jle02/ShyftLabs-Takehome/gen/restapi/operations/courses"
	"gorm.io/gorm"
)

func SetCourseAPIHandlers(api *operations.ShyftlabsIoAPI, db *gorm.DB, courseRepository *repositories.CourseRepository, resultRepository *repositories.ResultRepository) {
	//CREATE COURSE
	api.CoursesCreateCoursesHandler = courseOperations.CreateCoursesHandlerFunc(
		func(params courseOperations.CreateCoursesParams) middleware.Responder {
			err := courseRepository.CreateCourse(*params.Body.CourseName)
			if err != nil {
				errorMessage := err.Error()
				return courseOperations.NewCreateCoursesDefault(500).WithPayload(&swaggerModels.Error{Code: 500, Message: &errorMessage})
			}
			return courseOperations.NewCreateCoursesCreated()
		})

	//GET COURSE
	api.CoursesGetCoursesHandler = courseOperations.GetCoursesHandlerFunc(
		func(params courseOperations.GetCoursesParams) middleware.Responder {
			dbCourses, err := courseRepository.GetCourses()
			if err != nil {
				errorMessage := "Unable to get courses"
				return courseOperations.NewCreateCoursesDefault(500).WithPayload(&swaggerModels.Error{Code: 500, Message: &errorMessage})
			}

			var swaggerCourses []*swaggerModels.CourseOutput

			for _, dbCourse := range dbCourses {
				swaggerCourse := &swaggerModels.CourseOutput{
					ID:         int64(dbCourse.ID),
					CourseName: &dbCourse.CourseName,
				}
				swaggerCourses = append(swaggerCourses, swaggerCourse)
			}

			return courseOperations.NewGetCoursesOK().WithPayload(swaggerCourses)
		})

	//DELETE COURSE
	api.CoursesDeleteCourseHandler = courseOperations.DeleteCourseHandlerFunc(
		func(params courseOperations.DeleteCourseParams) middleware.Responder {
			tx := db.Begin() //initialize transaction so that we can rollback if either deletes fail
			err := courseRepository.DeleteCourse(tx, params.ID)
			if err != nil {
				tx.Rollback()
				errorMessage := "Unable to delete course with ID: " + fmt.Sprint(params.ID)
				return courseOperations.NewCreateCoursesDefault(500).WithPayload(&swaggerModels.Error{Code: 500, Message: &errorMessage})
			}
			err = resultRepository.DeleteResultByCourse(params.ID)
			if err != nil {
				tx.Rollback()
				errorMessage := "Unable to delete course with ID: " + fmt.Sprint(params.ID)
				return courseOperations.NewCreateCoursesDefault(500).WithPayload(&swaggerModels.Error{Code: 500, Message: &errorMessage})
			}
			tx.Commit()
			return courseOperations.NewDeleteCourseNoContent()
		})
}
