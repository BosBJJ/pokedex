package main

import (
	"fmt"
	"os"
)

func commandExit(config *Config, args []string) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

func commandMap(config *Config, args []string) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if config.Next != nil {
		url = *config.Next
	}
	getLocs, err := config.Client.GetLocations(url)
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}
	config.Next = getLocs.Next
	config.Previous = getLocs.Previous
	for _, loc := range getLocs.Results{
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(config *Config, args []string) error {
	if config.Previous == nil {
		fmt.Print("You're already on the first page")
		return nil
	}
	url := *config.Previous
	getLocs, err := config.Client.GetLocations(url)
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}
	config.Next = getLocs.Next
	config.Previous = getLocs.Previous
	for _, loc := range getLocs.Results{
		fmt.Println(loc.Name)
	}
	return nil
}

func commandExplore(config *Config, args []string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + args[1]
	explore, err := config.Client.ExploreLocation(url)
	if err != nil {
		return fmt.Errorf("Couldn't explore area due to error: %w", err)
	}
	for _, pokemon := range explore.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}