package main

import (
	"errors"
	"fmt"

	"github.com/murasakiwano/pokedexcli/internal/pokeapi"
)

func commandMap(conf *config) error {
	locations, err := fetchAndDisplayLocations(conf.nextLocationsUrl)
	if err != nil {
		return err
	}

	conf.SetConfig(locations.Next, locations.Previous)

	return nil
}

func commandMapB(conf *config) error {
	if conf.previousLocationsUrl == nil {
		return errors.New("can not go back from first page")
	}

	locations, err := fetchAndDisplayLocations(*conf.previousLocationsUrl)
	if err != nil {
		return err
	}

	conf.SetConfig(locations.Next, locations.Previous)

	return nil
}

func fetchAndDisplayLocations(endpoint string) (pokeapi.PokeApiLocationResponse, error) {
	locations, err := pokeapi.GetLocations(endpoint)
	if err != nil {
		return locations, err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return locations, nil
}
