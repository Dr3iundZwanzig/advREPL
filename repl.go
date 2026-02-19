package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	story := loadStory("Act1.json")

	continues := false // used to prevent printing the main string again after a command is executed or an unknown input is given

	charakter := createPlayer()
	items := loadItems()
	itemMap := make(map[int]Item)
	for _, item := range items.Items {
		itemMap[item.ItemID] = item
	}
	fmt.Println("Welcome to Adv")
	fmt.Println("Starting Adv...")
	fmt.Println("Press enter to continue...")
	reader.Scan()

	for {
		if !continues {
			fmt.Println(story.ChapterSteps[charakter.currentStep].MainString)
			if story.ChapterSteps[charakter.currentStep].HasEvent {
				for _, event := range story.ChapterSteps[charakter.currentStep].Events {
					triggerEvent(event, &charakter, itemMap)
				}
			}
			if story.ChapterSteps[charakter.currentStep].HasChoice {
				fmt.Println("Your choices:")
				for _, choice := range story.ChapterSteps[charakter.currentStep].TriggerChoice {
					fmt.Println(choice.ChoiceText)
				}
			}
		} else {
			continues = false
		}

		fmt.Print("Adv >>> ")
		reader.Scan()

		userInput := cleanInput(reader.Text())
		if len(userInput) == 0 {
			continue
		}
		// command input
		commandName := userInput[0]
		if strings.HasPrefix(commandName, "!") {
			continues = true
			err := executeCommand(userInput, &charakter)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
		// choice input
		choiceInput, err := strconv.Atoi(userInput[0])
		if err != nil {
			continues = true
			println("Unknown input")
			continue
		}
		if choiceInput < 1 || choiceInput > len(story.ChapterSteps[charakter.currentStep].TriggerChoice) {
			continues = true
			println("Unknown input")
			continue
		}
		charakter.currentStep = story.ChapterSteps[charakter.currentStep].TriggerChoice[choiceInput-1].ChoiceNextStep

	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
	playerCall  func(player) error
	playerInput func(*player, int) error
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
		"!items": {
			name:        "!items",
			description: "Displayes player items",
			playerCall:  commandPlayerItems,
		},
		"!use": {
			name:        "!use",
			description: "Use an item from your inventory. Usage: !use [itemID]",
			playerInput: commandUseItem,
		},
	}
}

func executeCommand(playerInput []string, p *player) error {
	commandName := playerInput[0]
	command, exists := getCommands()[commandName]
	if !exists {
		return fmt.Errorf("Unknown command: %s", commandName)
	}
	if command.playerCall != nil {
		return command.playerCall(*p)
	}
	if command.playerInput != nil {
		if len(playerInput) < 2 {
			return fmt.Errorf("Missing item ID for command %s", commandName)
		}
		itemID, err := strconv.Atoi(playerInput[1])
		if err != nil {
			return fmt.Errorf("Invalid item ID: %s", playerInput[1])
		}
		return command.playerInput(p, itemID)
	}
	if command.callback != nil {
		return command.callback()
	}
	return nil
}
