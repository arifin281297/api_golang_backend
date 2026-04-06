package routes

import (
	"go-backend/controllers"

	"github.com/gin-gonic/gin"
)

func UserCabangRoutes(r *gin.RouterGroup) {
	cabangGroup := r.Group("/cabangs")
	{
		cabangGroup.GET("", controllers.GetCabangs)
		cabangGroup.GET("/:id", controllers.GetCabang)
		cabangGroup.POST("", controllers.CreateCabang)
		cabangGroup.PUT("/:id", controllers.UpdateCabang)
		cabangGroup.DELETE("/:id", controllers.DeleteCabang)
	}
}
