package main

import "fmt"

func commandInspect(cfg *Config, params ...string) error {
	name := params[0]
	if pokemon, ok := cfg.pokedex[name]; ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, typeStruct := range pokemon.Types {
			fmt.Printf(" - %s\n", typeStruct.Type.Name)
		}
	} else {
		fmt.Println("You have not caught that pokemon")
	}
	return nil
}
