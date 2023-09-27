package main

import (
	"log"

	"github.com/MrWildanMD/room-booking/config"
	"github.com/MrWildanMD/room-booking/database"
	"github.com/MrWildanMD/room-booking/models"
)

func init() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Could not load env", err)
	}
	database.Init(cfg)
}

func main() {
	var dummyAdmin models.Users
	dummyAdmin.Name = "admin"
	dummyAdmin.PrivyID = "1234567"
	dummyAdmin.TypeUser = 1
	dummyAdmin.Email = "admin@test.com"
	dummyAdmin.TeleID = "717992772"
	database.DB.Create(&dummyAdmin)

	// var dummyAdmin models.Users
	// dummyAdmin.Name = "admin"
	// dummyAdmin.PrivyID = "1234567"
	// dummyAdmin.TypeUser = 1
	// dummyAdmin.Email = "admin@test.com"
	// dummyAdmin.TeleID = "717992772"
	// database.DB.Create(&dummyAdmin)
}
