package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
	playerCall  func(player) error
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to Adv")
	fmt.Println("Input your Name")
	reader.Scan()
	firstInput := reader.Text()
	charakter := createPlayer(firstInput)
	for {
		fmt.Print("Adv >>> ")
		reader.Scan()

		userInput := cleanInput(reader.Text())
		if len(userInput) == 0 {
			continue
		}

		commandName := userInput[0]
		if strings.HasPrefix(commandName, "!") {
			command, exists := getCommands()[commandName]
			if command.name == "!player" {
				err := command.playerCall(charakter)
				if err != nil {
					fmt.Println(err)
				}
				continue
			}
			if exists {
				err := command.callback()
				if err != nil {
					fmt.Println(err)
				}
				continue
			} else {
				fmt.Println("Unknown command")
				continue
			}
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"!help": {
			name:        "!help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"!exit": {
			name:        "!exit",
			description: "Exit the Programm",
			callback:    commandExit,
		},
		"!player": {
			name:        "!player",
			description: "Displayes player information",
			playerCall:  commandPlayerInfo,
		},
	}
}
