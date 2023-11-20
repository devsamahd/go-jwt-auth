package routes

import (
	"github.com/devsamahd/go-jwt-api/controllers"
	"github.com/devsamahd/go-jwt-api/middleware"
	"github.com/gin-gonic/gin" 
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET(".users", controllers.GetUsers())
	incomingRoutes.GET(".users/user_id", controllers.GetSingleUser())
}