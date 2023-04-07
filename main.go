package main

import (
	"log"
	"simp-service/pkg/api"
	"simp-service/pkg/config"
)

func main() {
	props := config.Properties()
	db := config.ConnectDB(props)

	router := api.SetupRouter(db)

	err := router.Run(":" + props.Server.Port)
	if err != nil {
		log.Fatal("Service startup failed", err)
	}
}
