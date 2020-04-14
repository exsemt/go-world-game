package main

import (
	"strings"
)

// Player returns player struct
type Player struct {
	Name              string
	availableCommands map[string]interface{}
	inventories       []*Inventory
	Location
}

// NewPlayer create new player
func NewPlayer(name string) Player {
	return Player{
		Name: name,
		availableCommands: map[string]interface{}{
			"идти":        Go,
			"осмотреться": LookAround,
			"взять":       Take,
			"одеть":       Clothe,
			"применить":   Apply,
		},
	}
}

// HandleInput parsed comand and interected player
func (p *Player) HandleInput(commandWithParams string) string {
	parsedCommandWithParams := strings.Split(commandWithParams, " ")
	params := parsedCommandWithParams[1:]

	if command, exist := p.availableCommands[parsedCommandWithParams[0]]; exist {
		return command.(func(*Player, []string) string)(p, params)
	}
	return "неизвестная команда"
}

// LookAround command: description from the location + some else
func LookAround(p *Player, _ []string) string {
	return p.Location.InventoriesAndAvailableLocations()
}

// Go command: relocate player
func Go(p *Player, params []string) string {
	newLocationName := params[0]

	location := p.Location.GetAvailableLocationFor(newLocationName)
	if location == nil {
		return "нет пути в " + newLocationName
	}
	// if p.Location.CanGoTo(location) {
	// 	return "нет пути в " + newLocationName
	// }

	p.Location = location.EntryLocation()
	return p.Location.DescriptionAndAvailableLocations()
}

// Clothe command
func Clothe(p *Player, params []string) string {
	for _, inventory := range p.Location.GetInventories() {
		if inventory.Name == params[0] && inventory.Clotheable {
			if p.Location.DeleteInventory(inventory) {
				p.inventories = append(p.inventories, inventory)
			}
		}
	}

	return "вы одели: " + params[0]
}

// Take command
func Take(p *Player, params []string) string {
	for _, inventory := range p.Location.GetInventories() {
		if inventory.Name == params[0] && inventory.Takeable {
			var inventoryForAdding *Inventory
			for _, includedInventory := range p.inventories {
				if includedInventory.CanAddInventories {
					inventoryForAdding = includedInventory
				}
			}

			if inventoryForAdding == nil {
				return "некуда класть"
			} else if p.Location.DeleteInventory(inventory) {
				inventoryForAdding.Inventories = append(inventoryForAdding.Inventories, inventory)
				return "предмет добавлен в инвентарь: " + params[0]
			}
		}
	}
	return "нет такого"
}

// Apply command
func Apply(p *Player, params []string) string {
	for _, inventory := range p.inventories {
		for _, includedInventory := range inventory.Inventories {
			if includedInventory.Name == params[0] {
				if includedInventory.ApplicableTo == params[1] {
					return params[1] + " открыта"
				}
				return "не к чему применить"
			}
		}
	}

	return "нет предмета в инвентаре - " + params[0]
}
