package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {

	router := gin.Default()

	router = SetStaticRoutes(router)
	router = SetUsersRoutes(router)
	router = SetTicketsRoutes(router)

	return router
}
