package main

import "fmt"



func commandPokeDex(config *Config, args []string) error {
	pokemon := config.Captured
	if len(pokemon) < 1 {
		fmt.Println("You do not have any Pokemon, try catching one with the catch command")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, name := range pokemon {
		fmt.Printf(" - %s\n", name.Name)
	}
	return nil
}