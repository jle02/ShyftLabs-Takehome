package repositories

import (
	"errors"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/jle02/ShyftLabs-Takehome/db/models"
	"gorm.io/gorm"
)

type StudentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{db: db}
}

func (studentRepo *StudentRepository) CreateStudent(firstName string, familyName string, dateOfBirth strfmt.Date, emailAddress strfmt.Email) error {

	dobTime := time.Time(dateOfBirth)
	currentTime := time.Now()

	yearsDifference := currentTime.Year() - dobTime.Year()

	if yearsDifference < 10 {
		return errors.New("student is less than 10 years old")
	}

	student := &models.Student{
		FirstName:    firstName,
		FamilyName:   familyName,
		DateOfBirth:  dateOfBirth,
		EmailAddress: emailAddress,
	}

	results := studentRepo.db.Create(student)

	return results.Error
}

func (studentRepo *StudentRepository) GetStudents() ([]*models.Student, error) {
	var students []*models.Student
	results := studentRepo.db.Find(&students)
	return students, results.Error
}

func (studentRepo *StudentRepository) GetStudent(studentId uint) (*models.Student, error) {
	var student models.Student
	results := studentRepo.db.Find(&student, studentId)
	if results.RowsAffected == 0 {
		return nil, nil
	}
	return &student, results.Error
}

func (studentRepo *StudentRepository) DeleteStudent(tx *gorm.DB, id int64) error {
	return tx.Delete(&models.Student{}, id).Error
}
