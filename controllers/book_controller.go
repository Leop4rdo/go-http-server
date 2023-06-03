package controllers

import (
	"net/http"

	"github.com/Leop4rdo/go-http-server/models"
	"github.com/gin-gonic/gin"
)

func BuildBookController(server *gin.Engine) {
	server.GET("/api/v1/books", listBooks)
	server.POST("/api/v1/books", createBook)
}

func listBooks(context *gin.Context) {
	var books []models.Book
	models.Database.Find(&books)
	context.JSON(http.StatusOK, books)
} 

func createBook(context *gin.Context) {
	var input createBookInput
	
	if err:= context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H { "error": err.Error()})
		return
	}

	book := models.Book {
		Title: input.Title,
		Author: input.Author,
	}

	models.Database.Create(&book)

	context.JSON(http.StatusCreated, book)
}

type createBookInput struct {
	Title string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}
