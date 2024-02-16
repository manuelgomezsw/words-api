package search

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"words-api/internal/domain/search/services"
)

func GetByID(c *gin.Context) {
	wordID, err := strconv.ParseInt(c.Param("word_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing name",
			"error":   err.Error(),
		})
		return
	}

	word, err := services.GetByID(wordID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting word",
			"error":   err.Error(),
		})
		return
	}

	if word.WordID == 0 {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, word)
}

func GetByWord(c *gin.Context) {
	if c.Param("word") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "word is required",
		})
		return
	}

	words, err := services.GetByWord(c.Param("word"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting word",
			"error":   err.Error(),
		})
		return
	}

	if words == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, words)
}
