package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Returns a configured gin router with all routes registered
func NewServer() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return router
}
