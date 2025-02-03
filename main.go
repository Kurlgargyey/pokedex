package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >> ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		fmt.Println("--------------------------------")
		if cmd, ok := commands[cleanedInput[0]]; ok {
			cmd.callback()
		} else {
			fmt.Println("Unknown command.")
		}
		fmt.Println("--------------------------------")
	}
}
