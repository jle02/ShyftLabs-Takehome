package repositories

import (
	"errors"

	"github.com/jle02/ShyftLabs-Takehome/db/models"
	"gorm.io/gorm"
)

type ResultRepository struct {
	studentRepo StudentRepository
	courseRepo  CourseRepository
	db          *gorm.DB
}

func NewResultRepository(db *gorm.DB, studentRepo StudentRepository, courseRepo CourseRepository) *ResultRepository {
	return &ResultRepository{db: db, studentRepo: studentRepo, courseRepo: courseRepo}
}

func (resultRepo *ResultRepository) CreateResult(courseId uint, studentId uint, score string) error {

	student, err := resultRepo.studentRepo.GetStudent(studentId)
	if err != nil {
		return errors.New("unable to retrieve student for result")
	}
	if student == nil {
		return errors.New("student does not exist")
	}
	course, err := resultRepo.courseRepo.GetCourse(courseId)
	if err != nil {
		return errors.New("unable to retrieve course for result")
	}
	if course == nil {
		return errors.New("course does not exist")
	}

	result := &models.Result{
		CourseID:  uint(courseId),
		Course:    *course,
		Student:   *student,
		StudentID: uint(studentId),
		Score:     score,
	}

	results := resultRepo.db.Create(result)

	return results.Error
}

func (resultRepo *ResultRepository) GetResults() ([]*models.Result, error) {
	var output []*models.Result
	results := resultRepo.db.Preload("Course").Preload("Student").Find(&output)
	return output, results.Error
}

func (resultRepo *ResultRepository) DeleteResultByCourse(courseID int64) error {
	return resultRepo.db.Where("course_id = ?", courseID).Delete(&models.Result{}).Error
}

func (resultRepo *ResultRepository) DeleteResultByStudent(tx *gorm.DB, studentID int64) error {
	return tx.Where("student_id = ?", studentID).Delete(&models.Result{}).Error
}
