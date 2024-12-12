package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type meeting struct {
	ID          string   `json:"id"`
	MeetingFrom string   `json:"meeting_from"`
	MeetingTo   []string `json:"meeting_to"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Date        string   `json:"date"`
	StartTime   string   `json:"start_time"`
	EndTime     string   `json:"end_time"`
}

var meetings = []meeting{}

func getMeetings(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, meetings)
}

func getMeetingByID(c *gin.Context) {
	id := c.Param("id")

	for _, m := range meetings {
		if m.ID == id {
			c.IndentedJSON(http.StatusOK, m)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "meetings not found"})
}

func postMeetings(c *gin.Context) {
	var newMeeting meeting

	if err := c.BindJSON(&newMeeting); err != nil {
		return
	}

	fmt.Println("POSTED!")

	meetings = append(meetings, newMeeting)
	c.IndentedJSON(http.StatusCreated, newMeeting)
}

func main() {
	r := gin.Default()
	/*
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	*/

	r.GET("/meetings", getMeetings)
	r.GET("/meetings/:id", getMeetingByID)
	r.POST("/meetings", postMeetings)

	r.Run("localhost:8080")
}
