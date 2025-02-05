package main

import "fmt"

func commandPokedex(params ...string) error {
	dexString := "Your Pokédex:\n"
	for _, pokemon := range cfg.pokedex {
		dexString += fmt.Sprintf("\t- %s\n", pokemon.Name)
	}
	if len(cfg.pokedex) == 0 {
		dexString += "The Pokédex is still empty...\n"
	}
	fmt.Print(dexString)
	return nil
}
