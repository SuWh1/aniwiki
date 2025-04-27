package main

import (
	"aniwiki/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/anime", handlers.HomeHandler)
	r.GET("/anime/:id", handlers.AnimeDetailsHandler)
	r.GET("/anime/search", handlers.AnimeSearchPageHandler)
	r.GET("/api/anime/search", handlers.AnimeSearchAPIHandler)

	r.GET("/manga", handlers.MangaHandler)
	r.GET("/manga/:id", handlers.MangaDetailsHandler)
	r.GET("/manga/search", handlers.MangaSearchPageHandler)
	r.GET("/api/manga/search", handlers.MangaSearchAPIHandler)

	r.GET("/characters", handlers.CharacterHandler)
	r.GET("/characters/:id", handlers.CharacterDetailsHandler)
	r.GET("/characters/search", handlers.CharacterSearchPageHandler)
	r.GET("/api/characters/search", handlers.CharacterSearchAPIHandler)

	r.Run(":8080")
}
