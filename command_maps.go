package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config) error {
	locations, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationsUrl = locations.Next
	cfg.previousLocationsUrl = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapB(cfg *config) error {
	if cfg.previousLocationsUrl == nil {
		return errors.New("can not go back from first page")
	}

	locations, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationsUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationsUrl = locations.Next
	cfg.previousLocationsUrl = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}
