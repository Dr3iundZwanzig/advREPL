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
	fmt.Printf("Your Gold: %v\n", config.player.Gold)
	for _, itemID := range keys {
		item := shopItems[itemID]
		fmt.Printf("ID: %v - %v (Cost: %v Gold, Amount: %v)\n", item.item.ItemID, item.item.ItemName, *item.item.ItemGoldCost, item.amount)
	}
	for {
		fmt.Print("ShopBuy >>> ")
		reader.Scan()
		result := reader.Text()
		if result == "!close" {
			fmt.Println("Exiting the shop.")
			break
		}
		if len(result) == 0 {
			fmt.Println("No item id entered.")
			continue
		}
		if result == "!sell" {
			err := sell(config)
			if err != nil {
				fmt.Println(err)
			}
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
			if config.player.Gold >= *shopItem.item.ItemGoldCost {
				config.player.Gold -= *shopItem.item.ItemGoldCost
				config.player.addItem(shopItem.item, 1)
				shopItems[itemID].amount -= 1
				fmt.Print(shopItems[itemID].amount)
				if shopItems[itemID].amount == 0 {

				}
				fmt.Printf("You bought %v for %v gold.\n", shopItem.item.ItemName, *shopItem.item.ItemGoldCost)
				fmt.Printf("Your remaining gold: %v\n", config.player.Gold)
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

func sell(config *config) error {
	reader := bufio.NewScanner(os.Stdin)
	p := config.player
	if len(p.Items) == 0 {
		return fmt.Errorf("You have no items to sell.")
	}
	fmt.Println("Use !close to exit the sell menu.")
	fmt.Println("Enter the ID of the item you want to sell: ")
	for itemID, items := range p.Items {
		if !items.Item.Sellable {
			continue
		}
		fmt.Printf("ID:%v -%v (Amount: %v)\n", itemID, items.Item.ItemName, items.Amount)
	}
	for {
		fmt.Print("ShopSell >>> ")
		reader.Scan()
		result := reader.Text()
		if result == "!close" {
			fmt.Println("Exiting the sell menu.")
			break
		}
		if len(result) == 0 {
			fmt.Println("No item id entered.")
			continue
		}
		itemID, err := strconv.Atoi(result)
		if err != nil {
			fmt.Println("Please enter a valid item ID.")
			continue
		}
		if item, ok := p.Items[itemID]; ok {
			if !item.Item.Sellable {
				fmt.Println("This item cannot be sold.")
				continue
			}
			p.Gold += *item.Item.ItemGoldCost / 2
			delete(p.Items, itemID)
			fmt.Printf("You sold %v for %v gold.\n", item.Item.ItemName, *item.Item.ItemGoldCost/2)
			fmt.Printf("Your current gold: %v\n", p.Gold)
			if len(p.Items) == 0 {
				fmt.Println("You have no more items to sell.")
				break
			}
			for itemID, items := range p.Items {
				if !items.Item.Sellable {
					continue
				}
				fmt.Printf("ID:%v -%v (Amount: %v)\n", itemID, items.Item.ItemName, items.Amount)
			}
		} else {
			fmt.Println("Invalid item ID.")
		}
	}
	return nil
}
