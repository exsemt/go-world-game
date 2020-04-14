package model

import (
	"strings"

	"github.com/elliotchance/orderedmap"
)

// Room returns location struct
type Room struct {
	AbstractLocation
	positionDescription string
	tasks               []*Task
}

// NewRoom returns new room object
func NewRoom(name string, description string, positionDescription string) *Room {
	room := Room{positionDescription: positionDescription}
	room.AbstractLocation = AbstractLocation{name: name, description: description, availableLocations: orderedmap.NewOrderedMap()}
	// room.AbstractLocation.Location = &room

	return &room
}

// DescriptionAndAvailableLocations returns room description and available locations
func (r *Room) DescriptionAndAvailableLocations() string {
	return r.Description() + r.AvailableLocationsNames()
}

// InventoriesAndAvailableLocations returns room description and available locations
func (r *Room) InventoriesAndAvailableLocations() string {
	var info string
	var infosSlice []string
	for _, furniture := range r.furnitures {
		furnitureInfo := furniture.Info()
		if furnitureInfo != "" {
			infosSlice = append(infosSlice, furnitureInfo)
		}
	}
	if r.positionDescription != "" {
		infosSlice = append([]string{r.positionDescription}, infosSlice...)
	}

	if r.openTaskInfos() != "" {
		infosSlice = append(infosSlice, r.openTaskInfos())
	}

	info = strings.Join(infosSlice[:], ", ")
	if info == "" {
		info = "пустая комната"
	}

	return info + "." + r.AvailableLocationsNames()
}

func (r *Room) openTaskInfos() string {
	var tasksInfos []string
	for _, task := range r.tasks {
		if task.Finished == false {
			tasksInfos = append(tasksInfos, task.Name)
		}
	}

	if len(tasksInfos) <= 0 {
		return ""
	}

	return "надо " + strings.Join(tasksInfos[:], " и ")
}

// EntryLocation returns entry location
func (r *Room) EntryLocation() Location {
	return r
}
