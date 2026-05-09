package main

import (
	"bufio"
	"fmt"
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
		fmt.Println(room.description)
		return nil
	}
	return nil
}
