package main

type config struct {
	previousLocationsUrl *string
	nextLocationsUrl     string
	pageSize             uint
}

func (c *config) SetConfig(next string, previous *string) {
	c.nextLocationsUrl = next
	c.previousLocationsUrl = previous
}

func NewConfig() config {
	return config{
		previousLocationsUrl: nil,
		nextLocationsUrl:     "https://pokeapi.co/api/v2/location-area?limit=20",
		pageSize:             20,
	}
}
