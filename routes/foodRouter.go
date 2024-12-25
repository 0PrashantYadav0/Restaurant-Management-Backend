package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(router *gin.Engine) {
	router.GET("/food", controller.GetFoods())
	router.GET("/food/:food_id", controller.GetFood())
	router.POST("/food", controller.CreateFood())
	router.PUT("/food/:food_id", controller.UpdateFood())
}
