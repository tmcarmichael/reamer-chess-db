package dal

import (
	"../model"
)

// Define CRUD operations for DB
func CreateGame(string gameID, string playerOneID, string playerTwoID) (model.Game, error) {
	// Starting FEN: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1
	newGame := model.Game{}
	if err != nil {
		return nil, err
	}
	return newGame
}

func FetchByID(string gameID) (model.Game, error) {
	// TODO
}

func UpdateGame(string gameID) error {
	// TODO
}

func Delete(string gameID) error {
	// TODO
}
