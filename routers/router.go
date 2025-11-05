package routers

import (
	"rest-api-coffee-server/controllers"

	"github.com/gin-gonic/gin"
)

func CallRouters() {
	router := gin.Default()

	router.GET("/drinks", controllers.GetAllDrinks)
	router.GET("/drinks/in-stock", controllers.GetDrinksInStock)
	router.GET("/drinks/:id", controllers.GetDrinkByID)
	router.POST("/drinks", controllers.CreateDrink)
	router.DELETE("/drinks/:id", controllers.DeleteByID)
	router.PATCH("/drinks/:id", controllers.PatchDrink)

	router.Run()
}
