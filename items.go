package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ItemEffect struct {
	HealthRestore int `json:"healthRestore,omitempty"`
	ManaRestore   int `json:"manaRestore,omitempty"`
	AttackBoost   int `json:"attackBoost,omitempty"`
	DefenseBoost  int `json:"defenseBoost,omitempty"`
}

type Item struct {
	ItemID          int         `json:"itemID"`
	ItemName        string      `json:"itemName"`
	ItemDescription string      `json:"itemDescription"`
	ItemType        string      `json:"itemType"`
	ItemEffect      *ItemEffect `json:"itemEffect"`
}

type ItemCollection struct {
	Items []Item `json:"items"`
}

func loadItems() ItemCollection {
	file, err := os.ReadFile("items.json")
	if err != nil {
		log.Fatal(fmt.Errorf("Error opening file items.json: %v", err))
	}

	var collection ItemCollection
	err = json.Unmarshal(file, &collection)
	if err != nil {
		log.Fatal(fmt.Errorf("Error creating struct from items.json: %v", err))
	}
	return collection
}
