package main

import (
	zinc "Indexador/zincsearch"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		result := zinc.SearchDocument("Gold")
		c.JSON(http.StatusOK, result)
	})

	r.Run(":3000")
}
