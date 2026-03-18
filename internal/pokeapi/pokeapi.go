package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *PokeClient) GetLocations(url string) (LocationArea, error){
	res, err := http.Get(url)
	if err != nil {
		return LocationArea{}, fmt.Errorf("Error: %w", err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	locAreaRes := LocationArea{}
	err = json.Unmarshal(body, &locAreaRes)
	if err != nil {
		return LocationArea{}, fmt.Errorf("Location Area Results could not be unmarshaled: %w", err)
	}
	return locAreaRes, nil

}

func (c *PokeClient) ExploreLocation(url string) (ExploreArea, error) {
	res, err := http.Get(url)
	if err != nil {
		return ExploreArea{}, fmt.Errorf("Error: %w", err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	exploreRes := ExploreArea{}
	err = json.Unmarshal(body, &exploreRes)
	if err != nil {
		return ExploreArea{}, fmt.Errorf("Explore results could not be unmarshaled: %w", err)
	}
	return exploreRes, nil
}

type LocationArea struct {
	Count int 			`json:"count"`
	Next *string 		`json:"next"`
	Previous *string 	`json:"previous"`
	Results []struct {
		Name string 	`json:"name"`
		URL string 		`json:"url"`
	} `json:"results"`
}

type ExploreArea struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string  	`json:"name"`
		} 					`json:"pokemon"`
	} 						`json:"pokemon_encounters"`
}
