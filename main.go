package main

import (
	"net/http"

	"github.com/Leop4rdo/go-http-server/models"
	"github.com/Leop4rdo/go-http-server/controllers"
	"github.com/gin-gonic/gin"
)


func main() {
	server := gin.Default()

	models.ConnectToDatabase()

	buildRoutes(server)

	server.Run()
}

func buildRoutes(server *gin.Engine) {
	
	server.GET("/health-check", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H {
				"message": "Go http server is working!!",
			},
		)
	})

	controllers.BuildBookController(server)
}
