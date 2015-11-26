package controllers

import (
	"github.com/gin-gonic/gin"
)

func SearchUsers(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"user": "bruno", "password": 123})
}
