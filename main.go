package main

func main() {
	files := loadFile()
	p := createPlayer()
	config := config{
		player: p,
		items:  loadItems(files[itemPath]),
		story:  loadStory(files[storyPath]),
		quests: loadQuests(files[questPath]),
		enemys: loadEnemys(files[enemyPath]),
	}
	startRepl(&config)
}
