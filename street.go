package main

import "github.com/elliotchance/orderedmap"

// Street returns location struct
type Street struct {
	AbstractLocation
}

// NewStreet returns new object
func NewStreet(name string, description string) *Street {
	street := Street{}
	street.AbstractLocation = AbstractLocation{name: name, description: description, availableLocations: orderedmap.NewOrderedMap()}
	// street.AbstractLocation.Location = &street

	return &street
}

// EntryLocation returns entry location
func (s *Street) EntryLocation() Location {
	return s
}
