package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	story := loadStory("Act1.json")

	continues := false

	fmt.Println("Welcome to Adv")
	fmt.Println("Input your Name")
	reader.Scan()
	firstInput := reader.Text()
	charakter := createPlayer(firstInput)
	fmt.Println("Starting Adv...")
	for {
		if !continues {
			fmt.Println(story.ChapterSteps[charakter.currentStep].MainString)
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

		commandName := userInput[0]
		if strings.HasPrefix(commandName, "!") {
			continues = true
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
		println(choiceInput)
		charakter.currentStep = story.ChapterSteps[charakter.currentStep].TriggerChoice[choiceInput-1].ChoiceNextStep
		println(charakter.currentStep)
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
