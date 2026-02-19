package main

import (
	"fmt"
)

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
		currentMana:    5,
		currentArmour:  0,
		currentAct:     1,
		currentChapter: 1,
		currentStep:    0,
		events:         []Event{},
		items:          []Item{},
	}
	return char
}

func (player *player) addItem(item Item) {
	player.items = append(player.items, item)
}

func (player *player) hasItem(itemID int) bool {
	for _, item := range player.items {
		if item.ItemID == itemID {
			return true
		}
	}
	return false
}

func (player *player) useItem(itemID int) {
	for i, item := range player.items {
		if item.ItemID == itemID {
			if item.ItemType == "Consumable" {
				if item.ItemEffect != nil {
					player.currentHealth += item.ItemEffect.HealthRestore
					if player.currentHealth > player.maxHealth {
						player.currentHealth = player.maxHealth
					}
					player.currentMana += item.ItemEffect.ManaRestore
					if player.currentMana > player.maxMana {
						player.currentMana = player.maxMana
					}
					player.currentArmour += item.ItemEffect.AttackBoost
					player.currentArmour += item.ItemEffect.DefenseBoost
					player.items = append(player.items[:i], player.items[i+1:]...)
					fmt.Printf("You used %v!\n", item.ItemName)
					return
				}
			} else {
				fmt.Printf("%v is not a consumable item!\n", item.ItemName)
			}
		}
	}
}
