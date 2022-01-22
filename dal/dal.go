package dal

import (
	"time"

	model "reamer-chess-db/model"
)

// Define CRUD operations for DB
func CreateGame(gameID string, playerOneID string, playerTwoID string) (model.Game, error) {
	// Init starting position, move 0
	startingPositionFEN := model.Move{MoveNumber: 0, FEN: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", CreatedAt: time.Now().Unix()}

	// Create new game with starting position FEN
	newGame := model.Game{GameID: gameID, Moves: []model.Move{startingPositionFEN}, CurrentMove: 1, PlayerOneID: playerOneID, PlayerTwoID: playerTwoID, CreatedAt: time.Now().Unix()}

	// DB Create
}

func FetchByID(gameID string) (model.Game, error) {

	// DB Retrieve
}

func UpdateGame(gameID string) error {

	// DB Update
}

func Delete(gameID string) error {

	// DB Delete
}

type DAO interface {
	CreateGame(gameID string, playerOneID string, playerTwoID string) (model.Game, error)
	FetchByID(gameID string) (model.Game, error)
	UpdateGame(gameID string) error
	Delete(gameID string) error
}
