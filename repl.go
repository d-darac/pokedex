package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func runRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		commandName := strings.ToLower(input[0])

		args := []string{}
		if len(input) > 1 {
			args = input[1:]
		}

		cmd, exists := getCommands()[commandName]

		if !exists {
			fmt.Printf("Unknown command: %s", commandName)
			fmt.Println()
			continue
		}

		if err := cmd.callback(cfg, args...); err != nil {
			fmt.Printf("%v", err)
			fmt.Println()
		}
	}

}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempts to catch the given pokemon.",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex.",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore <location_area_name>",
			description: "Displays a list of pokemon found in the given locatoin area.",
			callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message.",
			callback:    commandHelp,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Displays details about the given caught pokemon.",
			callback:    commandInspect,
		},
		"map": {
			name:        "map",
			description: "Displays a list of location area with 20 lines per page. Eachs subsequent call to map command displays the next 20 locations.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Same as the map command, but displays the previous 20 locations. Used to traverse the list backwards.",
			callback:    commandMapb,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays a list of caught pokemon.",
			callback:    commandPokedex,
		},
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	return words
}
