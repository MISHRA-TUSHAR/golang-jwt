package routes

import (
	controller "github.com/MISHRA-TUSHAR/golang-jwt/controllers"
	"github.com/MISHRA-TUSHAR/golang-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	userRoutes := incomingRoutes.Group("/users")
	{
		userRoutes.GET("/", middleware.TokenAuthMiddleware(), controller.GetUsers())
		userRoutes.GET("/:id", middleware.TokenAuthMiddleware(), controller.GetUser())
		userRoutes.PUT("/:id", middleware.TokenAuthMiddleware(), controller.UpdateUser())
		userRoutes.DELETE("/:id", middleware.TokenAuthMiddleware(), controller.DeleteUser())
	}
}
