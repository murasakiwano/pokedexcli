package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func runRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		fmt.Print("Pokedex ❯ ")
		scanner.Scan()

		params := extractParams(scanner.Text())

		command, ok := commands[params[0]]
		if ok {
			err := command.action(cfg, params)
			if err != nil {
				log.Printf("ERROR: %v\n", err)
			}
		} else {
			fmt.Printf("Unknown command: %s\n", params[0])
		}
	}
}

func extractParams(input string) []string {
	fields := strings.Fields(input)

	for i, field := range fields {
		fields[i] = strings.ToLower(field)
	}

	return fields
}
