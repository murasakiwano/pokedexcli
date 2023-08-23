package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, params []string) error {
	if len(params) <= 1 {
		return errors.New("must provide a pokemon to inspect")
	}

	if _, ok := cfg.pokemonData[params[1]]; !ok {
		fmt.Println("you must catch that pokemon first")
	}

	pokemon := cfg.pokemonData[params[1]]

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, pokeType := range pokemon.Types {
		fmt.Printf("  - %s\n", pokeType.Type.Name)
	}

	return nil
}
