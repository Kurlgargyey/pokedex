package main

import (
	"fmt"
	"io"
	"os"
	"pokedex/internal/pokemon_api"
	"strings"

	"github.com/chzyer/readline"
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
		description: "Exit the Pokédex",
		callback:    commandExit,
	}
	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
	commands["map"] = cliCommand{
		name:        "map",
		description: "Displays the next 20 locations in the world of pokémon",
		callback:    commandMap,
	}
	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays the previous 20 locations in the world of pokémon",
		callback:    commandMapB,
	}
	commands["explore"] = cliCommand{
		name:        "explore <area-id>",
		description: "Explores the given area",
		callback:    commandExplore,
	}
	commands["catch"] = cliCommand{
		name:        "catch <pokemon-id>",
		description: "Attempts to catch the given pokémon",
		callback:    commandCatch,
	}
	commands["inspect"] = cliCommand{
		name:        "inspect <pokemon-id>",
		description: "Inspect a Pokémon from your Pokédex",
		callback:    commandInspect,
	}

	cfg.areasNext = "https://pokeapi.co/api/v2/location-area/"
	cfg.areasPrev = ""
	cfg.pokedex = make(map[string]pokemon_api.Pokemon)
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	lowercased := strings.ToLower(trimmed)
	return strings.Fields(lowercased)
}

func commandExit(params ...string) error {
	fmt.Println("Closing the Pokédex... Goodbye!")
	os.Exit(0)
	return nil
}

func repl() {
	rl := initReadline()
	defer rl.Close()
	for {
		processCommand(parseCommands(rl))
	}
}

func initReadline() *readline.Instance {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:       "Pokedex >> ",
		HistoryLimit: 100,
	})
	if err != nil {
		fmt.Println("Error initializing readline:", err)
		os.Exit(1)
	}
	return rl
}

func parseCommands(rl *readline.Instance) []string {
	input, err := rl.Readline()
	if err == io.EOF {
		commands["exit"].callback()
	}
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	return cleanInput(input)
}

func processCommand(input []string) {
	fmt.Println("--------------------------------")
	if cmd, ok := commands[input[0]]; ok {
		cmd.callback(input[1:]...)
	} else {
		fmt.Println("Unknown command.")
	}
	fmt.Println("--------------------------------")
}
