package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func commandExit(config *config, _ ...string) error {
	fmt.Println("Closing the Adv!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config, _ ...string) error {
	fmt.Println("---")
	fmt.Println("Welcome to the Adv help page!")
	fmt.Println("Usage:")
	fmt.Println("Need to use ! in front of commands")
	fmt.Println("To selcet a choice type the number after !choice")
	fmt.Println("Some commands are not avalible while in the dungeon check the description below")
	fmt.Println("---")
	fmt.Println("Commands:")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandPlayerInfo(config *config, _ ...string) error {
	p := config.player
	fmt.Println("---")
	fmt.Println("Player Information:")
	fmt.Printf("Name: %v\n", p.PlayerName)
	fmt.Printf("Health: %v/%v\n", p.CurrentHealth, p.MaxHealth)
	fmt.Printf("Mana: %v/%v\n", p.CurrentMana, p.MaxMana)
	fmt.Printf("Armour: %v\n", p.CurrentArmour)
	fmt.Printf("Gold: %v\n", p.Gold)
	fmt.Printf("Level: %v\n", p.CurrentExperience.CurrentLevel)
	fmt.Printf("XP: %v/%v\n", p.CurrentExperience.CurrentXP, p.CurrentExperience.NextLevelXP)
	fmt.Printf("Guild Rank: %v\n", p.CurrentExperience.CurrentGuildRank)
	fmt.Printf("Guild XP: %v/%v\n", p.CurrentExperience.CurrentGuildXP, p.CurrentExperience.NextGuildLevelXP)
	fmt.Println("---")

	return nil
}

func commandPlayerItems(config *config, _ ...string) error {
	p := config.player
	if len(p.Items) == 0 {
		return fmt.Errorf("You have no items")
	}
	fmt.Println("---")
	fmt.Println("Player Items:")
	for itemID, items := range p.Items {
		fmt.Printf("ID:%v -%v (Amount: %v)\n", itemID, items.Item.Name, items.Amount)
	}
	fmt.Println("---")
	return nil
}

func commandUseItem(config *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("Usage: !use [itemID]")
	}
	itemID, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("Invalid item ID")
	}
	config.player.useItem(itemID)
	return nil
}

func commandSelectChoice(config *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("missing choice number")
	}
	choiceNumber, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid choice number")
	}
	currentStep := config.story.ChapterSteps[config.player.CurrentStep]
	if !currentStep.HasChoice {
		return fmt.Errorf("no choices available")
	}
	if choiceNumber < 1 || choiceNumber > len(currentStep.TriggerChoice) {
		return fmt.Errorf("invalid choice number")
	}
	config.player.CurrentStep = currentStep.TriggerChoice[choiceNumber-1].ChoiceNextStep
	return nil
}

func commandQuestInfo(config *config, _ ...string) error {
	if !config.player.CurrentQuests.HasQuest {
		return fmt.Errorf("You have no active quest.")
	}
	q := config.player.CurrentQuests.CurrentQuest
	fmt.Println("---")
	fmt.Println("Active quest:")
	fmt.Printf("Name: %v\n", q.Name)
	fmt.Printf("Description: %v\n", q.Description)
	fmt.Printf("Gold Reward: %v\n", q.GoldReward)
	fmt.Printf("Guild Experience Reward: %v\n", q.GuildExperience)
	fmt.Println("Item Reward:")
	for _, reward := range q.ItemRewards {
		if item, ok := config.items.GetItemByID(reward.ItemID); ok {
			fmt.Printf("  - %v (Amount: %v)\n", item.Name, reward.Amount)
		}
	}
	fmt.Println("---")
	return nil
}

func commandGo(config *config, args ...string) error {
	if config.player.CurrentChapter == 1 && config.player.CurrentStep < 4 {
		return fmt.Errorf("Command not unlocked")
	}
	if len(args) < 1 {
		return fmt.Errorf("missing location name")
	}
	locationName := args[0]
	location, exists := getLocations()[locationName]
	if !exists {
		return fmt.Errorf("Unknown Location")
	}
	err := location.callback(config)
	if err != nil {
		return err
	}
	return nil
}

func commandLocations(config *config, args ...string) error {
	if config.player.CurrentChapter == 1 && config.player.CurrentStep < 0 {
		return fmt.Errorf("Command not unlocked")
	}
	locations := getLocations()
	fmt.Println("Locations:")
	fmt.Println("---")
	for _, location := range locations {
		fmt.Printf("Name: %v ", location.name)
		fmt.Printf("/ %v\n", location.description)
	}
	fmt.Println("---")
	return nil
}

func commandExplore(config *config, args ...string) error {
	if !config.player.InDungeon {
		return fmt.Errorf("You are not in the dungeon")
	}
	rooms := getRooms()
	highestChance := 0
	for _, room := range rooms {
		if room.chance > highestChance {
			highestChance = room.chance
		}
	}
	randomNumber := rand.Intn(highestChance)
	currentRoomChance := highestChance + 1
	choosenRoom := room{}

	for _, room := range rooms {
		fmt.Println(randomNumber)
		if randomNumber <= room.chance && room.chance < currentRoomChance {
			currentRoomChance = room.chance
			choosenRoom = room
		}
	}
	err := triggerDungeonEvent(choosenRoom, config)
	return err
}

func commandSave(config *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("missing save file name")
	}
	saveFileName := args[0]
	file, err := os.Create(saveFileName + ".json")
	if err != nil {
		return fmt.Errorf("error creating file")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(config.player)
	if err != nil {
		return fmt.Errorf("error encoding config to json")
	}
	fmt.Println("Game saved successfully to", saveFileName+".json")
	return nil
}

func commandLoad(config *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("missing save file name")
	}
	saveFileName := args[0]
	file, err := os.Open(saveFileName + ".json")
	if err != nil {
		return fmt.Errorf("error opening file")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config.player)
	if err != nil {
		return fmt.Errorf("error decoding config from json")
	}
	fmt.Println("Game loaded successfully from", saveFileName+".json")
	return nil
}

func commandLeaveDungeon(config *config, args ...string) error {
	if !config.player.InDungeon {
		return fmt.Errorf("You are not in the dungeon")
	}
	config.player.InDungeon = false
	fmt.Println("You leave the dungeon and return to the city.")
	return nil
}
