package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(config *config) {
	reader := bufio.NewScanner(os.Stdin)

	continues := false // used to prevent printing the main string again after a command is executed or an unknown input is given

	fmt.Println("Welcome to Adv")
	fmt.Println("Type !help for a list of commands")
	fmt.Println("Press enter to continue...")
	reader.Scan()

	for {
		currentStep := config.story.ChapterSteps[config.player.currentStep]
		if !continues {
			continueStory(config)
			if !currentStep.HasChoice && currentStep.NextStep != nil {
				config.player.currentStep = *currentStep.NextStep
				continue
			}
		}

		fmt.Print("Adv >>> ")
		reader.Scan()

		userInput := cleanInput(reader.Text())
		if len(userInput) == 0 {
			fmt.Println("No input entered.")
			continues = true
			continue
		}
		// command input
		commandName := userInput[0]
		args := userInput[1:]
		continues = true
		command, exists := getCommands()[commandName]
		if !exists {
			println("Unknown command")
			continue
		}
		err := command.callback(config, args...)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if commandName == "!choice" {
			continues = false
		}
		if exists {
			continue
		}

		fmt.Println("Unknown input, type !help for a list of commands")
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type config struct {
	player player
	items  map[int]Item
	story  Story
	quests map[int]Quest
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
			callback:    commandPlayerInfo,
		},
		"!items": {
			name:        "!items",
			description: "Displayes player items",
			callback:    commandPlayerItems,
		},
		"!use": {
			name:        "!use",
			description: "Use an item from your inventory. Usage: !use [itemID]",
			callback:    commandUseItem,
		},
		"!choice": {
			name:        "!choice",
			description: "Select a choice. Usage: !choice [choiceNumber]",
			callback:    commandSelectChoice,
		},
		"!quest": {
			name:        "!quest",
			description: "View current quest information",
			callback:    commandQuestInfo,
		},
	}
}
