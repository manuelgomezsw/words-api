package registry

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"words-api/internal/domain"
	"words-api/internal/domain/registry/services"
)

func Create(c *gin.Context) {
	var newWord domain.Word
	if err := c.ShouldBindJSON(&newWord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing body",
			"error":   err.Error(),
		})
		return
	}

	if err := services.Create(&newWord); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error posting word",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newWord)
}

func Update(c *gin.Context) {
	var word domain.Word
	if err := c.ShouldBindJSON(&word); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing body",
			"error":   err.Error(),
		})
		return
	}

	wordID, err := strconv.ParseInt(c.Param("word_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing word_id",
			"error":   err.Error(),
		})
		return
	}

	if err := services.Update(wordID, &word); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating word",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, word)
}

func Delete(c *gin.Context) {
	wordID, err := strconv.ParseInt(c.Param("word_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing word_id",
			"error":   err.Error(),
		})
		return
	}

	if err := services.Delete(wordID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting word",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, nil)
}
