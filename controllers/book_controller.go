package controllers

import (
	"net/http"

	"github.com/Leop4rdo/go-http-server/models"
	"github.com/gin-gonic/gin"
)

func BuildBookController(server *gin.Engine) {
	server.GET("/api/v1/books", listBooks)
	server.POST("/api/v1/books", createBook)
	server.GET("/api/v1/books/:id", findBookById)
	server.DELETE("/api/v1/books/:id", deleteBookById)
	server.PUT("/api/v1/books/:id", updateBook)
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

func findBookById(context *gin.Context) {
	var book models.Book

	if err := models.Database.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{ "error": "Entity not found"})
		return
	}

	context.JSON(http.StatusOK, book)
}

func deleteBookById(context *gin.Context) {
	var book models.Book

	if err := models.Database.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{ "error": "Entity not found"})
		return
	}

	models.Database.Delete(&book)

	context.Status(http.StatusNoContent)
}

type updateBookInput struct {
	Title string `json:"title"`
	Author string `json:"author"`
}

func updateBook(context *gin.Context) {
	var input updateBookInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H { "error": "Invalid Request", "details": err.Error() })
		return
	}

	var book models.Book
	if err := models.Database.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "entity not found"})
	}

	models.Database.Model(&book).Updates(input)
	context.JSON(http.StatusOK, &book)
}
