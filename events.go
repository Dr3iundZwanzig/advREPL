package main

import (
	"bufio"
	"fmt"
	"os"
)

func triggerEvent(event Event, p *player, items ItemCollection) {
	switch event.EventName {
	case "Guild Registration":
		for _, item := range items.Items {
			if item.ItemName == "Bronze Badge" {
				p.items = append(p.items, item)
				break
			}
		}
		fmt.Println(event.EventDescription)
		p.events = append(p.events, event)
		namePlayer(p)
	}
}

func namePlayer(p *player) {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	reader.Scan()
	p.playerName = reader.Text()
}
