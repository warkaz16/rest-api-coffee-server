package main

import (
	"rest-api-coffee-server/config"

	"github.com/gin-gonic/gin"
)

func main()  {
	config.ConnectDB()

	router := gin.Default()

	

	router.Run()
}