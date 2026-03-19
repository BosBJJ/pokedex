package main

import (
	"fmt"
)

func commandExplore(config *Config, args []string) error {
	explore, err := config.Client.ExploreLocation(args[1])
	if err != nil {
		return fmt.Errorf("Couldn't explore area due to error: %w", err)
	}
	for _, pokemon := range explore.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}