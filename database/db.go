package database

import (
	"fmt"
	"log"
	"mygram/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host       = os.Getenv("DB_HOST")
	user       = os.Getenv("DB_USER")
	password   = os.Getenv("DB_PASSWORD")
	dbPort     = os.Getenv("DB_PORT")
	dbname     = os.Getenv("DB_NAME")
	debug_mode = os.Getenv("DEBUG_MODE")
	db         *gorm.DB
	err        error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database :", err)
	}

	fmt.Println("Sukses koneksi ke database")
	if debug_mode == "true" {
		db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
	}

	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})

}

func GetDB() *gorm.DB {
	return db
}
