package db

import (
	"github.com/jle02/ShyftLabs-Takehome/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=admin dbname=ShyftLabs port=5432 sslmode=disable TimeZone=America/New_York"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	conn.AutoMigrate(&models.Course{}, &models.Student{}, &models.Result{})

	return conn, err
}
