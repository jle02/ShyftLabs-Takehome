package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/jle02/ShyftLabs-Takehome/db/repositories"
	swaggerModels "github.com/jle02/ShyftLabs-Takehome/gen/models"
	"github.com/jle02/ShyftLabs-Takehome/gen/restapi/operations"
	resultOperations "github.com/jle02/ShyftLabs-Takehome/gen/restapi/operations/results"
	"gorm.io/gorm"
)

func SetResultAPIHandlers(api *operations.ShyftlabsIoAPI, db *gorm.DB, resultRepository *repositories.ResultRepository) {
	//CREATE RESULT
	api.ResultsCreateResultsHandler = resultOperations.CreateResultsHandlerFunc(
		func(params resultOperations.CreateResultsParams) middleware.Responder {
			err := resultRepository.CreateResult(uint(params.CourseID), uint(params.StudentID), string(*params.Body.Score))
			if err != nil {
				errorMessage := err.Error()
				return resultOperations.NewCreateResultsDefault(500).WithPayload(&swaggerModels.Error{Code: 500, Message: &errorMessage})
			}
			return resultOperations.NewCreateResultsCreated()
		})

	//GET RESULTS
	api.ResultsGetResultsHandler = resultOperations.GetResultsHandlerFunc(
		func(params resultOperations.GetResultsParams) middleware.Responder {
			dbResults, err := resultRepository.GetResults()
			if err != nil {
				errorMessage := "Unable to get results"
				return resultOperations.NewCreateResultsDefault(500).WithPayload(&swaggerModels.Error{Code: 500, Message: &errorMessage})
			}

			var swaggerResults []*swaggerModels.ResultOutput

			for _, dbResult := range dbResults {
				studentName := dbResult.Student.FirstName + " " + dbResult.Student.FamilyName
				swaggerResult := &swaggerModels.ResultOutput{
					ID:          int64(dbResult.ID),
					CourseID:    int64(dbResult.CourseID),
					CourseName:  &dbResult.Course.CourseName,
					StudentID:   int64(dbResult.StudentID),
					StudentName: &studentName,
				}
				swaggerResults = append(swaggerResults, swaggerResult)
			}

			return resultOperations.NewGetResultsOK().WithPayload(swaggerResults)
		})
	api.ResultsGetResultsScoreHandler = resultOperations.GetResultsScoreHandlerFunc(
		func(params resultOperations.GetResultsScoreParams) middleware.Responder {
			scores := []swaggerModels.ScoreEnum{
				swaggerModels.ScoreEnumA,
				swaggerModels.ScoreEnumB,
				swaggerModels.ScoreEnumC,
				swaggerModels.ScoreEnumD,
				swaggerModels.ScoreEnumE,
			}
			return resultOperations.NewGetResultsScoreOK().WithPayload(scores)
		})
}
