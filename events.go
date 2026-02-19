package main

import (
	"bufio"
	"fmt"
	"os"
)

func triggerEvent(event Event, p *player, items map[int]Item) {
	switch event.EventName {
	case "Guild Registration":
		p.items = append(p.items, items[1])
		p.items = append(p.items, items[2])
		p.items = append(p.items, items[3])
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
