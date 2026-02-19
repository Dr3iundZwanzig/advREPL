package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Closing the Adv!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("---")
	fmt.Println("Welcome to the Adv help page!")
	fmt.Println("Usage:")
	fmt.Println("Need to use ! in front of commands")
	fmt.Println("To selcet a choice type the number infront of it")
	fmt.Println("---")
	fmt.Println("Commands:")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandPlayerInfo(p player) error {
	fmt.Println("---")
	fmt.Println("Player Information:")
	fmt.Printf("Name: %v\n", p.playerName)
	fmt.Printf("Health: %v/%v\n", p.currentHealth, p.maxHealth)
	fmt.Printf("Mana: %v/%v\n", p.currentMana, p.maxMana)
	fmt.Printf("Armour: %v\n", p.currentArmour)
	fmt.Printf("Gold: %v\n", p.gold)
	fmt.Println("---")

	return nil
}

func commandPlayerItems(p player) error {
	if len(p.items) == 0 {
		fmt.Println("You have no items yet!")
		fmt.Println("---")
		return nil
	}
	fmt.Println("---")
	fmt.Println("Player Items:")
	for _, item := range p.items {
		fmt.Printf("- %v\n", item.ItemName)
	}
	fmt.Println("---")
	return nil
}
