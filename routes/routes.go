package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/events", createEvent)
	server.GET("/events", getAllEvents)
	server.POST("/events/registrations/:id", createRegistration)
	server.GET("/events/registrations", getAllRegistrations)
}
