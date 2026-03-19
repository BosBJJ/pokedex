package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *Config, args []string) error {
	pokeData, err := config.Client.GetPokemon(args[1])
	if err != nil {
		return fmt.Errorf("Couldn't fetch Pokemon data due to error: %w", err)
	}
	pokemon := pokeData.Name
	chance := pokeData.BaseExperience/10
	if chance < 1 {
		chance = 1
	}
	throwMessage := "Throwing a Pokeball at " + pokemon + "...\n"
	fmt.Print(throwMessage)
	roll := rand.Intn(chance)
	if roll == 0 {
		config.Captured[pokemon] = pokeData
		fmt.Printf("%s was caught!\n", pokemon)
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}
	return nil
}