package main

import (
	"fmt"
	"math/rand"
	"os"
	"pokedex/internal/pokemon_api"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(...string) error
}

type config struct {
	areasNext string
	areasPrev string
	area      string
	pokedex   map[string]pokemon_api.Pokemon
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
	commands["explore"] = cliCommand{
		name:        "explore <area-id>",
		description: "Explores the given area",
		callback:    commandExplore,
	}
	commands["catch"] = cliCommand{
		name:        "catch <pokemon-id>",
		description: "Attempts to catch the given pokemon",
		callback:    commandCatch,
	}

	cfg.areasNext = "https://pokeapi.co/api/v2/location-area/"
	cfg.areasPrev = ""
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	lowercased := strings.ToLower(trimmed)
	return strings.Fields(lowercased)
}

func commandExit(params ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(params ...string) error {
	commandNames := []string{"map", "mapb", "explore", "help", "exit"}

	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, cmd := range commandNames {
		fmt.Printf("%s - %s\n", commands[cmd].name, commands[cmd].description)
	}
	return nil
}

func commandMap(params ...string) error {
	res := pokemon_api.GetAreas(cfg.areasNext)
	cfg.areasNext = res.Next
	cfg.areasPrev = res.Previous
	for _, a := range res.Areas {
		fmt.Println(a.Name)
	}
	return nil
}

func commandMapB(params ...string) error {
	if cfg.areasPrev == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	res := pokemon_api.GetAreas(cfg.areasPrev)
	cfg.areasNext = res.Next
	cfg.areasPrev = res.Previous
	for _, a := range res.Areas {
		fmt.Println(a.Name)
	}
	return nil
}

func commandExplore(params ...string) error {
	if len(params) != 1 {
		fmt.Println("usage: explore <area-id>")
		return nil
	}
	area := params[0]
	info := pokemon_api.GetAreaInfo(area)
	fmt.Printf("Exploring %s...\n", area)
	for _, e := range info.PokemonEncounters {
		fmt.Println(e.Pokemon.Name)
	}
	cfg.area = area
	return nil
}

func commandCatch(params ...string) error {
	if len(params) != 1 {
		fmt.Println("usage: catch <pokemon-id>")
		return nil
	}
	pokemon := params[0]
	//info := pokemon_api.GetAreaInfo(cfg.area)
	pokemonInfo := pokemon_api.GetPokemonInfo(pokemon)
	fmt.Println("Throwing a pokeball at", pokemon, "...")
	luck := rand.Intn(700)
	if luck > pokemonInfo.BaseExperience {
		fmt.Println("You caught", pokemon)
	} else {
		fmt.Println(pokemon, "broke free")
	}
	return nil
}
