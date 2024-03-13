package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	CourseName string
}
