package main

import "fmt"

func commandHelp(_ *config, _ []string) error {
	fmt.Printf("\nWelcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")

	commands := getCommands()

	for entry := range commands {
		command := commands[entry]
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	fmt.Println()

	return nil
}
