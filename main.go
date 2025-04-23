package main

import (
	"aniwiki/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", handlers.HomeHandler)

	r.Run(":8080")
}
