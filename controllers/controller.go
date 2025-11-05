package controllers

import (
	"fmt"
	"net/http"
	"rest-api-coffee-server/config"
	"rest-api-coffee-server/models"

	"github.com/gin-gonic/gin"
)

func GetAllDrinks(c *gin.Context) {
	var drinks []models.DrinkShort
	if err := config.DB.Model(&models.Drink{}).Select("id", "name", "price").Find(&drinks).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, drinks)
}

func GetDrinkByID(c *gin.Context) {
	var drink models.Drink
	id := c.Param("id")

	if err := config.DB.First(&drink, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Напиток не найден"})
		return
	}
	c.JSON(200, drink)
}

func GetDrinksInStock(c *gin.Context) {
	var drinks []models.Drink

	if err := config.DB.Where("in_stock = ?", true).Find(&drinks).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, drinks)
}

func CreateDrink(c *gin.Context) {
	var drink models.Drink

	if err := c.ShouldBindJSON(&drink); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := config.DB.Create(&drink).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка при создании"})
		return
	}
	c.JSON(200, drink)
}

func DeleteByID(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Drink{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{"deleted": id})
}

func PatchDrink(c *gin.Context) {
	var drink models.Drink
	id := c.Param("id")
	if err := config.DB.First(&drink, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	var input models.DrinkRequestDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	fmt.Printf("%+v", input)
	config.DB.Model(&drink).Updates(input)

	c.JSON(200, drink)
}
