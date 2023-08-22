package pokeapi

const (
	baseUrl = "https://pokeapi.co/api/v2"
)

func (c *Client) ListLocations(pageURL *string) (PokeApiLocationResponse, error) {
	url := baseUrl + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	locationData, err := fetchData[PokeApiLocationResponse](url, c)
	if err != nil {
		return PokeApiLocationResponse{}, err
	}

	return locationData, nil
}

func (c *Client) ExploreLocation(location string) (PokeApiExploreLocationData, error) {
	url := baseUrl + "/location-area/" + location

	exploreData, err := fetchData[PokeApiExploreLocationData](url, c)
	if err != nil {
		return PokeApiExploreLocationData{}, err
	}

	return exploreData, nil
}
