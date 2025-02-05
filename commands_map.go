package main

import (
	"fmt"
	"pokedex/internal/pokemon_api"
)

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
