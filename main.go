package main

import (
	"log"

	"github.com/MrWildanMD/room-booking/config"
	"github.com/MrWildanMD/room-booking/database"
	"github.com/MrWildanMD/room-booking/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Could not load env", err)
	}
	gin.SetMode(cfg.APP_MODE)
	r := gin.Default()
	router := r.Group("api")
	database.Init(cfg)
	routers.RegisterRoutes(router)
	config.RegisterBot(cfg)
	log.Fatal(r.Run(":" + cfg.APP_PORT))
	config.TB.Start()
}
