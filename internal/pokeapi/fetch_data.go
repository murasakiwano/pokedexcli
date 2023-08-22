package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type pokeApiResponse interface {
	PokeApiExploreLocationData | PokeApiLocationResponse
}

type response[T pokeApiResponse] struct {
	data T
}

func fetchData[T pokeApiResponse](pageURL string, client *Client) (T, error) {
	if val, ok := client.cache.Get(pageURL); ok {
		res := response[T]{}
		err := json.Unmarshal(val, &res.data)
		if err != nil {
			return response[T]{}.data, err
		}

		return res.data, nil
	}

	req, err := http.NewRequest("GET", pageURL, nil)
	if err != nil {
		return response[T]{}.data, err
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return response[T]{}.data, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return response[T]{}.data, err
	}

	ret := response[T]{}
	err = json.Unmarshal(data, &ret.data)
	if err != nil {
		return response[T]{}.data, err
	}

	client.cache.Add(pageURL, data)
	return ret.data, nil
}
