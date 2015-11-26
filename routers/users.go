package routers

import (
	"github.com/gin-gonic/gin"
	"gojira/controllers"
)

func SetUsersRoutes(router *gin.Engine) *gin.Engine {

	router.GET("/users", controllers.SearchUsers)

	return router
}
