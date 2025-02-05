package main

import (
	"fmt"
	"strings"
)

func commandInspect(params ...string) error {
	if len(params) != 1 {
		fmt.Printf("usage: %s\n", commands["inspect"].name)
		return nil
	}
	pokemon := params[0]
	if pokemonInfo, ok := cfg.pokedex[pokemon]; ok {
		info := []string{}
		nameString := fmt.Sprintf("Name: %s", pokemonInfo.Name)
		info = append(info, nameString)
		heightString := fmt.Sprintf("Height: %d", pokemonInfo.Height)
		info = append(info, heightString)
		weightString := fmt.Sprintf("Weight: %d", pokemonInfo.Weight)
		info = append(info, weightString)
		statString := "Stats:\n"
		for _, stat := range pokemonInfo.Stats {
			statString += fmt.Sprintf("\t-%s: %d\n", stat.Stat.Name, stat.Value)
		}
		info = append(info, statString)
		typeString := "Types:\n"
		for _, pokeType := range pokemonInfo.Types {
			typeString += fmt.Sprintf("\t- %s\n", pokeType.Type.Name)
		}
		info = append(info, typeString)
		fmt.Print(strings.Join(info, "\n"))
	} else {
		fmt.Println("You have not caught that Pokémon")
	}
	return nil
}
