package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type PokeApiLocationResponse struct {
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
	Count int `json:"count"`
}

func GetLocations(endpoint string) (PokeApiLocationResponse, error) {
	res, err := http.Get(endpoint)
	if err != nil {
		return PokeApiLocationResponse{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	data := PokeApiLocationResponse{}
	json.Unmarshal(body, &data)

	return data, nil
}
