package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *Config, pokemon ...string) error {
	name := pokemon[0]
	response, err := cfg.client.getPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a pokeball at %s\n", name)
	if rand.IntN(700) > response.BaseExperience && rand.Float32() > 0.25 {
		fmt.Printf("%s was caught!\n", name)
		cfg.pokedex[name] = response
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil
}
