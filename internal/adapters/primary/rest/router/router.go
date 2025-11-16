package router

import (
	"github.com/gin-gonic/gin"
	"github.com/juancuadra2/crud-api-go/internal/adapters/primary/rest/handlers"
)

func SetupRoutes(userHandler *handlers.UserHandler) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.POST("", userHandler.Create)
			users.GET("", userHandler.GetAll)
			users.GET("/:id", userHandler.GetByID)
			users.PUT("/:id", userHandler.Update)
			users.DELETE("/:id", userHandler.Delete)
		}
	}

	return router
}
