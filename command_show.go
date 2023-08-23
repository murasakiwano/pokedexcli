package main

import "fmt"

func commandPokedex(cfg *config, _ []string) error {
	numPokemons := len(cfg.pokemonData)
	if numPokemons == 0 {
		fmt.Println("you have no pokemons!")
		return nil
	}

	fmt.Printf("You currently have %d pokemons. Here is your pokedex:\n", numPokemons)
	for name := range cfg.pokemonData {
		fmt.Printf("  - %s\n", name)
	}

	return nil
}
