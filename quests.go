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
	QuestID         int    `json:"questID"`
	Name            string `json:"name"`
	Objective       string `json:"objective"`
	Amount          int    `json:"amount"`
	Description     string `json:"description"`
	GoldReward      int    `json:"goldReward"`
	Repeatable      bool   `json:"repeatable"`
	GuildExperience int    `json:"guildExperience"`
	ItemRewards     []struct {
		ItemID int `json:"itemID"`
		Amount int `json:"amount"`
	} `json:"itemRewards"`
}

type QuestCollection struct {
	KillQuests  map[string]Quest `json:"killQuests"`
	FetchQuests map[string]Quest `json:"fetchQuests"`
}

func loadQuests(file []byte) QuestCollection {
	var collection QuestCollection
	err := json.Unmarshal(file, &collection)
	if err != nil {
		log.Fatal(fmt.Errorf("Error creating struct from quests.json: %v", err))
	}
	return collection
}

func (qc QuestCollection) GetQuestByID(id int) (Quest, bool) {
	for _, q := range qc.KillQuests {
		if q.QuestID == id {
			return q, true
		}
	}
	for _, q := range qc.FetchQuests {
		if q.QuestID == id {
			return q, true
		}
	}
	return Quest{}, false
}

func displayQuests(config *config) {
	killQuests := config.quests.KillQuests
	fetchQuests := config.quests.FetchQuests
	fmt.Println("Kill Quests:")
	for _, quest := range killQuests {
		fmt.Printf("-ID: %v Name: %v\n", quest.QuestID, quest.Name)
	}
	fmt.Println("Fetch Quests:")
	for _, quest := range fetchQuests {
		fmt.Printf("-ID: %v Name: %v\n", quest.QuestID, quest.Name)
	}
}

func chooseQuest(config *config) {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("You can only have one quest active at a time.")
	fmt.Println("Enter the ID of the quest you want to take: ")

	for {
		displayQuests(config)
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
		if config.player.CurrentQuests.HasQuest == true {
			fmt.Println("You already have an active quest. You can only have one active quest at a time.")
			break
		}
		if quest, ok := config.quests.GetQuestByID(questID); ok {
			config.player.CurrentQuests.HasQuest = true
			config.player.CurrentQuests.CurrentQuest = quest
			config.player.CurrentQuests.Progress = 0
			fmt.Println("You have taken the quest:", quest.Name)
		} else {
			fmt.Println("Quest with ID", questID, "does not exist.")
			continue
		}
		break
	}
}
func questComplete(config *config) bool {
	player := config.player
	if player.CurrentQuests.HasQuest && player.CurrentQuests.Progress >= player.CurrentQuests.CurrentQuest.Amount {
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
		return
	}
	quest := config.player.CurrentQuests.CurrentQuest
	fmt.Println("Current quest information:")
	fmt.Printf("Quest name: %v\n", quest.Name)
	fmt.Printf("Quest description: %v\n", quest.Description)
	fmt.Printf("Rewards:\n -Gold: %v\n -Guild experience: %v\n", quest.GoldReward, quest.GuildExperience)
	fmt.Println(" -Items: ")
	for _, item := range quest.ItemRewards {
		if itemObj, ok := config.items.GetItemByID(item.ItemID); ok {
			fmt.Printf("  -%v / Amount: %v\n", itemObj.Name, item.Amount)
		}
	}
	fmt.Printf("Quest progress: %v/%v\n", config.player.CurrentQuests.Progress, quest.Amount)
}
