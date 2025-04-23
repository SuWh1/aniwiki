package services

import (
	"encoding/json"
	"net/http"
)

type Anime struct {
	Title  string `json:"title"`
	Images struct {
		JPG struct {
			ImageURL string `json:"image_url"`
		} `json:"jpg"`
	} `json:"images"`
}

type TopAnimeResponse struct {
	Data []Anime `json:"data"`
}

func FetchTopAnime() ([]Anime, error) {
	resp, err := http.Get("https://api.jikan.moe/v4/top/anime")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result TopAnimeResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}
