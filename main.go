package main

func main() {
	files := loadFile()
	items := make(map[int]Item)
	for _, item := range loadItems(files[itemPath]).Items {
		items[item.ItemID] = item
	}
	quests := make(map[int]Quest)
	for _, quest := range loadQuests(files[questPath]).Quests {
		quests[quest.QuestID] = quest
	}
	p := createPlayer()
	config := config{
		player: p,
		items:  items,
		story:  loadStory(files[storyPath]),
		quests: quests,
		enemys: loadEnemys(files[enemyPath]),
	}
	startRepl(&config)
}
