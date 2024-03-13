package handlers

import (
	"github.com/jle02/ShyftLabs-Takehome/db/repositories"
	"github.com/jle02/ShyftLabs-Takehome/gen/restapi/operations"
	"gorm.io/gorm"
)

func SetAPIHandlers(api *operations.ShyftlabsIoAPI, db *gorm.DB, studentRepository *repositories.StudentRepository, courseRepository *repositories.CourseRepository, resultRepository *repositories.ResultRepository) {
	SetStudentAPIHandlers(api, db, studentRepository, resultRepository)
	SetCourseAPIHandlers(api, db, courseRepository, resultRepository)
	SetResultAPIHandlers(api, db, resultRepository)
}
