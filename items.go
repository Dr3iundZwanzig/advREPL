package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type ItemEffect struct {
	HealthRestore int `json:"healthRestore,omitempty"`
	ManaRestore   int `json:"manaRestore,omitempty"`
	AttackBoost   int `json:"attackBoost,omitempty"`
	DefenseBoost  int `json:"defenseBoost,omitempty"`
}

type Item struct {
	ItemID      int         `json:"itemID"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Type        string      `json:"type"`
	GoldCost    *int        `json:"goldCost"`
	DropChance  float64     `json:"dropChance"`
	Effect      *ItemEffect `json:"effect"`
	Sellable    bool        `json:"sellable"`
}

type ItemCollection struct {
	Consumables map[string]Item `json:"consumables"`
	Trinkets    map[string]Item `json:"trinkets"`
	Valuables   map[string]Item `json:"valuables"`
}

func loadItems(file []byte) ItemCollection {
	var collection ItemCollection
	err := json.Unmarshal(file, &collection)
	if err != nil {
		log.Fatal(fmt.Errorf("Error creating struct from items.json: %v", err))
	}
	return collection
}

func (ic ItemCollection) GetItemByID(id int) (Item, bool) {
	for _, it := range ic.Consumables {
		if it.ItemID == id {
			return it, true
		}
	}
	for _, it := range ic.Trinkets {
		if it.ItemID == id {
			return it, true
		}
	}
	for _, it := range ic.Valuables {
		if it.ItemID == id {
			return it, true
		}
	}
	return Item{}, false
}
