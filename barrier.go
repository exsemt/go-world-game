package main

// BarrierInterface interface
type BarrierInterface interface {
}

// Barrier struct
type Barrier struct {
	Name string
	// To        Location
	Closed    bool
	CanOpenBy *Inventory
}

// NewBarrier returns Barrier object
func NewBarrier() *Barrier {
	return &Barrier{}
}
