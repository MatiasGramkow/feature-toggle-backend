package database

import (
	"github.com/backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB OBJECT
var DB *gorm.DB

// Connect to mySQL database
func Connect() {
	connection, err := gorm.Open(mysql.Open("root:addwishdev@/feature-toggle-db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{}, &models.PasswordReset{}, &models.Feature{})
}
