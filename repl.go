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
		currentStep := config.story.ChapterSteps[config.player.CurrentStep]
		if !continues {
			continueStory(config)
			if !currentStep.HasChoice && currentStep.NextStep != nil {
				config.player.CurrentStep = *currentStep.NextStep
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
			println("Unknown command: " + commandName + ", type !help for a list of commands")
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
	player Player
	items  map[int]Item
	story  Story
	quests map[int]Quest
	enemys EnemyCollection
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type location struct {
	name        string
	description string
	callback    func(*config) error
}

func getLocations() map[string]location {
	return map[string]location{
		"shop": {
			name:        "shop",
			description: "A shop to buy items from",
			callback:    regularShop,
		},
		"dungeon": {
			name:        "dungeon",
			description: "The dungeons inside the tree",
			callback:    dungeon,
		},
		"statue": {
			name:        "statue",
			description: "A mysterious statue that allows the player to level up",
			callback:    statue,
		},
		"guild": {
			name:        "guild",
			description: "The Guild",
			callback:    guild,
		},
	}
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
		"!go": {
			name:        "!go",
			description: "Go to the given Location (cannot be uses in the dungeon)",
			callback:    commandGo,
		},
		"!locations": {
			name:        "!locations",
			description: "Show all the Locations",
			callback:    commandLocations,
		},
		"!explore": {
			name:        "!explore",
			description: "Explore the dungeon further (only works in the dungeon)",
			callback:    commandExplore,
		},
		"!save": {
			name:        "!save",
			description: "Save the current game state",
			callback:    commandSave,
		},
		"!load": {
			name:        "!load",
			description: "Load a saved game state",
			callback:    commandLoad,
		},
		"!leavedungeon": {
			name:        "!leavedungeon",
			description: "Leave the dungeon and return to the city",
			callback:    commandLeaveDungeon,
		},
	}
}
