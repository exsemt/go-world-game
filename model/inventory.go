package model

// Inventory struct
type Inventory struct {
	Name              string
	Clotheable        bool
	CanAddInventories bool
	Takeable          bool
	ApplicableTo      string
	Inventories       []*Inventory
}

// NewInventory create new inventory
func NewInventory(name string, clotheable bool, canAddInventories bool, takeable bool, applicableTo string) *Inventory {
	return &Inventory{Name: name, Clotheable: clotheable, CanAddInventories: canAddInventories, Takeable: takeable, ApplicableTo: applicableTo}
}
