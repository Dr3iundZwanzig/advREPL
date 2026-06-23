package main

import (
	"bufio"
	"fmt"
	"os"
)

func guild(config *config) error {
	player := config.player
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("You enter the guild throu the main entrance and stand in the lobby.")
	fmt.Println("Your guild rank can be checked with the !player command outside the guild")
	fmt.Println("--Current options:\n -!exit: leave the guild\n -!quest: turn in your quest if ready or choose a new one")
	for {
		fmt.Print("Guild >>> ")
		reader.Scan()
		userInput := cleanInput(reader.Text())
		if len(userInput) == 0 {
			fmt.Println("No input entered.")
			continue
		}
		if userInput[0] == "!exit" {
			return nil
		}
		if userInput[0] == "!quest" {
			if questComplete(config) {
				fmt.Printf("Quest %v finished\n", player.CurrentQuests.CurrentQuest.Name)
				resetQuest(config)
				continue
			}
			if player.CurrentQuests.HasQuest {
				fmt.Println("---Current quest is unfinished")
				currentQuestData(config)
				continue
			}
			if !player.CurrentQuests.HasQuest {
				fmt.Println("You dont have a quest currently choose a new one now")
				chooseQuest(config)
				continue
			}
		}
		fmt.Println("wrong command")
	}
}
