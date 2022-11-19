package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/GabrielLuizSF/cockroach-db/server/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		Photos := main.Group("photos")
		{
			Photos.GET("/", controllers.ShowAllPhotos)
			Photos.GET("/:id", controllers.ShowPhoto)
			Photos.POST("/", controllers.CreatePhoto)
			Photos.PUT("/", controllers.EditPhoto)
			Photos.DELETE("/:id", controllers.DeletePhoto)
		}
	}

	return router
}