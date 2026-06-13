package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

type fightCommand struct {
	name        string
	description string
	callback    func(*Player, *Enemy, string)
}

func triggerFight(config *config, room room) error {
	reader := bufio.NewScanner(os.Stdin)
	currentEnemy := Enemy{}
	player := &config.player
	switch room.name {
	case "normalfight":
		for _, enemy := range config.enemys.NormalEnemys {
			currentEnemy = enemy
			break
		}
	case "hardfight":
		for _, enemy := range config.enemys.HardEnemys {
			currentEnemy = enemy
			break
		}
	}
	fmt.Printf("You have encountered a %v with %v\n", room.name, currentEnemy.Name)
	fmt.Println("Current commands avalible:")
	fmt.Println("!attack / attacks the enemy with your current attack value")
	fmt.Println("!defend / increase your armour by 10-30 for this fight taking damage reduces this amount by the damage dealt")
	fmt.Println("!flee / attempt to flee the fight")
	fmt.Println("You cannot use items during a fight")
	fmt.Println("The enemy and you will take turns monsters have a higher chance to attack")

	for {
		fleeChance := rand.Intn(50) + rand.Intn(player.CurrentHealth/2)
		fmt.Printf("P: Your life: %v\n", player.CurrentHealth)
		fmt.Printf("P: Your armour: %v\n", player.CurrentArmour)
		fmt.Printf("E: Enemy life: %v\n", currentEnemy.Hp)
		fmt.Printf("E: Enemy armour: %v\n", currentEnemy.Defense)
		fmt.Print("Fight >>> ")
		reader.Scan()
		userInput := cleanInput(reader.Text())
		if len(userInput) == 0 {
			fmt.Println("No input entered.")
			continue
		}
		commandName := userInput[0]
		command, exists := getFightCommands()[commandName]
		if !exists && commandName != "!flee" {
			fmt.Printf("Command: %v does not exist\n", commandName)
			continue
		}
		if commandName == "!flee" {
			if currentEnemy.Hp < fleeChance {
				fmt.Println("You successfully fled the fight")
				return nil
			}
		} else {
			command.callback(player, &currentEnemy, "player")
		}

		if currentEnemy.Hp <= 0 {
			fmt.Printf("---You have killed: %v\n", currentEnemy.Name)
			fmt.Printf("---You got: %v XP\n", currentEnemy.Experience)
			player.CurrentExperience.CurrentXP += currentEnemy.Experience
			if player.CurrentExperience.CurrentXP >= player.CurrentExperience.NextLevelXP {
				fmt.Println("***Ready to level up visit the statue in town")
				player.CurrentExperience.CurrentXP = player.CurrentExperience.NextLevelXP
			}
			if player.CurrentQuests.HasQuest && player.CurrentQuests.CurrentQuest.QuestObjective == currentEnemy.Name {
				player.CurrentQuests.Progress += 1
				if questComplete(config) {
					fmt.Println("***Quest completed return to the guild to turn it in***")
				}
			}
			return nil
		}

		rnd := rand.Intn(100)
		if rnd < 30 {
			defend(player, &currentEnemy, "enemy")
		} else {
			attack(player, &currentEnemy, "enemy")
		}

		if config.player.CurrentHealth <= 0 {
			fmt.Println("you died")
			os.Exit(1)
		}
	}
}

func getFightCommands() map[string]fightCommand {
	return map[string]fightCommand{
		"!attack": {
			name:        "!attack",
			description: "Attacks an enemy",
			callback:    attack,
		},
		"!defend": {
			name:        "!defend",
			description: "Defends with current armour",
			callback:    defend,
		},
	}
}

func attack(player *Player, enemy *Enemy, attacker string) {
	switch attacker {
	case "player":
		damage := rand.Intn(player.CurrentAttack) + (player.CurrentAttack / 2)
		fmt.Println(player.CurrentAttack)
		fmt.Println(damage)
		finaldamage := damage - enemy.Defense
		enemy.Defense -= damage
		if enemy.Defense <= 0 {
			enemy.Defense = 0
		}
		if finaldamage <= 0 {
			fmt.Printf("E: %v damage absorbt by enemy armour\n", damage)
			return
		}
		enemy.Hp -= finaldamage
		fmt.Printf("P: You dealt %v damage to the enemy\n", finaldamage)
	case "enemy":
		damage := rand.Intn(enemy.Attack) + (enemy.Attack / 2)
		finaldamage := damage - player.CurrentArmour
		player.CurrentArmour -= damage
		if player.CurrentArmour <= 0 {
			player.CurrentArmour = 0
		}
		if finaldamage <= 0 {
			fmt.Printf("P: %v damage absorbt by your armour\n", damage)
			return
		}
		player.CurrentHealth -= finaldamage
		fmt.Printf("E: Enemy dealt %v damage to you\n", damage)
	}
}

func defend(player *Player, enemy *Enemy, defender string) {
	switch defender {
	case "player":
		defence := rand.Intn(20) + 10
		player.CurrentArmour += defence
		fmt.Printf("+You got %v armour\n", defence)
	case "enemy":
		defence := rand.Intn(10) + 1
		enemy.Defense += defence
		fmt.Printf("-Enemy got %v armour\n", defence)
	}
}
