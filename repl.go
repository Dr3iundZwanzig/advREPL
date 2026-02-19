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
			command, exists := getCommands()[commandName]
			if command.name == "!player" || command.name == "!items" {
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
	}
}
