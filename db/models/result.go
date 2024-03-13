package models

import (
	"gorm.io/gorm"
)

type Result struct {
	gorm.Model
	CourseID  uint
	StudentID uint
	Course    Course  `gorm:"foreignKey:CourseID"`
	Student   Student `gorm:"foreignKey:StudentID"`
	Score     string
}
