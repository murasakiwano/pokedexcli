package main

import (
	"bufio"
	"fmt"
	"strings"
)

func runRepl() {
	var w string
	scanner := bufio.NewScanner(strings.NewReader(w))

	commandMap := getCommands()
	conf := config{}

	for {
		fmt.Print("Pokedex ❯ ")
		fmt.Scanln(&w)
		scanner.Scan()

		if command, ok := commandMap[w]; ok {
			command.action(conf)
		} else {
			fmt.Printf("Unknown command: %s\n", w)
		}
	}
}
