package handler

import (
	"bookstore/model"
	"bookstore/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	service service.BookService
}

func NewBookHandler(s service.BookService) *BookHandler {
	return &BookHandler{service: s}
}

func (h *BookHandler) GetAll(c *gin.Context) {
	books, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) Create(c *gin.Context) {
	var input model.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if input.ReleaseYear < 1980 || input.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "release_year must be between 1980 and 2024"})
		return
	}
	if input.TotalPage > 100 {
		input.Thickness = "tebal"
	} else {
		input.Thickness = "tipis"
	}

	err := h.service.Create(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func (h *BookHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
