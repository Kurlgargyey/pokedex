package pokemon_api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"pokedex/internal/pokecache"
	"time"
)

type area struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type areaResponse struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Areas    []area `json:"results"`
}

type areaInfo struct {
	PokemonEncounters []Encounter `json:"pokemon_encounters"`
}
type Encounter struct {
	Pokemon pokemonInfo `json:"pokemon"`
}
type pokemonInfo struct {
	Name string `json:"name"`
}

type Pokemon struct {
	Name           string        `json:"name"`
	Height         int           `json:"height"`
	BaseExperience int           `json:"base_experience"`
	Weight         int           `json:"weight"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
}

type PokemonStat struct {
	Value int `json:"base_stat"`
	Stat  struct {
		Name string `json:"name"`
	} `json:"stat"`
}

type PokemonType struct {
	Slot int `json:"slot"`
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

var cache *pokecache.Cache

func init() {
	cache = pokecache.NewCache(5 * time.Minute)
}

func GetAreas(url string) areaResponse {
	if res, ok := cache.Get(url); ok {
		return unmarshalAreas(res)
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body := readHTMLBody(res)

	cache.Add(url, body)
	return unmarshalAreas(body)
}

func unmarshalAreas(body []byte) areaResponse {
	var areas areaResponse
	err := json.Unmarshal(body, &areas)
	if err != nil {
		log.Fatal(err)
	}
	return areas
}

func GetAreaInfo(area string) areaInfo {
	if res, ok := cache.Get("https://pokeapi.co/api/v2/location-area/" + area); ok {
		info := unmarshalAreaInfo(res)
		return info
	}

	res, err := http.Get("https://pokeapi.co/api/v2/location-area/" + area)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body := readHTMLBody(res)

	cache.Add("https://pokeapi.co/api/v2/location-area/"+area, body)
	return unmarshalAreaInfo(body)
}

func unmarshalAreaInfo(body []byte) areaInfo {
	var info areaInfo
	err := json.Unmarshal(body, &info)
	if err != nil {
		log.Fatal(err)
	}
	return info
}

func GetPokemonInfo(pokemon string) Pokemon {
	if res, ok := cache.Get("https://pokeapi.co/api/v2/pokemon/" + pokemon); ok {
		info := unmarshalPokemon(res)
		return info
	}

	res, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + pokemon)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body := readHTMLBody(res)

	cache.Add("https://pokeapi.co/api/v2/pokemon/"+pokemon, body)
	return unmarshalPokemon(body)
}

func unmarshalPokemon(body []byte) Pokemon {
	var info Pokemon
	err := json.Unmarshal(body, &info)
	if err != nil {
		log.Fatal(err)
	}
	return info
}

func readHTMLBody(res *http.Response) []byte {
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		log.Fatalf("Failed to fetch pokemon info with status code %d: %s", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	return body
}
