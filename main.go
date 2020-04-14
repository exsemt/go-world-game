package main

func main() {}

var gameWord Game

// Game returns game world
type Game struct {
	Players []Player
}

func initGame() *Game {
	corridor := NewRoom("коридор", "ничего интересного.", "")

	kitchen := NewRoom("кухня", "кухня, ничего интересного.", "ты находишься на кухне")
	kitchen.furnitures = []*Furniture{NewFurniture("стол", "на столе", []*Inventory{NewInventory("чай", false, false, false, "")})}
	kitchen.tasks = []*Task{NewTask("собрать рюкзак"), NewTask("идти в универ")}

	myRoom := NewRoom("комната", "ты в своей комнате.", "")
	myRoom.furnitures = []*Furniture{
		NewFurniture("стол", "на столе:", []*Inventory{NewInventory("ключи", false, false, true, "дверь"), NewInventory("конспекты", false, false, true, "")}),
		NewFurniture("стул", "на стуле -", []*Inventory{NewInventory("рюкзак", true, true, false, "")}),
	}

	house := NewHouse("домой", "")
	house.AddEntryLocation(corridor)

	street := NewStreet("улица", "на улице весна.")

	myRoom.AddAvailableLocation(corridor, nil)
	corridor.AddAvailableLocation(kitchen, nil)
	corridor.AddAvailableLocation(myRoom, nil)
	kitchen.AddAvailableLocation(corridor, nil)

	street.AddAvailableLocation(house, nil)
	corridor.AddAvailableLocation(street, nil)

	player := NewPlayer("SuperPlayer")
	player.Location = kitchen
	gameWord = Game{Players: []Player{player}}

	return &gameWord
}

func handleCommand(commandWithParams string) string {
	return gameWord.Players[0].HandleInput(commandWithParams)
}
