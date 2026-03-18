package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/BosBJJ/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name 		string
	description string
	callback 	func(*Config, []string)error
}

type Config struct {
	Next 		*string
	Previous 	*string
	Client 		pokeapi.PokeClient
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
		name: 			"map back",
		description: 	"Displays previous 20 locations",
		callback: 		commandMapb,
	},
	"explore": {
		name: 			"explore",
		description: 	"Displays pokemon in the area",
		callback: 		commandExplore,
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
	config := &Config{
		Client: pokeapi.NewPokeClient(5*time.Second, 5*time.Minute),
	}
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
		err := command.callback(config, cleaned)
		if err != nil {
			fmt.Println(err)
		}
	}
}
