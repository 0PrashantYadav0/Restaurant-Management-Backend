package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(router *gin.Engine) {
	router.GET("/orderItems", controller.GetOrderItems())
	router.GET("/orderItems/:orderItem_id", controller.GetOrderItem())
	router.GET("/orderItems/order/:order_id", controller.GetOrderItemsByOrder())
	router.POST("/orderItems", controller.CreateOrderItem())
	router.PUT("/orderItems/:orderItem_id", controller.UpdateOrderItem())
}
