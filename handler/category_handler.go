package handler

import (
    "bookstore/service"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type CategoryHandler struct {
    service service.CategoryService
}

func NewCategoryHandler(s service.CategoryService) *CategoryHandler {
    return &CategoryHandler{service: s}
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
    categories, err := h.service.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, categories)
}

func (h *CategoryHandler) Create(c *gin.Context) {
    var input struct {
        Name string `json:"name" binding:"required"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := h.service.Create(input.Name)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusCreated)
}

func (h *CategoryHandler) GetByID(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    category, err := h.service.GetByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
        return
    }
    c.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) Delete(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    err := h.service.Delete(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
        return
    }
    c.Status(http.StatusNoContent)
}

func (h *CategoryHandler) GetBooksByCategory(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    books, err := h.service.GetBooksByCategory(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, books)
}