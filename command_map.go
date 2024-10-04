package main

import (
	"fmt"
)

func commandMap(cfg *Config, _ ...string) error {
	if cfg.next == "" {
		fmt.Println()
		return nil
	}

	response, err := cfg.client.listLocations(cfg.next)
	if err != nil {
		return err
	}

	cfg.next = response.Next
	cfg.prev = response.Previous
	for _, location := range response.Results {
		fmt.Println(location.Name)
	}

	return nil
}
