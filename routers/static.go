package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetStaticRoutes(router *gin.Engine) *gin.Engine {

	router.LoadHTMLFiles("templates/index.html")
	router.Static("/static", "./static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	return router
}
