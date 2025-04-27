package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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

	CorrectedYear     string
	CorrectedEpisodes string
}

type AnimeDetailsResponse struct {
	Data AnimeDetails `json:"data"`
}

func GetAnimeList(filter string) ([]AnimeDetails, error) {
	resp, err := http.Get("https://api.jikan.moe/v4/top/anime?" + filter + "&limit=12")
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

	for i := range result.Data {
		result.Data[i].CorrectedYear = correctDate(result.Data[i].Year)
		result.Data[i].CorrectedEpisodes = correctEpisodes(result.Data[i].Episodes)
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

	result.Data.CorrectedYear = correctDate(result.Data.Year)
	result.Data.CorrectedEpisodes = correctEpisodes(result.Data.Episodes)

	return &result.Data, nil
}

func SearchAnime(query string) ([]AnimeDetails, error) {
	encodedQuery := url.QueryEscape(query)
	resp, err := http.Get("https://api.jikan.moe/v4/anime?q=" + encodedQuery + "&limit=12")
	if err != nil {
		return nil, fmt.Errorf("error making request to Jikan API: %v", err)
	}
	defer resp.Body.Close()

	var result struct {
		Data []AnimeDetails `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	for i := range result.Data {
		result.Data[i].CorrectedYear = correctDate(result.Data[i].Year)
		result.Data[i].CorrectedEpisodes = correctEpisodes(result.Data[i].Episodes)
	}

	return result.Data, nil
}

func SearchAnimeWithPagination(query string, page int) ([]AnimeDetails, error) {
	encodedQuery := url.QueryEscape(query)
	resp, err := http.Get(fmt.Sprintf("https://api.jikan.moe/v4/anime?q=%s&limit=12&page=%d", encodedQuery, page))
	if err != nil {
		return nil, fmt.Errorf("error making request to Jikan API: %v", err)
	}
	defer resp.Body.Close()

	var result struct {
		Data []AnimeDetails `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	for i := range result.Data {
		result.Data[i].CorrectedYear = correctDate(result.Data[i].Year)
		result.Data[i].CorrectedEpisodes = correctEpisodes(result.Data[i].Episodes)
	}

	return result.Data, nil
}

func correctDate(year int) string {
	if year == 0 {
		return "nd"
	}
	return fmt.Sprintf("%d", year)
}

func correctEpisodes(ep int) string {
	if ep == 0 {
		return "N/A"
	}
	return fmt.Sprintf("%d", ep)
}
