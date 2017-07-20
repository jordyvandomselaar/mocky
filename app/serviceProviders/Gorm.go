package serviceProviders

import (
	"github.com/jinzhu/gorm"
	"github.com/jordyvandomselaar/mock-backend/app/models"
	"log"
)

// Gorm service provider. Used to initialize Gorm.
type Gorm struct {
	Db *gorm.DB
}

// NewGormServiceProvider returns a new Gorm service provider.
func NewGormServiceProvider() *Gorm {
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
