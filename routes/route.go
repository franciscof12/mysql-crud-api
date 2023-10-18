package routes

import (
	"github.com/franciscof12/v1/crud_api/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRoutes(r *gin.Engine, db *gorm.DB) {
	userHandler := handlers.UserHandler{DB: db}
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users/:id", userHandler.HandleGetUsersByID)
		v1.POST("/users", userHandler.HandleCreateUser)
		v1.DELETE("/users/:id", userHandler.HandleDeleteUser)
		v1.PUT("/users/:id", userHandler.HandleUpdateUsers)
	}
}
