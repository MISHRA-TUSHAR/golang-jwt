package routes

import (
	controller "github.com/MISHRA-TUSHAR/golang-jwt/contollers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/login", contoller.Login())
	incomingRoutes.POST("users/register", contoller.Signup())
}
