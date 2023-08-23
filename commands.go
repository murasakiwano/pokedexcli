package main

type cliCommand struct {
	action      func(*config, []string) error
	name        string
	description string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			name:        "catch",
			description: "Try to catch a pokemon!",
			action:      commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			action:      commandExit,
		},
		"explore": {
			name:        "explore",
			description: "Displays the names of the pokemons in a given area",
			action:      commandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			action:      commandHelp,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon in your pokedex",
			action:      commandInspect,
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 location areas",
			action:      commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas",
			action:      commandMapB,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays the pokemons in your pokedex",
			action:      commandPokedex,
		},
	}
}
