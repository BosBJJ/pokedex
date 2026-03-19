package pokeapi

import (
	"net/http"
	"fmt"
	"io"
	"encoding/json"
)


func (c *PokeClient) GetLocations(pageURL *string) (LocationArea, error){
	url := baseURL + "/location-area"
	if pageURL != nil {
        url = *pageURL
    }
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