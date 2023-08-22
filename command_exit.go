package main

import (
	"fmt"
	"os"
)

func commandExit(_ *config, _ []string) error {
	fmt.Println("Exiting program...")
	os.Exit(0)
	return nil
}
