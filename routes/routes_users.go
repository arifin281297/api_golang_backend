package routes

import (
	"go-backend/controllers"

	"github.com/gin-gonic/gin"
)

// func UserRoutes(r *gin.Engine) {
func UserRoutes(r *gin.RouterGroup) {
	userGroup := r.Group("/users")
	{
		userGroup.GET("", controllers.GetUsers)
		userGroup.GET("/:id", controllers.GetUser)
		userGroup.POST("", controllers.CreateUser)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
	}
}
