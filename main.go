package main

import (
	"time"

	"github.com/murasakiwano/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{pokeapiClient: pokeClient, pokemonData: map[string]pokeapi.Pokemon{}}

	runRepl(cfg)
}
