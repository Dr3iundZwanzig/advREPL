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
		p.addItem(config.items.Consumables["health_potion"], 1)
		p.addItem(config.items.Trinkets["bronze_badge"], 1)
		fmt.Println(event.EventDescription)
		namePlayer(p)
	case "Old man Shop":
		p.Events[event.EventName] = event
		fmt.Println(event.EventDescription)
		err := consumableShop(config)
		if err != nil {
			return err
		}
	case "Get Quest":
		chooseQuest(config)
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
		err := triggerFight(config, room)
		if err != nil {
			return err
		}
		return nil
	case "hardfight":
		err := triggerFight(config, room)
		if err != nil {
			return err
		}
		return nil
	case "treasure":
		treasureEvent(config)
		return nil
	}
	return nil
}

func treasureEvent(config *config) {
	fmt.Println("You found a treasure chest!")

	if rand.Float64() < 0.1 {
		fmt.Println("The chest was empty.")
		return
	}

	itemPool := []Item{}
	for _, item := range config.items.Consumables {
		if item.DropChance > 0 {
			itemPool = append(itemPool, item)
		}
	}
	for _, item := range config.items.Valuables {
		if item.DropChance > 0 {
			itemPool = append(itemPool, item)
		}
	}

	totalWeight := 0.0
	for _, item := range itemPool {
		totalWeight += item.DropChance
	}

	itemCount := 1
	roll := rand.Float64()
	switch {
	case roll < 0.1:
		itemCount = 3
	case roll < 0.35:
		itemCount = 2
	}

	fmt.Printf("The chest contains %d item(s)!\n", itemCount)
	for i := 0; i < itemCount; i++ {
		rnd := rand.Float64() * totalWeight
		current := 0.0
		var chosenItem Item
		for _, item := range itemPool {
			current += item.DropChance
			if rnd <= current {
				chosenItem = item
				break
			}
		}
		config.player.addItem(chosenItem, 1)
		fmt.Printf("You found: %v!\n", chosenItem.Name)
	}
}
