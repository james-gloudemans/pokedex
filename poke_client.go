package main

import (
	"io"
	"net/http"
	"time"

	"github.com/james-gloudemans/pokedex/internal/pokecache"
)

type PokeClient struct {
	httpClient http.Client
	cache      pokecache.Cache
	baseURL    string
}

func newClient(timeout time.Duration) PokeClient {
	return PokeClient{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache:   pokecache.NewCache(5 * time.Minute),
		baseURL: "https://pokeapi.co/api/v2/",
	}
}

func (c *PokeClient) listLocations(url string) (ResponseLocations, error) {

	var locationsResponse ResponseLocations
	if r, ok := c.cache.Get(url); ok {
		lr, err := unmarshalLocations(r)
		if err != nil {
			return ResponseLocations{}, err
		}
		locationsResponse = lr
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return ResponseLocations{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return ResponseLocations{}, err
		}
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return ResponseLocations{}, err
		}
		c.cache.Add(url, data)

		lr, err := unmarshalLocations(data)
		if err != nil {
			return ResponseLocations{}, err
		}
		locationsResponse = lr
	}

	return locationsResponse, nil
}

func (c *PokeClient) listLocationPokemon(location string) (ResponseLocationAreas, error) {
	url := c.baseURL + "location-area/" + location
	var areaResponse ResponseLocationAreas
	if r, ok := c.cache.Get(url); ok {
		ar, err := unmarshalLocationAreas(r)
		if err != nil {
			return ResponseLocationAreas{}, err
		}
		areaResponse = ar
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return ResponseLocationAreas{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return ResponseLocationAreas{}, err
		}
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return ResponseLocationAreas{}, err
		}
		c.cache.Add(url, data)

		ar, err := unmarshalLocationAreas(data)
		if err != nil {
			return ResponseLocationAreas{}, err
		}
		areaResponse = ar
	}
	return areaResponse, nil
}
