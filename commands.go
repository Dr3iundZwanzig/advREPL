package main

import (
	"fmt"
	"os"
	"strconv"
)

func commandExit(config *config, _ ...string) error {
	fmt.Println("Closing the Adv!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config, _ ...string) error {
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

func commandPlayerInfo(config *config, _ ...string) error {
	p := config.player
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

func commandPlayerItems(config *config, _ ...string) error {
	p := config.player
	if len(p.items) == 0 {
		fmt.Println("You have no items yet!")
		fmt.Println("---")
		return nil
	}
	fmt.Println("---")
	fmt.Println("Player Items:")
	for itemName, items := range p.items {
		fmt.Printf("ID:%v -%v (Amount: %v)\n", items.item.ItemID, itemName, items.amount)
	}
	fmt.Println("---")
	return nil
}

func commandUseItem(config *config, args ...string) error {
	if len(args) < 1 {
		fmt.Println("Usage: !use [itemID]")
		return nil
	}
	itemID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid item ID")
		return nil
	}
	config.player.useItem(itemID)
	return nil
}

func commandSelectChoice(config *config, args ...string) error {
	if len(args) < 1 {
		fmt.Println("Usage: !choice [choiceNumber]")
		return nil
	}
	choiceNumber, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid choice number")
		return nil
	}
	currentStep := config.story.ChapterSteps[config.player.currentStep]
	if choiceNumber < 1 || choiceNumber > len(currentStep.TriggerChoice) {
		fmt.Println("Invalid choice number")
		return nil
	}
	config.player.currentStep = currentStep.TriggerChoice[choiceNumber-1].ChoiceNextStep
	continueStory(config)
	return nil
}
