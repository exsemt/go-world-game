package main

import (
	"github.com/exsemt/go-world-game/model"
)

func main() {}

var gameWord Game

// Game returns game world
type Game struct {
	Players []model.Player
}

func initGame() *Game {
	corridor := model.NewRoom("коридор", "ничего интересного.", "")

	kitchen := model.NewRoom("кухня", "кухня, ничего интересного.", "ты находишься на кухне")
	// kitchen.furnitures = []*model.Furniture{model.NewFurniture("стол", "на столе", []*model.Inventory{model.NewInventory("чай", false, false, false, "")})}
	// kitchen.tasks = []*model.Task{model.NewTask("собрать рюкзак"), model.NewTask("идти в универ")}

	myRoom := model.NewRoom("комната", "ты в своей комнате.", "")
	// myRoom.furnitures = []*model.Furniture{
	// 	model.NewFurniture("стол", "на столе:", []*model.Inventory{model.NewInventory("ключи", false, false, true, "дверь"), model.NewInventory("конспекты", false, false, true, "")}),
	// 	model.NewFurniture("стул", "на стуле -", []*model.Inventory{model.NewInventory("рюкзак", true, true, false, "")}),
	// }

	house := model.NewHouse("домой", "")
	house.AddEntryLocation(corridor)

	street := model.NewStreet("улица", "на улице весна.")

	myRoom.AddAvailableLocation(corridor, nil)
	corridor.AddAvailableLocation(kitchen, nil)
	corridor.AddAvailableLocation(myRoom, nil)
	kitchen.AddAvailableLocation(corridor, nil)

	street.AddAvailableLocation(house, nil)
	corridor.AddAvailableLocation(street, nil)

	player := model.NewPlayer("SuperPlayer")
	player.Location = kitchen
	gameWord = Game{Players: []model.Player{player}}

	return &gameWord
}

func handleCommand(commandWithParams string) string {
	return gameWord.Players[0].HandleInput(commandWithParams)
}
