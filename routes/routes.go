package routes

import (
	"rest-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventByID)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authentication)
	authenticated.POST("/events", postEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
