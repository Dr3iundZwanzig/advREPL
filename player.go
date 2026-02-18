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
}

func createPlayer(name string) player {
	char := player{
		playerName:     name,
		gold:           100,
		maxHealth:      50,
		currentHealth:  50,
		maxMana:        20,
		currentMana:    20,
		currentArmour:  0,
		currentAct:     1,
		currentChapter: 1,
		currentStep:    0,
	}
	return char
}
