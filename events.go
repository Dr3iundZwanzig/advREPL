package main

import (
	"bufio"
	"fmt"
	"os"
)

func triggerEvent(event Event, config *config) {
	items := make(map[int]Item)
	for _, item := range config.items.Items {
		items[item.ItemID] = item
	}
	if _, ok := config.player.events[event.EventName]; ok {
		return
	}
	p := &config.player
	switch event.EventName {
	case "Guild Registration":
		p.items = append(p.items, items[1])
		p.items = append(p.items, items[2])
		p.items = append(p.items, items[3])
		fmt.Println(event.EventDescription)
		p.events[event.EventName] = event
		namePlayer(p)
	}
}

func namePlayer(p *player) {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	reader.Scan()
	p.playerName = reader.Text()
}
