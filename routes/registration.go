package routes

import (
	"fmt"
	"go-mysql-phpmyadmin/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func createRegistration(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = services.CreateRegistration(eventId)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, fmt.Sprintf("Registration created for event %d", eventId))
}

func getAllRegistrations(context *gin.Context) {
	page, _ := context.GetQuery("page")
	pageSize, _ := context.GetQuery("page_size")

	pageValue, _ := strconv.ParseInt(page, 10, 64)
	pageSizeValue, _ := strconv.ParseInt(pageSize, 10, 64)
	// logic above can become a helper

	registrations, err := services.GetAllRegistrations(int(pageValue), int(pageSizeValue))
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, registrations)
}
