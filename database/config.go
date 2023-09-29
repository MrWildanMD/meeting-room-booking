package database

import (
	"fmt"
	"log"

	"github.com/MrWildanMD/room-booking/config"
	"github.com/MrWildanMD/room-booking/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	host string
	DB   *gorm.DB
)

func Init(config *config.Config) *gorm.DB {
	if config.APP_MODE == "docker" {
		host = "postgresdb"
	} else if config.APP_MODE == "local" {
		host = "localhost"
	} else {
		host = config.DBHost
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, config.DBUser, config.DBPassword, config.DBName, config.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalln("Failed to connect to database: ", err)
	}

	db.AutoMigrate(&models.Users{}, &models.Room{}, &models.Booking{}, &models.Notification{}, &models.Office{})

	DB = db
	return db
}
