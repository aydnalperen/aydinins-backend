package db

import (
	"aydin-tracker/aydinins-backend/pkg/common/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.Product{})
	return db
}
