package main

import (
	"fmt"
	"log"
	"os"
)

const enemyPath = "internal/json/enemys.json"
const itemPath = "internal/json/items.json"
const questPath = "internal/json/quests.json"
const storyPath = "internal/json/Chapter1.json"

func loadFile() map[string][]byte {
	files := map[string][]byte{}

	enemyFile, err := os.ReadFile(enemyPath)
	if err != nil {
		log.Fatal(fmt.Errorf("Error opening file %v: %v", enemyPath, err))
	}
	files[enemyPath] = enemyFile

	itemFile, err := os.ReadFile(itemPath)
	if err != nil {
		log.Fatal(fmt.Errorf("Error opening file %v: %v", itemPath, err))
	}
	files[itemPath] = itemFile

	questFile, err := os.ReadFile(questPath)
	if err != nil {
		log.Fatal(fmt.Errorf("Error opening file %v: %v", questPath, err))
	}
	files[questPath] = questFile

	storyFile, err := os.ReadFile(storyPath)
	if err != nil {
		log.Fatal(fmt.Errorf("Error opening file %v: %v", storyPath, err))
	}
	files[storyPath] = storyFile

	return files
}
