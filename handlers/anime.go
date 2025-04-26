package handlers

import (
	"aniwiki/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	animes, err := services.GetAnimeList()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error fetching anime details")
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":  "AniWiki - Home",
		"Animes": animes,
	})
}

func AnimeHandler(c *gin.Context) {
	animes, err := services.GetAnimeList()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error fetching anime details")
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":  "AniWiki - Home",
		"Animes": animes,
	})
}

func AnimeDetailsHandler(c *gin.Context) {
	id := c.Param("id")

	animeDetails, err := services.GetAnimeDetails(id)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error fetching anime details")
		return
	}

	c.HTML(http.StatusOK, "animeDetails.html", gin.H{
		"title": animeDetails.Title + " - AniWiki",
		"Anime": animeDetails,
	})
}

// func SearchHandler(c *gin.Context) {
// 	query := c.Query("q")
// 	if query == "" {
// 		c.Redirect(http.StatusFound, "/")
// 		return
// 	}

// 	searchResults, err := services.SearchAnime(query)
// 	if err != nil {
// 		c.String(http.StatusInternalServerError, "Error searching anime")
// 		return
// 	}

// 	c.HTML(http.StatusOK, "searchResults.html", gin.H{
// 		"title":   "Search Results - " + query,
// 		"query":   query,
// 		"results": searchResults,
// 	})
// }
