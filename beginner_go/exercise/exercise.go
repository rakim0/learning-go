package exercise

import (
	"fmt"
)

type Item struct {
	Name     string
	ItemType string
}

type Player struct {
	Name      string
	Inventory []Item
}

func (p *Player) PickUpAnItem(item Item) {
	p.Inventory = append(p.Inventory, item)
	fmt.Print("%s picked up %s!\n", p.Name, item.Name)
}

func (p *Player) DropAnItem(itemName string) {
	for index, item := range p.Inventory {
		if item.Name == itemName {
			p.Inventory = append(p.Inventory[:index], p.Inventory[index+1:]...)
			fmt.Printf("%s dropped %s!\n", p.Name, itemName)
			return
		}
	}
	fmt.Printf("%s doesn't have the item: %s", p.Name, itemName)
}

func (p *Player) UseItem(itemName string) string {
	for index, item := range p.Inventory {
		if item.Name == itemName {
			p.Inventory = append(p.Inventory[:index], p.Inventory[index+1:]...)
		}
	}
	return "Using " + itemName
}
