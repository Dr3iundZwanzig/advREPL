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
	currentAttack  int
	currentAct     int
	currentChapter int
	currentStep    int
	events         map[string]Event
	items          map[int]*playerItem
}

type playerItem struct {
	amount int
	item   Item
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
		currentAttack:  5,
		currentAct:     1,
		currentChapter: 1,
		currentStep:    0,
		events:         map[string]Event{},
		items:          map[int]*playerItem{},
	}
	return char
}

func (player *player) addItem(item Item, amount int) {
	if existingItem, ok := player.items[item.ItemID]; ok {
		existingItem.amount += amount
		fmt.Println("item up")
		fmt.Println(amount)
	} else {
		fmt.Println("item NEW")
		player.items[item.ItemID] = &playerItem{
			amount: amount,
			item:   item,
		}
	}
}

func (player *player) useItem(itemID int) {
	if existingItem, ok := player.items[itemID]; !ok {
		fmt.Printf("You don't have an item with the ID %v!\n", itemID)
		return
	} else {
		if existingItem.item.ItemType != "Consumable" {
			fmt.Printf("You cannot use this item!\n")
			return
		}
		existingItem.amount -= 1
		if existingItem.amount <= 0 {
			delete(player.items, itemID)
		}
		effect := existingItem.item.ItemEffect
		player.currentHealth += effect.HealthRestore
		if player.currentHealth > player.maxHealth {
			player.currentHealth = player.maxHealth
		}
		player.currentMana += effect.ManaRestore
		if player.currentMana > player.maxMana {
			player.currentMana = player.maxMana
		}
		player.currentArmour += effect.DefenseBoost
		player.currentAttack += effect.AttackBoost
		fmt.Printf("You used %v!\n", existingItem.item.ItemName)
	}

}
