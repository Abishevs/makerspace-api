package api

import (
	"makerspace-api/internal/handlers"
	// "makerspace-api/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// r.Use(middleware.AuthMiddleware())

	r.POST("/order", handlers.AddOrder)
	r.GET("/order/:id", handlers.GetOrderByID)
	r.GET("/orders", handlers.GetOrders)
	return r
}
