package main

import "fmt"

func statue(config *config) error {
	if config.player.CurrentChapter == 1 && config.player.CurrentStep < 4 {
		return fmt.Errorf("Location not unlocked")
	}
	fmt.Println("You see a statue of an unknown figure, it looks like it has been there for a long time")
	if config.player.CurrentExperience.CurrentXP < config.player.CurrentExperience.NextLevelXP {
		fmt.Println("Your current experience is not enough to level up, keep fighting to gain more experience")
		return nil
	}
	player := &config.player
	fmt.Println("You feel a surge of power as you approach the statue, you have leveled up!")
	player.CurrentExperience.CurrentXP = 0
	player.CurrentExperience.NextLevelXP += 50
	player.CurrentExperience.CurrentLevel += 1
	player.MaxHealth += 10
	player.CurrentHealth = player.MaxHealth
	player.MaxMana += 5
	player.CurrentAttack += 5
	fmt.Printf("Your health has increased to %v\n", player.MaxHealth)
	fmt.Printf("Your mana has increased to %v\n", player.MaxMana)
	fmt.Printf("Your attack has increased to %v\n", player.CurrentAttack)
	fmt.Printf("Current Level: %v\n", player.CurrentExperience.CurrentLevel)
	fmt.Printf("XP needed for next Level: %v\n", player.CurrentExperience.NextLevelXP)
	return nil
}
