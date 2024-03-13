package models

import (
	"github.com/go-openapi/strfmt"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	DateOfBirth  strfmt.Date
	EmailAddress strfmt.Email
	FamilyName   string
	FirstName    string
}
