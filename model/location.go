package model

import (
	"strings"

	"github.com/elliotchance/orderedmap"
)

// Location interface
type Location interface {
	Name() string
	Description() string
	AvailableLocations() []Location
	AddAvailableLocation(Location, BarrierInterface)
	GetAvailableLocationFor(string) Location
	DescriptionAndAvailableLocations() string
	InventoriesAndAvailableLocations() string
	GetInventories() []*Inventory
	DeleteInventory(*Inventory) bool
	AddEntryLocation(location Location)
	EntryLocation() Location
}

// AbstractLocation struct
type AbstractLocation struct {
	Location
	name        string
	description string
	// availableLocations map[Location]BarrierInterface
	availableLocations *orderedmap.OrderedMap
	furnitures         []*Furniture
	entryLocation      Location
}

// Name returns name
func (a *AbstractLocation) Name() string {
	return a.name
}

// Description returns description
func (a *AbstractLocation) Description() string {
	return a.description
}

// AvailableLocations returns available locations
func (a *AbstractLocation) AvailableLocations() []Location {
	var locations []Location
	for _, availableLocation := range a.availableLocations.Keys() {
		locations = append(locations, availableLocation.(Location))
	}

	return locations
}

// AddAvailableLocation add available location
func (a *AbstractLocation) AddAvailableLocation(location Location, barrier BarrierInterface) {
	// if a.availableLocations == nil {
	// 	a.availableLocations = map[Location]BarrierInterface{}
	// }

	// a.availableLocations[location] = barrier
	a.availableLocations.Set(location, barrier)
}

// GetAvailableLocationFor returns matched location
func (a *AbstractLocation) GetAvailableLocationFor(locationName string) Location {
	// for _, availableLocation := range r.availableLocations {
	// 	if availableLocation.Name() == locationName {
	// 		return availableLocation
	// 	}
	// }

	for _, availableLocation := range a.availableLocations.Keys() {
		l := availableLocation.(Location)
		if l.Name() == locationName {
			return l
		}
	}

	return nil
}

// AvailableLocationsNames returns available locations names
func (a *AbstractLocation) AvailableLocationsNames() string {
	var locationNames []string

	for _, location := range a.availableLocations.Keys() {
		locationNames = append(locationNames, location.(Location).Name())
	}

	return " можно пройти - " + strings.Join(locationNames, ", ")
}

// DescriptionAndAvailableLocations returns description and available locations
func (a *AbstractLocation) DescriptionAndAvailableLocations() string {
	return a.Description() + a.AvailableLocationsNames()
}

// InventoriesAndAvailableLocations returns description and available locations
func (a *AbstractLocation) InventoriesAndAvailableLocations() string {
	return "" + a.AvailableLocationsNames()
}

// GetInventories returns inventories
func (a *AbstractLocation) GetInventories() []*Inventory {
	var inventories []*Inventory
	for _, furniture := range a.furnitures {
		for _, inventory := range furniture.inventories {
			inventories = append(inventories, inventory)
		}
	}

	return inventories
}

// DeleteInventory remove inventory
func (a *AbstractLocation) DeleteInventory(inventory *Inventory) bool {
	for _, furniture := range a.furnitures {
		for index, i := range furniture.inventories {
			if inventory == i {
				furniture.inventories = append(furniture.inventories[:index], furniture.inventories[index+1:]...)
				return true
			}
		}
	}

	return false
}

// AddEntryLocation add entry location
func (a *AbstractLocation) AddEntryLocation(location Location) {
	a.entryLocation = location
}
