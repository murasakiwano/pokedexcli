package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, params []string) error {
	if len(params) <= 1 {
		return errors.New("must provide a pokemon to catch")
	}

	if _, ok := cfg.pokemonData[params[1]]; ok {
		fmt.Println("you have already caught that pokemon!")
		return nil
	}

	pokemon, err := cfg.pokeapiClient.FetchPokemon(params[1])
	if err != nil {
		return err
	}

	chance := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)

	if chance > 40 {
		fmt.Printf("%v was caught!\n", pokemon.Name)
		cfg.pokemonData[params[1]] = pokemon
	} else {
		fmt.Printf("%v escaped!\n", pokemon.Name)
	}

	return nil
}
