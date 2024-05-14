package routes

import (
	"net/http"
	"rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func updateEvent(c *gin.Context) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	// Check if the event exists
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Bind the request body to the event struct
	var updateEvent models.Event
	err = c.ShouldBindJSON(&updateEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the event exists with the given ID
	_, err = models.GetEventByID(eventID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the event
	updateEvent.ID = eventID
	err = updateEvent.UpdateEvent()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
	})
}

func deleteEvent(c *gin.Context) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the event exists with the given ID
	_, err = models.GetEventByID(eventID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.DeleteEvent(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
	})
}
