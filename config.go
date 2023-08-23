package main

import (
	"github.com/murasakiwano/pokedexcli/internal/pokeapi"
)

type config struct {
	previousLocationsUrl *string
	nextLocationsUrl     *string
	pokemonData          map[string]pokeapi.Pokemon
	pokeapiClient        pokeapi.Client
}
