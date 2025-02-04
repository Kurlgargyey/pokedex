package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mtslzr/pokeapi-go"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	locationPage int
}

var commands map[string]cliCommand
var cfg config

func init() {
	commands = make(map[string]cliCommand)

	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
	commands["map"] = cliCommand{
		name:        "map",
		description: "Displays the next 20 locations in the world of pokemon",
		callback:    commandMap,
	}
	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays the previous 20 locations in the world of pokemon",
		callback:    commandMapB,
	}
	cfg.locationPage = 0
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	lowercased := strings.ToLower(trimmed)
	return strings.Fields(lowercased)
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	commandNames := []string{"map", "mapb", "help", "exit"}

	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, cmd := range commandNames {
		fmt.Printf("%s - %s\n", commands[cmd].name, commands[cmd].description)
	}
	return nil
}

func commandMap() error {
	locs, err := pokeapi.Resource("location-area", cfg.locationPage*20)
	if err != nil {
		return err
	}
	for _, loc := range locs.Results {
		fmt.Println(loc.Name)
	}
	cfg.locationPage++
	return nil
}
func commandMapB() error {
	if cfg.locationPage <= 1 {
		fmt.Println("you're on the first page")
		return nil
	}
	cfg.locationPage-=2
	locs, err := pokeapi.Resource("location-area", cfg.locationPage*20)
	if err != nil {
		return err
	}
	for _, loc := range locs.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
