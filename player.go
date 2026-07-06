package main

import (
	"fmt"
)

type Player struct {
	PlayerName        string              `json:"playerName"`
	Gold              int                 `json:"gold"`
	MaxHealth         int                 `json:"maxHealth"`
	CurrentHealth     int                 `json:"currentHealth"`
	MaxMana           int                 `json:"maxMana"`
	CurrentMana       int                 `json:"currentMana"`
	CurrentArmour     int                 `json:"currentArmour"`
	CurrentAttack     int                 `json:"currentAttack"`
	CurrentAct        int                 `json:"currentAct"`
	CurrentChapter    int                 `json:"currentChapter"`
	CurrentStep       int                 `json:"currentStep"`
	InDungeon         bool                `json:"inDungeon"`
	Events            map[string]Event    `json:"events"`
	Items             map[int]*PlayerItem `json:"items"`
	CurrentQuests     PlayerQuest         `json:"currentQuests"`
	CurrentExperience PlayerExperience    `json:"experience"`
	AutoSaveFileName  string              `json:"autoSaveFileName"`
	AutoSave          bool                `json:"autoSave"`
}

type PlayerQuest struct {
	HasQuest     bool  `json:"hasQuest"`
	CurrentQuest Quest `json:"currentQuest"`
	Progress     int   `json:"progress"`
}

type PlayerItem struct {
	Amount int  `json:"amount"`
	Item   Item `json:"item"`
}

type PlayerExperience struct {
	CurrentLevel     int    `json:"currentLevel"`
	CurrentXP        int    `json:"currentXP"`
	NextLevelXP      int    `json:"nextLevelXP"`
	CurrentGuildRank string `json:"currentGuildRank"`
	CurrentGuildXP   int    `json:"currentGuildXP"`
	NextGuildLevelXP int    `json:"nextGuildLevelXP"`
}

func createPlayer() Player {
	char := Player{
		PlayerName:     "nameless",
		Gold:           100,
		MaxHealth:      50,
		CurrentHealth:  50,
		MaxMana:        20,
		CurrentMana:    5,
		CurrentArmour:  0,
		CurrentAttack:  15,
		CurrentAct:     1,
		CurrentChapter: 1,
		CurrentStep:    0,
		InDungeon:      false,
		Events:         map[string]Event{},
		Items:          map[int]*PlayerItem{},
		CurrentQuests: PlayerQuest{
			HasQuest:     false,
			CurrentQuest: Quest{},
			Progress:     0,
		},
		CurrentExperience: PlayerExperience{
			CurrentLevel:     1,
			CurrentXP:        0,
			NextLevelXP:      100,
			CurrentGuildRank: "Unregistered",
			CurrentGuildXP:   0,
			NextGuildLevelXP: 100,
		},
	}
	return char
}

func (player *Player) addItem(item Item, amount int) {
	if existingItem, ok := player.Items[item.ItemID]; ok {
		existingItem.Amount += amount
	} else {
		player.Items[item.ItemID] = &PlayerItem{
			Amount: amount,
			Item:   item,
		}
	}
	fmt.Printf("+++ Got item: %v\n", item.Name)
}

func (player *Player) useItem(itemID int) {
	if existingItem, ok := player.Items[itemID]; !ok {
		fmt.Printf("You don't have an item with the ID %v!\n", itemID)
		return
	} else {
		if existingItem.Item.Type != "Consumable" {
			fmt.Printf("You cannot use this item!\n")
			return
		}
		existingItem.Amount -= 1
		if existingItem.Amount <= 0 {
			delete(player.Items, itemID)
		}
		effect := existingItem.Item.Effect
		if effect != nil {
			player.CurrentHealth += effect.HealthRestore
			if player.CurrentHealth > player.MaxHealth {
				player.CurrentHealth = player.MaxHealth
			}
			player.CurrentMana += effect.ManaRestore
			if player.CurrentMana > player.MaxMana {
				player.CurrentMana = player.MaxMana
			}
			player.CurrentArmour += effect.DefenseBoost
			player.CurrentAttack += effect.AttackBoost
		}
		fmt.Printf("You used %v!\n", existingItem.Item.Name)
	}

}
