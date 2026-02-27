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
		p.addItem(config.items[1], 1)
		p.addItem(config.items[3], 1)
		fmt.Println(event.EventDescription)
		p.events[event.EventName] = event
		namePlayer(p)
	case "Old man Shop":
		fmt.Println(event.EventDescription)
		regularShopEvent(config)
	}
}

func namePlayer(p *player) {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	reader.Scan()
	p.playerName = reader.Text()
}
