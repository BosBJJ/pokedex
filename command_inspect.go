package main

import "fmt"

func commandInspect(config *Config, args []string) error {
	if len(args) < 2 {
		fmt.Println(`Please type the Pokemon name after "inspect"`)
		return nil
	}
	pokemon, ok := config.Captured[args[1]]
	if !ok {
		fmt.Print("You have not caught that Pokemon yet\n")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, elem := range pokemon.Types {
		fmt.Printf("  - %s\n", elem.Type.Name)
	}
	
	return nil
}
