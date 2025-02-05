package main

import (
	"fmt"
	"math/rand"
	"pokedex/internal/pokemon_api"
)

func commandCatch(params ...string) error {
	if len(params) != 1 {
		fmt.Printf("usage: %s\n", commands["catch"].name)
		return nil
	}
	pokemon := params[0]
	//info := pokemon_api.GetAreaInfo(cfg.area)
	pokemonInfo := pokemon_api.GetPokemonInfo(pokemon)
	fmt.Printf("Throwing a Pokéball at %s...\n", pokemonInfo.Name)
	luck := rand.Float64()
	chance := 1.0 / (1 + (float64(pokemonInfo.BaseExperience) / 90))
	if luck < chance {
		cfg.pokedex[pokemon] = pokemonInfo
		fmt.Println("You caught", pokemonInfo.Name)
	} else {
		fmt.Println(pokemonInfo.Name, "broke free")
	}
	return nil
}
