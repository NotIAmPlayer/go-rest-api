package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID      string  `json:"id"`
    Title   string  `json:"title"`
    Artist  string  `json:"artist"`
    Price   float64 `json:"price"`
}

var albums = []album {
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"message": "pong",
		})
	})
    r.GET("/albums", getAlbums)

	r.Run("localhost:8080")
}