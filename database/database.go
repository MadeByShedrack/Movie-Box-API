package database

import (
	"github.com/MadeByShedrack/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	myDatabase, err := gorm.Open(sqlite.Open("movies.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to DB")
	}

	err = myDatabase.AutoMigrate(&models.Movies{})

	if err != nil {
		return
	}

	DB = myDatabase
}
