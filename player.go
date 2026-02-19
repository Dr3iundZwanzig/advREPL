package main

type player struct {
	playerName     string
	gold           int
	maxHealth      int
	currentHealth  int
	maxMana        int
	currentMana    int
	currentArmour  int
	currentAct     int
	currentChapter int
	currentStep    int
	events         []Event
	items          []Item
}

func createPlayer() player {
	char := player{
		playerName:     "nameless",
		gold:           100,
		maxHealth:      50,
		currentHealth:  50,
		maxMana:        20,
		currentMana:    20,
		currentArmour:  0,
		currentAct:     1,
		currentChapter: 1,
		currentStep:    0,
		events:         []Event{},
		items:          []Item{},
	}
	return char
}
