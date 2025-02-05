package main

import (
	"fmt"
	"io"
	"os"

	"github.com/chzyer/readline"
)

func main() {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:       "Pokedex >> ",
		HistoryLimit: 100,
	})
	if err != nil {
		fmt.Println("Error initializing readline:", err)
		os.Exit(1)
	}
	defer rl.Close()
	for {
		input, err := rl.Readline()
		if err == io.EOF {
			commands["exit"].callback()
		}
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}
		cleanedInput := cleanInput(input)
		fmt.Println("--------------------------------")
		if cmd, ok := commands[cleanedInput[0]]; ok {
			cmd.callback(cleanedInput[1:]...)
		} else {
			fmt.Println("Unknown command.")
		}
		fmt.Println("--------------------------------")
	}
}
