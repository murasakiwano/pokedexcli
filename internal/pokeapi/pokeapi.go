package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	baseUrl = "https://pokeapi.co/api/v2"
)

type PokeApiLocationResponse struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
	Count int `json:"count"`
}

func (c *Client) ListLocations(pageURL *string) (PokeApiLocationResponse, error) {
	url := baseUrl + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := PokeApiLocationResponse{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return PokeApiLocationResponse{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeApiLocationResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokeApiLocationResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokeApiLocationResponse{}, err
	}

	locationsResp := PokeApiLocationResponse{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return PokeApiLocationResponse{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}
