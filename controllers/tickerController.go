package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetTicket(ctx *gin.Context) {
	//params := ctx.Params.ByName("id")
	ctx.JSON(200, gin.H{"user":"bruno", "password": 123})
	
}
