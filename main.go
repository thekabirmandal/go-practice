package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/uuid", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"uuid": uuid.New().String(),
		})
	})

	r.Run(":8080") // listens and serves on 0.0.0.0:8080
}
