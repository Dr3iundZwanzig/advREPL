package main

func main() {
	items := make(map[int]Item)
	for _, item := range loadItems().Items {
		items[item.ItemID] = item
	}
	config := config{
		player: createPlayer(),
		items:  items,
		story:  loadStory("Act1.json"),
	}
	startRepl(&config)
}
