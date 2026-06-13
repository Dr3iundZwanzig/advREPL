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
	QuestObjective       string `json:"questObjective"`
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

func loadQuests(file []byte) QuestCollection {
	var quests QuestCollection
	err := json.Unmarshal(file, &quests)
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
		if config.player.CurrentQuests.HasQuest == true {
			fmt.Println("You already have an active quest. You can only have one active quest at a time.")
			break
		}
		config.player.CurrentQuests.HasQuest = true
		config.player.CurrentQuests.CurrentQuest = config.quests[questID]
		config.player.CurrentQuests.Progress = 0
		fmt.Println("You have taken the quest:", config.quests[questID].QuestName)
		break
	}
}
func questComplete(config *config) bool {
	player := config.player
	if player.CurrentQuests.HasQuest && player.CurrentQuests.Progress >= player.CurrentQuests.CurrentQuest.QuestAmount {
		return true
	}
	return false
}

func resetQuest(config *config) {
	config.player.CurrentQuests = PlayerQuest{
		HasQuest:     false,
		CurrentQuest: Quest{},
		Progress:     0,
	}
}

func currentQuestData(config *config) {
	if !config.player.CurrentQuests.HasQuest {
		fmt.Println("You dont have any quest active")
	}
	quest := config.player.CurrentQuests.CurrentQuest
	fmt.Println("Current quest information:")
	fmt.Printf("Quest name: %v\n", quest.QuestName)
	fmt.Printf("Quest description: %v\n", quest.QuestDescription)
	fmt.Printf("Rewards:\n -Gold: %v\n -Guild experience: %v\n", quest.QuestGoldReward, quest.QuestGuildExperience)
	fmt.Println(" -Items: ")
	for _, item := range quest.QuestItemRewards {
		fmt.Printf("  -%v / Amount: %v\n", config.items[item.ItemID].ItemName, item.Amount)
	}
	fmt.Printf("Quest progress: %v/%v\n", config.player.CurrentQuests.Progress, quest.QuestAmount)
}
