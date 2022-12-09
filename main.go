package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
}

var Server *App

func NewServer() {
	if Server != nil {
		return
	}

	app := &App{}
	router := gin.Default()

	router.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
			AllowCredentials: true,
			AllowHeaders:     []string{"withCredentials", "Content-Type"},
			MaxAge:           0,
		},
	))

	router.GET("/stress1", stress1)
	router.POST("/stress2", stress2)

	app.Router = router
	Server = app
}

func Start(address string) error {
	return Server.Router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func main() {

	NewServer()

	err := Start("http://103.218.157.114")
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
