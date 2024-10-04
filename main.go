package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Command struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

type Config struct {
	client  PokeClient
	next    string
	prev    string
	pokedex map[string]Pokemon
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &Config{
		client:  newClient(5 * time.Second),
		next:    "https://pokeapi.co/api/v2/location-area/",
		prev:    "",
		pokedex: make(map[string]Pokemon),
	}
	for {
		fmt.Printf("pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		if len(input) == 0 {
			continue
		}
		words := strings.Fields(input)
		command, exists := getCommands()[words[0]]
		if !exists {
			fmt.Println("Unknown command.")
			continue
		} else {
			if len(words) == 1 {
				err := command.callback(cfg, []string{}...)
				if err != nil {
					fmt.Println(err)
				}
			} else if len(words) == 2 {
				err := command.callback(cfg, words[1])
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Invalid command.")
			}
			continue
		}
	}
}

func getCommands() map[string]Command {
	return map[string]Command{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Explore nearby locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Explore previous locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List all the Pokemon in a given location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon by name",
			callback:    commandCatch,
		},
	}
}
