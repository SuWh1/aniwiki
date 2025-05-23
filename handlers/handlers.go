package handlers

import (
	"aniwiki/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	animes, err := services.GetAnimeList("airing")
	if err != nil {
		log.Println("Error fetching anime list:", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.HTML(http.StatusOK, "anime.html", gin.H{
		"title":  "AniWiki - Anime",
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

func AnimeSearchPageHandler(c *gin.Context) {
	query := c.DefaultQuery("q", "")

	if query == "" {
		c.HTML(http.StatusOK, "searchAnime.html", gin.H{
			"query":  "",
			"Animes": nil,
		})
		return
	}

	searchResults, err := services.SearchAnime(query)
	if err != nil {
		c.HTML(http.StatusOK, "searchAnime.html", gin.H{
			"query":  query,
			"Animes": nil,
			"error":  err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "searchAnime.html", gin.H{
		"title":  "Search - " + query,
		"query":  query,
		"Animes": searchResults,
	})
}

func MangaHandler(c *gin.Context) {
	mangas, err := services.GetMangaList("publishing")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error fetching manga")
		return
	}

	c.HTML(http.StatusOK, "manga.html", gin.H{
		"title":  "AniWiki - Manga",
		"Mangas": mangas,
	})
}

func MangaDetailsHandler(c *gin.Context) {
	id := c.Param("id")

	mangaDetails, err := services.GetMangaDetails(id)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error fetching anime details")
		return
	}

	c.HTML(http.StatusOK, "mangaDetails.html", gin.H{
		"title": mangaDetails.Title + " - AniWiki",
		"Manga": mangaDetails,
	})
}

func MangaSearchPageHandler(c *gin.Context) {
	query := c.DefaultQuery("q", "")

	if query == "" {
		c.HTML(http.StatusOK, "searchManga.html", gin.H{
			"query":  "",
			"Mangas": nil,
		})
		return
	}

	searchResults, err := services.SearchManga(query)
	if err != nil {
		c.HTML(http.StatusOK, "searchManga.html", gin.H{
			"query":  query,
			"Mangas": nil,
			"error":  err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "searchManga.html", gin.H{
		"title":  "Search - " + query,
		"query":  query,
		"Mangas": searchResults,
	})
}

func CharacterHandler(c *gin.Context) {
	characters, err := services.GetCharacterList()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error fetching manga")
		return
	}

	c.HTML(http.StatusOK, "characters.html", gin.H{
		"title":      "AniWiki - Characters",
		"Characters": characters,
	})
}

func CharacterDetailsHandler(c *gin.Context) {
	id := c.Param("id")

	characterDetails, err := services.GetCharacterDetails(id)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error fetching anime details")
		return
	}

	c.HTML(http.StatusOK, "characterDetails.html", gin.H{
		"name":      characterDetails.Name + " - AniWiki",
		"Character": characterDetails,
	})
}

func CharacterSearchPageHandler(c *gin.Context) {
	query := c.DefaultQuery("q", "")

	if query == "" {
		c.HTML(http.StatusOK, "searchCharacter.html", gin.H{
			"query":      "",
			"Characters": nil,
		})
		return
	}

	searchResults, err := services.SearchCharacter(query)
	if err != nil {
		c.HTML(http.StatusOK, "searchCharacter.html", gin.H{
			"query":      query,
			"Characters": nil,
			"error":      err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "searchCharacter.html", gin.H{
		"title":      "Search - " + query,
		"query":      query,
		"Characters": searchResults,
	})
}
