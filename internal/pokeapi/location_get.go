package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	if data, ok := c.cache.Get(url); ok {
		var location Location
		if err := json.Unmarshal(data, &location); err != nil {
			return Location{}, err
		}
		return location, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return Location{}, fmt.Errorf("location %q not found (status %d)", locationName, res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	var location Location
	if err := json.Unmarshal(data, &location); err != nil {
		return Location{}, err
	}

	c.cache.Add(url, data)
	return location, nil
}
