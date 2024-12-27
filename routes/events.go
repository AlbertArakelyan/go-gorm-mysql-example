package routes

import (
	"go-mysql-phpmyadmin/models"
	"go-mysql-phpmyadmin/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = services.CreateEvent(&event)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, event)
}

func getAllEvents(context *gin.Context) {
	page, _ := context.GetQuery("page")
	pageSize, _ := context.GetQuery("page_size")

	pageValue, _ := strconv.ParseInt(page, 10, 64)
	pageSizeValue, _ := strconv.ParseInt(pageSize, 10, 64)
	// logic above can become a helper

	events, err := services.GetAllEvents(int(pageValue), int(pageSizeValue))
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// withPagination(events, pageValue, pageSizeValue) // where events is type of "any"

	context.JSON(http.StatusOK, events)
}
