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

		if config.player.AutoSave {
			commandSave(config, config.player.AutoSaveFileName, "autosave")
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
	items  ItemCollection
	story  Story
	quests QuestCollection
	enemys EnemyCollection
}

type cliCommand struct {
	name            string
	description     string
	canUseInDungeon bool
	callback        func(*config, ...string) error
}

type location struct {
	name        string
	description string
	callback    func(*config) error
}

func getLocations() map[string]location {
	return map[string]location{
		"consumableShop": {
			name:        "Consumable Shop",
			description: "A shop to buy items from",
			callback:    consumableShop,
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
			name:            "!help",
			description:     "Displays a help message",
			canUseInDungeon: true,
			callback:        commandHelp,
		},
		"!exit": {
			name:            "!exit",
			description:     "Exit the Programm",
			canUseInDungeon: false,
			callback:        commandExit,
		},
		"!player": {
			name:            "!player",
			description:     "Displayes player information",
			canUseInDungeon: true,
			callback:        commandPlayerInfo,
		},
		"!items": {
			name:            "!items",
			description:     "Displayes player items",
			canUseInDungeon: true,
			callback:        commandPlayerItems,
		},
		"!use": {
			name:            "!use",
			description:     "Use an item from your inventory. Usage: !use [itemID]",
			canUseInDungeon: true,
			callback:        commandUseItem,
		},
		"!choice": {
			name:            "!choice",
			description:     "Select a choice. Usage: !choice [choiceNumber]",
			canUseInDungeon: false,
			callback:        commandSelectChoice,
		},
		"!quest": {
			name:            "!quest",
			description:     "View current quest information",
			canUseInDungeon: true,
			callback:        commandQuestInfo,
		},
		"!go": {
			name:            "!go",
			description:     "Go to the given Location (cannot be uses in the dungeon)",
			canUseInDungeon: false,
			callback:        commandGo,
		},
		"!locations": {
			name:            "!locations",
			description:     "Show all the Locations",
			canUseInDungeon: false,
			callback:        commandLocations,
		},
		"!explore": {
			name:            "!explore",
			description:     "Explore the dungeon further (only works in the dungeon)",
			canUseInDungeon: true,
			callback:        commandExplore,
		},
		"!save": {
			name:            "!save",
			description:     "Save the current game state",
			canUseInDungeon: false,
			callback:        commandSave,
		},
		"!load": {
			name:            "!load",
			description:     "Load a saved game state",
			canUseInDungeon: false,
			callback:        commandLoad,
		},
		"!autosave": {
			name:            "!autosave",
			description:     "Enable autosave with \"-on\" or \"-of\" and a file name. If no additional arguments are given will show the state and the filename.\n Saves the game when you enter the city does not work in the dungeon",
			canUseInDungeon: false,
			callback:        commandAutoSave,
		},
		"!leavedungeon": {
			name:            "!leavedungeon",
			description:     "Leave the dungeon and return to the city",
			canUseInDungeon: true,
			callback:        commandLeaveDungeon,
		},
	}
}
