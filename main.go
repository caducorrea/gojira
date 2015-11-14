package main

import (
	"github.com/gin-gonic/gin"
	"github.com/brunocodeman/gojira/controllers"
	"net/http"

)

func main() {
	  r := gin.Default()
	  r.LoadHTMLFiles("templates/index.html")
	  r.Static("/static", "./static")
    r.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })
    r.GET("/tickets", controllers.GetTicket)
    r.GET("/",func( c *gin.Context) {
    	 c.HTML(http.StatusOK, "index.html", gin.H{})
    })
    r.Run(":8081") // listen and serve on 0.0.0.0:8080
}