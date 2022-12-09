package main

import (
	"log"

	"mockserver/pkg/handler"

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

	router.GET("/stress1", handler.Stress1)
	router.POST("/stress2", handler.Stress2)

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

	err := Start("0.0.0.0:8080")
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
