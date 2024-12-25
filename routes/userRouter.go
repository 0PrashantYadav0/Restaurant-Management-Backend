package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/users", controller.GetUsers())
	router.GET("/users/:user_id", controller.GetUser())
	router.POST("/users/signup", controller.SignUp())
	router.POST("/users/login", controller.Login())
}
