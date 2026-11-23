package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// cache hit: unmarshal the saved bytes, no HTTP request
	if data, ok := c.cache.Get(url); ok {
		fmt.Println("[cache hit]", url) // TEMP
		var locations RespShallowLocations
		if err := json.Unmarshal(data, &locations); err != nil {
			return RespShallowLocations{}, err
		}
		return locations, nil
	}

	// cache miss: make the request
	fmt.Println("[network request]", url) // TEMP
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return RespShallowLocations{}, fmt.Errorf("bad status code: %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	var locations RespShallowLocations
	if err := json.Unmarshal(data, &locations); err != nil {
		return RespShallowLocations{}, err
	}

	// only cache after the request and decoding both succeeded
	c.cache.Add(url, data)
	return locations, nil
}
