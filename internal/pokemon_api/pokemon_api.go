package pokemon_api

import (
	"encoding/json"
	"fmt"
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

var cache *pokecache.Cache

func init() {
	cache = pokecache.NewCache(5 * time.Minute)
}

func GetAreas(url string) areaResponse {
	if res, ok := cache.Get(url); ok {
		fmt.Println("I found this in the cache!")
		return unmarshalAreas(res)
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		log.Fatalf("Failed to fetch areas with status code %d: %s", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	cache.Add(url, body)
	fmt.Println("The cache contains", len(cache.Entries), "entries.")
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