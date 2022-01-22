package model

import (
	"time"
)

// Move contains the structure for a chess move with FEN notation
type Move struct {
	MoveNumber int       `json:"move_number"`
	FEN        string    `json:"fen"`
	CreatedAt  time.Time `json:"created_at"`
}

// Game contains the structure for a chess game with an array of moves
type Game struct {
	GameID      string    `json:"game_id"`
	Moves       []Move    `json:"moves"`
	CurrentMove int       `json:"current_move"`
	PlayerOneID string    `json:"player_one_id"`
	PlayerTwoID string    `json:"player_two_id"`
	CreatedAt   time.Time `json:"created_at"`
}
