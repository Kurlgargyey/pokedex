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
		fmt.Printf("Your command was: %v\n", cleanedInput[0])
		if cleanedInput[0] == "quit" {
			fmt.Println("Quitting...")
			break
		}
	}
}