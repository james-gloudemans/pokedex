package main

import "fmt"

func commandExplore(cfg *Config, location ...string) error {
	response, err := cfg.client.listLocationPokemon(location[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", location)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range response.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
