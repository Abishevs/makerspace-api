package api

import (
	"makerspace-api/internal/handlers"
	// "makerspace-api/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

    // Set the templates directory
    r.LoadHTMLGlob("templates/*") 

	// Auth and authorization middleware placeholder
	// r.Use(middleware.AuthMiddleware())
	r.GET("/",  handlers.LoadHomePage) 

	r.GET("/auth/google/login", handlers.OauthGoogleLogin)
	r.GET("/api/sessions/oauth/google", handlers.OauthGoogleCallback)

	r.POST("/order", handlers.AddOrder)
	r.GET("/order/:id", handlers.GetOrderByID)
	r.GET("/orders", handlers.GetOrders)
	return r

}
