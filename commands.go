package main

type cliCommand struct {
	action      func(config) error
	name        string
	description string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			action:      commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			action:      commandExit,
		},
	}
}
