package database

import (
	"log"
	"os"

	"github.com/superduperkevin/fiber-rest/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to db \n", err.Error())
		os.Exit(2)
	}
	log.Println("Connected to the db")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations")
	// TODO: Add migrations
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = DbInstance{Db: db}
}