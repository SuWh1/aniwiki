package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type CharacterDetails struct {
	MalID     int      `json:"mal_id"`
	URL       string   `json:"url"`
	Name      string   `json:"name"`
	NameKanji string   `json:"name_kanji"`
	Nicknames []string `json:"nicknames"`
	Favorites int      `json:"favorites"`
	About     string   `json:"about"`
	Images    struct {
		JPG struct {
			ImageURL      string `json:"image_url"`
			SmallImageURL string `json:"small_image_url"`
		} `json:"jpg"`
	} `json:"images"`
}

type CharacterDetailsResponse struct {
	Data CharacterDetails `json:"data"`
}

func GetCharacterList() ([]CharacterDetails, error) {
	resp, err := http.Get("https://api.jikan.moe/v4/top/characters?&limit=12")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Data []CharacterDetails `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

func GetCharacterDetails(id string) (*CharacterDetails, error) {
	resp, err := http.Get("https://api.jikan.moe/v4/characters/" + id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result CharacterDetailsResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	if len(result.Data.Nicknames) == 0 {
		result.Data.Nicknames = []string{"No nicknames available"}
	}

	return &result.Data, nil
}

func SearchCharacter(query string) ([]CharacterDetails, error) {
	encodedQuery := url.QueryEscape(query)
	resp, err := http.Get("https://api.jikan.moe/v4/characters?q=" + encodedQuery + "&limit=12")
	if err != nil {
		return nil, fmt.Errorf("error making request to Jikan API: %v", err)
	}
	defer resp.Body.Close()

	var result struct {
		Data []CharacterDetails `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

func SearchCharacterWithPagination(query string, page int) ([]CharacterDetails, error) {
	encodedQuery := url.QueryEscape(query)
	resp, err := http.Get(fmt.Sprintf("https://api.jikan.moe/v4/characters?q=%s&limit=12&page=%d", encodedQuery, page))
	if err != nil {
		return nil, fmt.Errorf("error making request to Jikan API: %v", err)
	}
	defer resp.Body.Close()

	var result struct {
		Data []CharacterDetails `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}
