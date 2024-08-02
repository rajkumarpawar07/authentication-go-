package routes

import (
	"golang-jwt-project/controller"
	"golang-jwt-project/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	// middle ware
	incomingRoutes.Use(middleware.Authenticate())

	// routes
	incomingRoutes.GET("/users", controller.GetUsers())

	incomingRoutes.GET("/users/:user_id", controller.GetUserById())


}