package main

import "fmt"

func commandHelp(params ...string) error {
	commandNames := []string{
		"pokedex",
		"map",
		"mapb",
		"explore",
		"catch",
		"help",
		"exit",
	}

	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, cmd := range commandNames {
		fmt.Printf("%s - %s\n", commands[cmd].name, commands[cmd].description)
	}
	return nil
}
