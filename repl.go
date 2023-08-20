package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func runRepl(cfg *config) {
	var w string
	scanner := bufio.NewScanner(strings.NewReader(w))
	commandMap := getCommands()

	for {
		fmt.Print("Pokedex ❯ ")
		fmt.Scanln(&w)
		scanner.Scan()

		command, ok := commandMap[w]
		if ok {
			err := command.action(cfg)
			if err != nil {
				log.Printf("ERROR: %v\n", err)
			}
		} else {
			fmt.Printf("Unknown command: %s\n", w)
		}
	}
}
