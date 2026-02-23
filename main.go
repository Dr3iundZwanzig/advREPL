package main

func main() {
	config := config{
		player: createPlayer(),
		items:  loadItems(),
		story:  loadStory("Act1.json"),
	}
	startRepl(&config)
}
