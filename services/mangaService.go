package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type MangaDetails struct {
	MalID         int     `json:"mal_id"`
	Title         string  `json:"title"`
	TitleEnglish  string  `json:"title_english"`
	TitleJapanese string  `json:"title_japanese"`
	Type          string  `json:"type"`
	Chapters      int     `json:"chapters"`
	Volumes       int     `json:"volumes"`
	Status        string  `json:"status"`
	Score         float64 `json:"score"`
	Synopsis      string  `json:"synopsis"`
	Published     struct {
		From string `json:"from"`
		To   string `json:"to"`
	} `json:"published"`
	Images struct {
		JPG struct {
			ImageURL      string `json:"image_url"`
			LargeImageURL string `json:"large_image_url"`
		} `json:"jpg"`
	} `json:"images"`
	Genres []struct {
		Name string `json:"name"`
	} `json:"genres"`

	PublishedFrom string
	PublishedTo   string

	CorrectedChapters string
	CorrectedVolumes  string
}

type MangaDetailsResponse struct {
	Data MangaDetails `json:"data"`
}

func GetMangaList(filter string) ([]MangaDetails, error) {
	resp, err := http.Get("https://api.jikan.moe/v4/top/manga?" + filter + "&limit=12")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Data []MangaDetails `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	for i := range result.Data {
		result.Data[i].PublishedFrom = cleanDate(result.Data[i].Published.From)
		result.Data[i].PublishedTo = cleanDate(result.Data[i].Published.To)

		result.Data[i].CorrectedChapters = correctChapterVolume(result.Data[i].Chapters)
		result.Data[i].CorrectedVolumes = correctChapterVolume(result.Data[i].Volumes)
	}

	return result.Data, nil
}

func GetMangaDetails(id string) (*MangaDetails, error) {
	resp, err := http.Get("https://api.jikan.moe/v4/manga/" + id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result MangaDetailsResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	manga := &result.Data
	manga.PublishedFrom = cleanDate(manga.Published.From)
	if manga.Published.To == "" {
		manga.PublishedTo = "Present"
	} else {
		manga.PublishedTo = cleanDate(manga.Published.To)
	}

	manga.CorrectedChapters = correctChapterVolume(manga.Chapters)
	manga.CorrectedVolumes = correctChapterVolume(manga.Volumes)

	return manga, nil
}

func SearchManga(query string) ([]MangaDetails, error) {
	encodedQuery := url.QueryEscape(query)
	resp, err := http.Get("https://api.jikan.moe/v4/manga?q=" + encodedQuery + "&limit=12")
	if err != nil {
		return nil, fmt.Errorf("error making request to Jikan API: %v", err)
	}
	defer resp.Body.Close()

	var result struct {
		Data []MangaDetails `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	for i := range result.Data {
		result.Data[i].PublishedFrom = cleanDate(result.Data[i].Published.From)
		result.Data[i].PublishedTo = cleanDate(result.Data[i].Published.To)

		result.Data[i].CorrectedChapters = correctChapterVolume(result.Data[i].Chapters)
		result.Data[i].CorrectedVolumes = correctChapterVolume(result.Data[i].Volumes)
	}

	return result.Data, nil
}

func cleanDate(datetime string) string {
	if datetime == "" {
		return "nd"
	}
	parts := strings.Split(datetime, "T")
	return parts[0]
}

func correctChapterVolume(chapterOrVolume int) string {
	if chapterOrVolume == 0 {
		return "N/A"
	}
	return fmt.Sprintf("%d", chapterOrVolume)
}
