package handlers

import (
	"aniwiki/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	animes, err := services.FetchTopAnime()
	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка при получении аниме")
		return
	}

	c.HTML(http.StatusOK, "layout.html", gin.H{
		"Animes": animes,
	})
}
