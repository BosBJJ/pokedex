package main

import "fmt"

func commandMap(config *Config, args []string) error {
	getLocs, err := config.Client.GetLocations(config.Next)
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
    getLocs, err := config.Client.GetLocations(config.Previous)
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
