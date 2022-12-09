package handler

import "github.com/gin-gonic/gin"

func Stress1(c *gin.Context) {

	fib(10)

	c.JSON(200, gin.H{
		"message": fib(10),
	})
}

func fib(n int) int {
	if n >= 0 && n <= 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}
