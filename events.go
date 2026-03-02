package main

import (
	"bufio"
	"fmt"
	"os"
)

func triggerEvent(event Event, config *config) {

	if _, ok := config.player.events[event.EventName]; ok {
		return
	}
	p := &config.player
	switch event.EventName {
	case "Guild Registration":
		p.events[event.EventName] = event
		p.experience.currentGuildRank = "Bronze"
		p.addItem(config.items[1], 1)
		p.addItem(config.items[3], 1)
		fmt.Println(event.EventDescription)
		namePlayer(p)
	case "Old man Shop":
		p.events[event.EventName] = event
		fmt.Println(event.EventDescription)
		regularShopEvent(config)
	case "Get Quest":
		fmt.Println(event.EventDescription)
	case "Open World":
		p.events[event.EventName] = event
		fmt.Println("-------------------------------")
		fmt.Println("You are now free to explore the world! Type !help for a list of commands and !playerinfo to see your stats.\nTo continue with the story reach Level 5 and the silver guild rank by doing quests.")
		fmt.Println("-------------------------------")
	}
}

func namePlayer(p *player) {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	reader.Scan()
	p.playerName = reader.Text()
}
