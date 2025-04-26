package services

import (
	"encoding/json"
	"net/http"
)

type AnimeDetails struct {
	MalID         int     `json:"mal_id"`
	Title         string  `json:"title"`
	TitleEnglish  string  `json:"title_english"`
	TitleJapanese string  `json:"title_japanese"`
	Type          string  `json:"type"`
	Episodes      int     `json:"episodes"`
	Status        string  `json:"status"`
	Rating        string  `json:"rating"`
	Score         float64 `json:"score"`
	Synopsis      string  `json:"synopsis"`
	Year          int     `json:"year"`
	Images        struct {
		JPG struct {
			ImageURL      string `json:"image_url"`
			LargeImageURL string `json:"large_image_url"`
		} `json:"jpg"`
	} `json:"images"`
	Genres []struct {
		Name string `json:"name"`
	} `json:"genres"`
}

type AnimeDetailsResponse struct {
	Data AnimeDetails `json:"data"`
}

func GetAnimeList() ([]AnimeDetails, error) {
	resp, err := http.Get("https://api.jikan.moe/v4/anime?limit=20")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Data []AnimeDetails `json:"data"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

func GetAnimeDetails(id string) (*AnimeDetails, error) {
	resp, err := http.Get("https://api.jikan.moe/v4/anime/" + id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result AnimeDetailsResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result.Data, nil
}

// func SearchAnime(query string) ([]AnimeDetails, error) {
// 	encodedQuery := url.QueryEscape(query)
// 	apiURL := fmt.Sprintf("https://api.jikan.moe/v4/anime?q=%s&limit=20", encodedQuery)

// 	resp, err := http.Get(apiURL)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	var result struct {
// 		Data []AnimeDetails `json:"data"`
// 	}
// 	err = json.NewDecoder(resp.Body).Decode(&result)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return result.Data, nil
// }
