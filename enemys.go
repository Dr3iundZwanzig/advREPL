package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type EnemyCollection struct {
	NormalEnemys map[string]Enemy `json:"normalEnemys"`
	HardEnemys   map[string]Enemy `json:"hardEnemys"`
}

type Enemy struct {
	Name        string  `json:"name"`
	Hp          int     `json:"hp"`
	Attack      int     `json:"attack"`
	Defense     int     `json:"defense"`
	Description string  `json:"description"`
	SpawnChance float32 `json:"spawnChance"`
	Experience  int     `json:"experience"`
}

func loadEnemys() EnemyCollection {
	file, err := os.ReadFile("enemys.json")
	if err != nil {
		log.Fatal(fmt.Errorf("Error opening file quests.json: %v", err))
	}

	var enemys EnemyCollection
	err = json.Unmarshal(file, &enemys)
	if err != nil {
		log.Fatal(fmt.Errorf("Error creating struct from quests.json: %v", err))
	}
	return enemys
}
