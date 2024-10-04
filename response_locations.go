package main

import "encoding/json"

type ResponseLocations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func unmarshalLocations(data []byte) (ResponseLocations, error) {
	locationsResponse := ResponseLocations{}
	err := json.Unmarshal(data, &locationsResponse)
	if err != nil {
		return ResponseLocations{}, err
	}
	return locationsResponse, nil
}
