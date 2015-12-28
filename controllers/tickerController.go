package controllers

import (
	"github.com/gin-gonic/gin"
	"gojira/models"
	"time"
	"fmt"
)

func GetTicket(ctx *gin.Context) {
	search_term := ctx.Query("search_term")
	tickets := []models.Ticket{models.Ticket{ID: 99, CreatedAt: time.Now(), Title:"First Ticket", Department: "IT", Description: "This is the first ticket"}}	
	fmt.Println("started at " + time.Now().String())
	for i := 0; i < 1000; i++ {
		ticket := models.Ticket{ID:i, Title:"Other Ticket", CreatedAt: time.Now(), Department: "IT", Description: search_term}
		tickets = append(tickets, ticket)
		
	}
	fmt.Println("finished at " + time.Now().String())
	
	ctx.JSON(200, gin.H{"tickets":tickets})
	
}
