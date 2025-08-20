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

		commandName := input[0]

		cmd, exists := getCommands()[commandName]

		if !exists {
			fmt.Printf("Unknown command: %s", commandName)
			fmt.Println()
			continue
		}

		if err := cmd.callback(cfg); err != nil {
			fmt.Printf("%v", err)
			fmt.Println()
		}
	}

}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex.",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message.",
			callback:    commandHelp,
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
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}
