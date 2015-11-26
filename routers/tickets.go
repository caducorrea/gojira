package routers

import (
	"github.com/gin-gonic/gin"
	"gojira/controllers"
)

func SetTicketsRoutes(router *gin.Engine) *gin.Engine {

	router.GET("/tickets", controllers.GetTicket)

	return router
}
