package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(router *gin.Engine) {
	router.GET("/menu", controller.GetMenus())
	router.GET("/menu/:menu_id", controller.GetMenu())
	router.POST("/menu", controller.CreateMenu())
	router.PUT("/menu/:menu_id", controller.UpdateMenu())
}
