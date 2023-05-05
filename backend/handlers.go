package main

import (
	"github.com/leeyenter/books/backend/books"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getBooks(c *gin.Context) {
	foundBooks, err := books.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, foundBooks)
}

func addBook(c *gin.Context) {
	var book books.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := books.Add(c.Request.Context(), &book)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": book.ID})
}

type markBookAsBoughtPayload struct {
	BoughtType string `json:"boughtType"`
}

func markBookAsBought(c *gin.Context) {
	var payload markBookAsBoughtPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	if err := books.MarkBought(c.Request.Context(), id, payload.BoughtType); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

type setBookPricePayload struct {
	Source string `json:"source"`
	Price  int    `json:"price"`
}

func setBookPrice(c *gin.Context) {
	var payload setBookPricePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	if err := books.SetBookPrice(c.Request.Context(), id, payload.Source, payload.Price); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

type removeBookPricePayload struct {
	Source string `json:"source"`
	Price  int    `json:"price"`
}

func removeBookPrice(c *gin.Context) {
	var payload removeBookPricePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	if err := books.RemoveBookPrice(c.Request.Context(), id, payload.Source); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

type setBookStatusPayload struct {
	Status string `json:"status"`
}

func setBookStatus(c *gin.Context) {
	var payload setBookStatusPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	if err := books.SetBookStatus(c.Request.Context(), id, payload.Status); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}
