package main

import (
	zinc "Indexador/zincsearch"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")

	v1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Enron-Email Index ZincSearch API v1",
		})
	})

	v1.GET("/search", func(c *gin.Context) {
		searchTerm := c.Query("word")
		if searchTerm == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'word' is required"})
			return
		}
		result := zinc.SearchDocument(searchTerm)
		c.JSON(http.StatusAccepted, result)

	})

	r.Run(":3000")
}
