package main

import (
	"net/http"
	"rest-api/db"
	"rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	r := gin.Default()
	r.GET("/events", getEvents)
	r.GET("/events/:id", getEventByID)
	r.POST("/events", postEvent)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func postEvent(c *gin.Context) {
	var event models.Event

	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	res, err := models.GetAllEvent()
	event.UserID = len(res) + 1

	models.Save(event)

	c.JSON(http.StatusCreated, gin.H{
		"data": event,
	})
}

func getEvents(c *gin.Context) {
	res, err := models.GetAllEvent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func getEventByID(c *gin.Context) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	event, err := models.GetEventByID(eventID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, event)
}
