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

	meetings = append(meetings, newMeeting)
	c.IndentedJSON(http.StatusCreated, newMeeting)
}

func putMeeting(c *gin.Context) {
	var updatedMeeting meeting

	if err := c.BindJSON(&updatedMeeting); err != nil {
		return
	}

	id := c.Param("id")
	found := false

	for _, m := range meetings {
		if m.ID == id {
			found = true

			fmt.Println(updatedMeeting)

			if updatedMeeting.MeetingFrom != "" {
				m.MeetingFrom = updatedMeeting.MeetingFrom
			}

			if len(updatedMeeting.MeetingTo) != 0 {
				m.MeetingTo = updatedMeeting.MeetingTo
			}

			if updatedMeeting.Title != "" {
				m.Title = updatedMeeting.Title
			}

			if updatedMeeting.Description != "" {
				m.Description = updatedMeeting.Description
			}

			if updatedMeeting.Date != "" {
				m.Date = updatedMeeting.Date
			}

			if updatedMeeting.StartTime != "" {
				m.StartTime = updatedMeeting.StartTime
			}

			if updatedMeeting.EndTime != "" {
				m.EndTime = updatedMeeting.EndTime
			}

			c.IndentedJSON(http.StatusOK, m)
			break
		}
	}

	if !found {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "meetings not found"})
	}
}

func deleteMeeting(c *gin.Context) {
	id := c.Param("id")
	done := false

	for i, m := range meetings {
		if m.ID == id {
			meetings = append(meetings[:i], meetings[i+1:]...)
			done = true

			c.IndentedJSON(http.StatusOK, gin.H{"message": "meeting deleted successfully"})
			break
		}
	}

	if !done {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "meetings not found"})
	}
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
	r.PUT("/meetings/:id", putMeeting)
	r.DELETE("/meetings/:id", deleteMeeting)

	r.Run("localhost:8080")
}
