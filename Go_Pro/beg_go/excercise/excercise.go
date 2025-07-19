package excercise

import (
	"fmt"
)

type Item struct {
	Name string
	Type string
}

type Player struct {
	Name      string
	Inventory []Item
}

func (p *Player) PickUpItem(item Item) {
	p.Inventory = append(p.Inventory, item)
	fmt.Printf("%s picked up %s!\n", p.Name, item.Name)
}

func (p *Player) DropItem(itemName string) {
	for i, item := range p.Inventory {
		if item.Name == itemName {
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			fmt.Printf("%s dropped!\n", itemName)
			return
		}
	}
	fmt.Printf("%s does not have %s in invetroty\n", p.Name, itemName)
}

func (p *Player) UseItem(itemName string) {
	for i, item := range p.Inventory {
		if item.Name == itemName {
			if item.Type == "potion" {
				fmt.Printf("%s used %s and feels rejuvenated! \n", p.Name, itemName)
				p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			} else {
				fmt.Printf("%s used %s.", p.Name, itemName)
			}
			return
		}
	}
	fmt.Printf("%s does not have %s in inventory. \n", p.Name, itemName)
}
