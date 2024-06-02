package main

import "fmt"

func callbackHelp() error {
	fmt.Println("Welcome to the PokeDex help menu.")
	fmt.Println("Here are you available commands")

	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf(" - %s: %s \n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}
