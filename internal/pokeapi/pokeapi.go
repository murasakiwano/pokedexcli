package pokeapi

const (
	baseUrl = "https://pokeapi.co/api/v2"
)

func (c *Client) ListLocations(pageURL *string) (PokeApiLocationsResponse, error) {
	url := baseUrl + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	locationData, err := fetchData[PokeApiLocationsResponse](url, c)
	if err != nil {
		return PokeApiLocationsResponse{}, err
	}

	return locationData, nil
}

func (c *Client) ExploreLocation(location string) (PokeApiLocationArea, error) {
	url := baseUrl + "/location-area/" + location

	exploreData, err := fetchData[PokeApiLocationArea](url, c)
	if err != nil {
		return PokeApiLocationArea{}, err
	}

	return exploreData, nil
}

func (c *Client) FetchPokemon(pokename string) (Pokemon, error) {
	url := baseUrl + "/pokemon/" + pokename

	pokemonData, err := fetchData[Pokemon](url, c)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemonData, nil
}
