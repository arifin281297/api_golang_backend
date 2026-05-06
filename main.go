package main

import (
	"go-backend/config"
	"go-backend/controllers"
	"go-backend/middleware"
	"go-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Init Gin
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	r.POST("/login", controllers.LoginHandler)

	// Connect database & migrate
	config.ConnectDatabase()

	// Register routes
	api := r.Group("/api")
	api.Use(middleware.APIKeyAuth())
	routes.UserRoutes(api)
	routes.UserCabangRoutes(api)
	routes.UserGroupRoutes(api)

	// routes.UserRoutes(r)

	// Run server
	r.Run(":8080")
}
