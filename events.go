package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func triggerStoryEvent(event Event, config *config) error {

	if _, ok := config.player.Events[event.EventName]; ok {
		return nil
	}
	p := &config.player
	switch event.EventName {
	case "Guild Registration":
		p.Events[event.EventName] = event
		p.CurrentExperience.CurrentGuildRank = "Bronze"
		p.addItem(config.items[1], 1)
		p.addItem(config.items[3], 1)
		fmt.Println(event.EventDescription)
		namePlayer(p)
	case "Old man Shop":
		p.Events[event.EventName] = event
		fmt.Println(event.EventDescription)
		err := regularShop(config)
		if err != nil {
			return err
		}
	case "Get Quest":
		quests := []int{1, 2, 3}
		chooseQuest(quests, config)
	case "Open World":
		p.Events[event.EventName] = event
		fmt.Println("-------------------------------")
		fmt.Println("You are now free to explore the world! Type !help for a list of commands and !playerinfo to see your stats.\nTo continue with the story reach Level 5 and the silver guild rank by doing quests.")
		fmt.Println("-------------------------------")
	}
	return nil
}

func namePlayer(p *Player) {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	reader.Scan()
	p.PlayerName = reader.Text()
}

func triggerDungeonEvent(room room, config *config) error {
	switch room.name {
	case "normalfight":
		fmt.Println(room.description)
		return nil
	case "hardfight":
		fmt.Println(room.description)
		return nil
	case "treasure":
		err := treasureEvent(config)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func treasureEvent(config *config) error {
	fmt.Println("You found a treasure chest!")
	randNum := rand.Float64()
	fmt.Println(randNum)
	var foundItem Item
	for _, item := range config.items {
		if randNum <= item.ItemDropChance {
			foundItem = item
			break
		}
	}
	if foundItem.ItemID != 0 {
		fmt.Printf("You found a %v!\n", foundItem.ItemName)
		config.player.addItem(foundItem, 1)
		return nil
	}
	fmt.Println("The chest was empty.")
	return nil
}
