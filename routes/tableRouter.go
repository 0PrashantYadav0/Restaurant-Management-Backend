package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func TableRoutes(router *gin.Engine) {
	router.GET("/tables", controller.GetTables())
	router.GET("/tables/:table_id", controller.GetTable())
	router.POST("/tables", controller.CreateTable())
	router.PUT("/tables/:table_id", controller.UpdateTable())
}
