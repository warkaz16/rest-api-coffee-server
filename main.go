package main

import (
	"rest-api-coffee-server/config"
	"rest-api-coffee-server/models"
	"rest-api-coffee-server/routers"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Drink{})
	routers.CallRouters()
}
