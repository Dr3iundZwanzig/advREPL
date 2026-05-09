package main

import (
	"bufio"
	"fmt"
	"os"
)

type room struct {
	name        string
	description string
	chance      int
}

func dungeon(cnf *config) error {
	cnf.player.InDungeon = true
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("As you arrived at the gates a guard stands in your way and demands a proof you belong to the guild.")
	switch rank := cnf.player.CurrentExperience.CurrentGuildRank; rank {
	case "Bronze":
		fmt.Println("You show the guard your Bronze badge and he lets you in but warns you not to go beyond the first floor.")
		fmt.Println("The bronze rank allows you to explore the first floor of the dungeon which expands over multible rooms.")
	default:
		return fmt.Errorf("You are not a guild member yet")
	}
	fmt.Println("Some commands are not avalible in the dungeon check !help for more infos")
	for {
		fmt.Println("You step further into the dungeon.")
		fmt.Print("Dungeon >>> ")
		reader.Scan()

		userInput := cleanInput(reader.Text())
		if len(userInput) == 0 {
			fmt.Println("No input entered.")
			continue
		}

		commandName := userInput[0]
		args := userInput[1:]
		command, exists := getCommands()[commandName]
		if !exists {
			println("Unknown command")
			continue
		}
		err := command.callback(cnf, args...)
		if command.name == "!leavedungeon" {
			return nil
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func getRooms() map[string]room {
	return map[string]room{
		"normalfight": {
			name:        "normalfight",
			description: "a normal fight",
			chance:      70,
		},
		"hardfight": {
			name:        "hardfight",
			description: "a hard fight",
			chance:      30,
		},
		"treasure": {
			name:        "treasure",
			description: "a treasure chest",
			chance:      10,
		},
	}
}
