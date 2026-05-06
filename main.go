package main

import (
	"go-backend/controllers"
	"go-backend/middleware"
	"go-backend/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Init Gin
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	r.POST("/login", controllers.LoginHandler)

	// Connect database & migrate
	// config.ConnectDatabase()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API hidup 🚀",
		})
	})

	// Register routes
	api := r.Group("/api")
	api.Use(middleware.APIKeyAuth())
	routes.UserRoutes(api)
	routes.UserCabangRoutes(api)
	routes.UserGroupRoutes(api)

	// routes.UserRoutes(r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Run server
	r.Run(":" + port)
}
