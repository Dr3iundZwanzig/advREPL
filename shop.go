package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type ShopItems struct {
	amount int
	item   Item
}

func fillShopItems(config *config, itemIDs []int) (map[int]*ShopItems, error) {
	shopItems := make(map[int]*ShopItems)
	for _, itemID := range itemIDs {
		if item, ok := config.items[itemID]; ok {
			shopItems[itemID] = &ShopItems{
				amount: 5,
				item:   item,
			}
		} else {
			return nil, fmt.Errorf("Item with ID %v not found in config", itemID)
		}
	}
	return shopItems, nil
}

func regularShop(config *config) error {
	reader := bufio.NewScanner(os.Stdin)
	shopItems, ok := fillShopItems(config, []int{1, 2})
	if ok != nil {
		return ok
	}
	keys := make([]int, 0, len(shopItems))
	for k := range shopItems {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println("Use !close to exit the shop.")
	fmt.Println("Enter the ID of the item you want to buy: ")
	fmt.Printf("Your Gold: %v\n", config.player.gold)
	for _, itemID := range keys {
		item := shopItems[itemID]
		fmt.Printf("ID: %v - %v (Cost: %v Gold, Amount: %v)\n", item.item.ItemID, item.item.ItemName, *item.item.ItemGoldCost, item.amount)
	}
	for {
		fmt.Print("Adv >>> ")
		reader.Scan()
		result := reader.Text()
		if result == "!close" {
			fmt.Println("Exiting the shop.")
			break
		}

		if len(result) == 0 {
			fmt.Println("No tiem id entered.")
			continue
		}
		itemID, err := strconv.Atoi(result)
		if err != nil {
			fmt.Println("Please enter a valid item ID.")
			continue
		}
		if shopItem, ok := shopItems[itemID]; ok {
			if shopItems[itemID].amount == 0 {
				fmt.Println("This item is out of stock.")
				continue
			}
			if config.player.gold >= *shopItem.item.ItemGoldCost {
				config.player.gold -= *shopItem.item.ItemGoldCost
				config.player.addItem(shopItem.item, 1)
				shopItems[itemID].amount -= 1
				fmt.Print(shopItems[itemID].amount)
				if shopItems[itemID].amount == 0 {

				}
				fmt.Printf("You bought %v for %v gold.\n", shopItem.item.ItemName, *shopItem.item.ItemGoldCost)
				fmt.Printf("Your remaining gold: %v\n", config.player.gold)
				for _, itemID := range keys {
					item := shopItems[itemID]
					fmt.Printf("ID: %v - %v (Cost: %v Gold, Amount: %v)\n", item.item.ItemID, item.item.ItemName, *item.item.ItemGoldCost, item.amount)
				}
			} else {
				fmt.Println("You don't have enough gold.")
			}
		} else {
			fmt.Println("Invalid item ID.")
		}
	}
	return nil
}
