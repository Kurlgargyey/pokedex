package pokemon_api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
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

func GetAreas(url string) areaResponse {
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
	var areas areaResponse
	err = json.Unmarshal(body, &areas)
	if err != nil {
		log.Fatal(err)
	}
	return areas
}
