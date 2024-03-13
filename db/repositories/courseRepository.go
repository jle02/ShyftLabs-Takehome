package repositories

import (
	"github.com/jle02/ShyftLabs-Takehome/db/models"
	"gorm.io/gorm"
)

type CourseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

func (courseRepo *CourseRepository) CreateCourse(courseName string) error {

	course := &models.Course{
		CourseName: courseName,
	}

	results := courseRepo.db.Create(course)

	return results.Error
}

func (courseRepo *CourseRepository) GetCourses() ([]*models.Course, error) {
	var courses []*models.Course
	results := courseRepo.db.Find(&courses)
	return courses, results.Error
}

func (courseRepo *CourseRepository) GetCourse(courseId uint) (*models.Course, error) {
	var course models.Course
	results := courseRepo.db.Find(&course, courseId)
	if results.RowsAffected == 0 {
		return nil, nil
	}
	return &course, results.Error
}

func (courseRepo *CourseRepository) DeleteCourse(tx *gorm.DB, id int64) error {
	return tx.Delete(&models.Course{}, id).Error
}
