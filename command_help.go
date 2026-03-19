package main

import (
	"fmt"
)



func commandHelp(config *Config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for name, desc := range getCommands(){
		fmt.Printf("%s: %s\n", name, desc.description)
		
	}
	return nil
}

