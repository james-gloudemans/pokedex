package main

import "fmt"

func commandMapb(cfg *Config, _ ...string) error {
	if cfg.prev == "" {
		fmt.Println()
		return nil
	}
	response, err := cfg.client.listLocations(cfg.prev)
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
