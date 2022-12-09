package handler

import "github.com/gin-gonic/gin"

func Stress2(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "user created",
	})
}
