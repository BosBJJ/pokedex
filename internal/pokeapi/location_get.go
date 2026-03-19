package pokeapi

import (
	"net/http"
	"fmt"
	"io"
	"encoding/json"
)

func (c *PokeClient) ExploreLocation(name string) (ExploreArea, error) {
	url := baseURL + "/location-area/" + name
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