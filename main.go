package main

import (
	"gin_notes/controllers"
	"gin_notes/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Static("/static", "./static")

	r.LoadHTMLGlob("templates/**/**")

	models.ConnectDatabase()
	models.DbMigrate()

	r.GET("/notes", controllers.NotesIndex)
	r.GET("/notes/new", controllers.NotesNew)
	r.POST("notes", controllers.NotesCreate)

	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "views/index.html", gin.H{
			"title": "Notes application",
		})

	})

	log.Println("Server Started")
	r.Run() // Default localhost is 8080
}
