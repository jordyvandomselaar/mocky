package serviceProviders

import (
	"github.com/jinzhu/gorm"
	"github.com/jordyvandomselaar/mock-backend/app/models"
	"log"
)

type Gorm struct {
	Db *gorm.DB
}

func NewGormServiceProvder() *Gorm {
	// Create connection, log any errors
	db, err := gorm.Open("mysql", "root:r6vegas2@/mock_backend?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	// Migrate database
	db.AutoMigrate(&models.User{})

	return &Gorm{
		Db: db,
	}
}
