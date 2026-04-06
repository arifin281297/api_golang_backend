package routes

import (
	"go-backend/controllers"

	"github.com/gin-gonic/gin"
)

func UserGroupRoutes(r *gin.RouterGroup) {
	userGroupGroup := r.Group("/groups")
	{
		userGroupGroup.GET("", controllers.GetGroups)
		userGroupGroup.GET("/:id", controllers.GetGroup)
		userGroupGroup.POST("", controllers.CreateGroup)
		userGroupGroup.PUT("/:id", controllers.UpdateGroup)
		userGroupGroup.DELETE("/:id", controllers.DeleteGroup)
	}
}
