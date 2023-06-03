package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default();

	server.GET("/health-check", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H {
				"message": "Go http server is working!!",
			},
		)
	})

	server.Run()
}