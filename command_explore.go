package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, params []string) error {
	if len(params) <= 1 {
		return errors.New("must provide location to explore")
	}

	exploreData, err := cfg.pokeapiClient.ExploreLocation(params[1])
	if err != nil {
		return err
	}

	pokemons := make([]string, len(exploreData.PokemonEncounters))

	for i, encounter := range exploreData.PokemonEncounters {
		pokemons[i] = encounter.Pokemon.Name
	}

	for _, pokemon := range pokemons {
		fmt.Printf("- %s\n", pokemon)
	}

	return nil
}
