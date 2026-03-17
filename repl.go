package main
import ("strings" 
		"fmt"
		"bufio"
		"os"
)

type cliCommand struct {
	name string
	description string
	callback func()error
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
}

}



func cleanInput(text string) []string{
	textLower := strings.ToLower(text)
	words := strings.Fields(textLower)
	return words
}

func startRepl(){
	scanner := bufio.NewScanner(os.Stdin)
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
		err := command.callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}