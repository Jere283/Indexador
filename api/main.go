package main

import (
	zinc "Indexador/zincsearch"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	v1 := r.Group("/api/v1")

	v1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Enron-Email Index ZincSearch API v1",
		})
	})

	v1.GET("/search/:word", func(c *gin.Context) {
		searchTerm := c.Param("word")
		if searchTerm == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'word' is required"})
			return
		}
		result := zinc.SearchDocument(searchTerm)
		c.JSON(http.StatusAccepted, result)

	})

	r.Run(":3000")
}
