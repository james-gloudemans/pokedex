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
	callback    func(cfg *Config, params []string) error
}

type Config struct {
	client PokeClient
	next   string
	prev   string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &Config{
		client: newClient(5 * time.Second),
		next:   "https://pokeapi.co/api/v2/location-area/",
		prev:   "",
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
				err := command.callback(cfg, nil)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				err := command.callback(cfg, words[1:])
				if err != nil {
					fmt.Println(err)
				}
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
	}
}
