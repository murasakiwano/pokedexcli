package main

import "github.com/murasakiwano/pokedexcli/internal/pokeapi"

type config struct {
	previousLocationsUrl *string
	nextLocationsUrl     *string
	pokeapiClient        pokeapi.Client
}
