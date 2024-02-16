package router

import (
	"github.com/gin-gonic/gin"
	"words-api/internal/infraestructure/controller/registry"
	"words-api/internal/infraestructure/controller/search"
)

func mapURLs(router *gin.Engine) {
	registryURLs(router)
	searchURLs(router)
}

func registryURLs(router *gin.Engine) {
	router.POST("/words", registry.Create)
	router.PUT("/words/:word_id", registry.Update)
	router.DELETE("/words/:word_id", registry.Delete)
}

func searchURLs(router *gin.Engine) {
	router.GET("/words/:word", search.GetByWord)
	router.GET("/word/:word_id", search.GetByID)
}
