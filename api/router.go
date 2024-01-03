package api

import (
	"fmt"
	"makerspace-api/internal/handlers"
	"math/rand"
	"net/http"

	// "makerspace-api/pkg/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Set the templates directory
	r.LoadHTMLGlob("templates/*.html") 

	// Setup cors middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	r.Use(cors.New(config))

	// Auth and authorization middleware placeholder
	// r.Use(middleware.AuthMiddleware())
	r.GET("/",  handlers.LoadHomePage) 

	r.GET("/auth/google/login", handlers.OauthGoogleLogin)
	r.GET("/api/sessions/oauth/google", handlers.OauthGoogleCallback)
	

	r.GET("/random", func (c *gin.Context) {
		randomNumber := rand.Intn(100)
		c.String(http.StatusOK, fmt.Sprintf("Random number: %d", randomNumber))
	})
	r.POST("/order", handlers.AddOrder)
	r.GET("/order/:id", handlers.GetOrderByID)
	r.GET("/orders", handlers.GetOrders)
	return r

}
