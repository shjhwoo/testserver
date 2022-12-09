package main

import "github.com/gin-gonic/gin"

func stress2(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "user created",
	})
}
