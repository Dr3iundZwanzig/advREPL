package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Quest struct {
	QuestID              int    `json:"questID"`
	QuestName            string `json:"questName"`
	QuestType            string `json:"questType"`
	QuestAmount          int    `json:"questAmount"`
	QuestDescription     string `json:"questDescription"`
	QuestGoldReward      int    `json:"questGoldReward"`
	Repeatable           bool   `json:"repeatable"`
	QuestGuildExperience int    `json:"questGuildExperience"`
	QuestItemRewards     []struct {
		ItemID int `json:"itemID"`
		Amount int `json:"amount"`
	} `json:"questItemRewards"`
}

type QuestCollection struct {
	Quests []Quest `json:"quests"`
}

func loadQuests() QuestCollection {
	file, err := os.ReadFile("quests.json")
	if err != nil {
		log.Fatal(fmt.Errorf("Error opening file quests.json: %v", err))
	}

	var quests QuestCollection
	err = json.Unmarshal(file, &quests)
	if err != nil {
		log.Fatal(fmt.Errorf("Error creating struct from quests.json: %v", err))
	}
	return quests
}

func chooseQuest(questIDs []int, config *config) {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("You can only have one quest active at a time.")
	fmt.Println("Enter the ID of the quest you want to take: ")
	for _, questID := range questIDs {
		if quest, ok := config.quests[questID]; ok {
			fmt.Printf("ID: %v - %v\n", quest.QuestID, quest.QuestName)
		} else {
			fmt.Printf("Quest with ID %v not found in config\n", questID)
		}
	}
	for {
		fmt.Print("Adv >>> ")
		reader.Scan()
		userInput := cleanInput(reader.Text())
		if len(userInput) == 0 {
			fmt.Println("No input entered.")
			continue
		}
		questID, err := strconv.Atoi(userInput[0])
		if err != nil {
			fmt.Println("Invalid quest ID entered.")
			continue
		}
		if _, ok := config.quests[questID]; !ok {
			fmt.Println("Quest with ID", questID, "does not exist.")
			continue
		}
		if config.player.currentQuests.hasQuest == true {
			fmt.Println("You already have an active quest. You can only have one active quest at a time.")
			break
		}
		config.player.currentQuests.hasQuest = true
		config.player.currentQuests.currentQuest = config.quests[questID]
		fmt.Println("You have taken the quest:", config.quests[questID].QuestName)
		break
	}
}
