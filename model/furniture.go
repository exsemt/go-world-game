package model

import "strings"

// Furniture struct
type Furniture struct {
	Name        string
	Description string
	inventories []*Inventory
}

// NewFurniture create new furniture
func NewFurniture(name string, description string, inventories []*Inventory) *Furniture {
	return &Furniture{name, description, inventories}
}

// Info returns information regarding furniture and his inventories
func (f *Furniture) Info() string {
	var info string
	var inventoriesNames []string
	for _, i := range f.inventories {
		inventoriesNames = append(inventoriesNames, i.Name)
	}

	if len(inventoriesNames) > 0 {
		info = f.Description + " " + strings.Join(inventoriesNames[:], ", ")
	}

	return info
}
