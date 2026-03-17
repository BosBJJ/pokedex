package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/BosBJJ/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name 		string
	description string
	callback 	func(*Config)error
}

type Config struct {
	Next 		*string
	Previous 	*string
}

func getCommands() map[string]cliCommand{
	return map[string]cliCommand {
	"exit": {
		name: 			"exit",
		description: 	"Exit the Pokedex",
		callback: 		commandExit,
	},
	"help": {
		name:  			"help",
		description: 	"Displays a help message",
		callback: 		commandHelp,
	},
	"map": {
		name: 			"map",
		description: 	"Displays 20 locations",
		callback: 		commandMap,
	},
	"mapb": {
		name: "map back",
		description: "Displays previous 20 locations",
		callback: commandMapb,
	},
}

}



func cleanInput(text string) []string{
	textLower := strings.ToLower(text)
	words := strings.Fields(textLower)
	return words
}

func startRepl(){
	scanner := bufio.NewScanner(os.Stdin)
	config := &Config{}
	for{
		fmt.Print("Pokedex > ")
		if !scanner.Scan(){
			return
		}
		scanToText := scanner.Text()
		cleaned := cleanInput(scanToText)
		if len(cleaned) < 1 {
			continue
		}
		commandName := cleaned[0]
		commands := getCommands()
		command, exists := commands[commandName]
		if exists == false {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(config)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func commandExit(config *Config) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

func commandMap(config *Config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if config.Next != nil {
		url = *config.Next
	}
	getLocs, err := pokeapi.GetLocations(url)
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

func commandMapb(config *Config) error {
	if config.Previous == nil {
		fmt.Print("You're already on the first page")
		return nil
	}
	url := *config.Previous
	getLocs, err := pokeapi.GetLocations(url)
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