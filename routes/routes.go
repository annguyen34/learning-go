package routes

import "github.com/gin-gonic/gin"

func RegisterRoute(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventByID)
	server.POST("/events", postEvent)
}