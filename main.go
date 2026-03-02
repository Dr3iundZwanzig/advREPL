package main

import (
	"fmt"
)

func main() {
	items := make(map[int]Item)
	for _, item := range loadItems().Items {
		items[item.ItemID] = item
	}
	p := createPlayer()
	config := config{
		player: p,
		items:  items,
		story:  loadStory(fmt.Sprintf("Chapter%v.json", p.currentChapter)),
	}
	startRepl(&config)
}
