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
		log.Panic("Failed to connect to database: ", err)
	}
	if err := db.Migrator().DropTable(&models.Users{}, &models.Room{}, &models.Booking{}, &models.Notification{}, &models.Office{}); err != nil {
		log.Fatal("Failed to drop table:", err)
	}

	db.AutoMigrate(&models.Users{}, &models.Room{}, &models.Booking{}, &models.Notification{}, &models.Office{})

	var dummyAdmin models.Users
	dummyAdmin.Name = "admin"
	dummyAdmin.PrivyID = "12345"
	dummyAdmin.TypeUser = 1
	dummyAdmin.Email = "admin@test.com"
	dummyAdmin.TeleID = "717992772"
	db.Create(&dummyAdmin)
	var dummyGuest models.Users
	dummyGuest.Name = "guest"
	dummyGuest.PrivyID = "12345"
	dummyGuest.TypeUser = 0
	dummyGuest.Email = "guest@test.com"
	dummyGuest.TeleID = "101010101"
	db.Create(&dummyGuest)

	var dummyRoom models.Room
	dummyRoom.Title = "Room 1"
	dummyRoom.Description = "Room 1 Description"
	dummyRoom.Location = "Room 1 Location"
	dummyRoom.Type = 1
	dummyRoom.Capacity = 5
	dummyRoom.Status = 0
	db.Create(&dummyRoom)
	var dummRoom2 models.Room
	dummRoom2.Title = "Room 2"
	dummRoom2.Description = "Room 2 Description"
	dummRoom2.Location = "Room 2 Location"
	dummRoom2.Type = 2
	dummRoom2.Capacity = 10
	dummRoom2.Status = 0
	db.Create(&dummRoom2)
	var dummyRoom3 models.Room
	dummyRoom3.Title = "Room 3"
	dummyRoom3.Description = "Room 3 Description"
	dummyRoom3.Location = "Room 3 Location"
	dummyRoom3.Type = 3
	dummyRoom3.Capacity = 15
	dummyRoom3.Status = 0
	db.Create(&dummyRoom3)

	DB = db
	return db
}
