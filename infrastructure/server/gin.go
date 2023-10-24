package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "active",
			"message": "Welcome to the Toko Belanja API",
			"version": "1.0.0",
		})
	})

	r.Run()
}
