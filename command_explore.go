package main

import (
	"fmt"
	"pokedex/internal/pokemon_api"
)

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
