package main

import (
	"log"
	"simp-service/pkg/api"
	"simp-service/pkg/config"
)

func main() {
	props := config.Properties()

	db := config.ConnectDB(props)
	handler := api.Handler{DB: db}
	router := api.SetupRouter(&handler)

	err := router.Run(":" + props.Server.Port)
	if err != nil {
		log.Fatal("Service startup failed", err)
	}
}
