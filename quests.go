package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Quest struct {
	QuestID          int    `json:"questID"`
	QuestName        string `json:"questName"`
	QuestType        string `json:"questType"`
	QuestAmount      int    `json:"questAmount"`
	QuestDescription string `json:"questDescription"`
	QuestGoldReward  int    `json:"questGoldReward"`
	Repeatable       bool   `json:"repeatable"`
	QuestItemRewards []struct {
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
