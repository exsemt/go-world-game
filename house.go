package main

import "github.com/elliotchance/orderedmap"

// House returns location struct
type House struct {
	AbstractLocation
}

// NewHouse returns new object
func NewHouse(name string, description string) *House {
	house := House{}
	house.AbstractLocation = AbstractLocation{name: name, description: description, availableLocations: orderedmap.NewOrderedMap()}
	// house.AbstractLocation.Location = &house

	return &house
}

// EntryLocation returns entry location
func (h *House) EntryLocation() Location {
	return h.entryLocation
}
