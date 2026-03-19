package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


func (c *PokeClient) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name
	res, err := http.Get(url)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Error: %w", err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	pokemonRes := Pokemon{}
	err = json.Unmarshal(body, &pokemonRes)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Pokemon results could not be unmarshaled: %w", err)
	}
	return pokemonRes, nil
}
