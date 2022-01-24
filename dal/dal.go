package dal

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	model "reamer-chess-db/model"
)

type MongoClient struct {
	*mongo.Client
}

// Define CRUD operations for DB
func (m MongoClient) CreateGame(gameID string, playerOneID string, playerTwoID string) error {
	// DB Create init
	collection := m.Database("ReamerDB").Collection("games")

	// Init starting position, move 0
	startingPositionFEN := model.Move{
		MoveNumber: 0,
		FEN:        "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		CreatedAt:  time.Now().Unix(),
	}

	// Create new game with starting position FEN
	newGame := model.Game{
		GameID:      gameID,
		Moves:       []model.Move{startingPositionFEN},
		CurrentMove: 1,
		PlayerOneID: playerOneID,
		PlayerTwoID: playerTwoID,
		CreatedAt:   time.Now().Unix(),
	}

	// DB Create
	insertGame, err := collection.InsertOne(context.TODO(), newGame)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("Inserted game with MongoDB id: ", insertGame.InsertedID)

	return nil
}

func (m MongoClient) FetchGameByID(gameID string) (model.Game, error) {
	// DB Retrieve init
	collection := m.Database("ReamerDB").Collection("games")
	var result model.Game
	filter := bson.D{{Key: "game_id", Value: gameID}}

	// DB Retrieve
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Retrieved game with GameID: ", gameID)

	return result, nil
}

func (m MongoClient) UpdateGame(gameID string, move model.Move) error {
	// DB Update init
	collection := m.Database("ReamerDB").Collection("games")
	filter := bson.D{{Key: "game_id", Value: gameID}}
	update := bson.D{
		{Key: "$push", Value: bson.D{
			{Key: "moves", Value: move},
		}},
	}

	// DB Update
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v games with gameID: %v\n", updateResult.ModifiedCount, gameID)

	return nil
}

func (m MongoClient) DeleteGame(gameID string) error {
	// DB Delete init
	collection := m.Database("ReamerDB").Collection("games")

	// DB Delete
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v games with gameID: %v\n", deleteResult.DeletedCount, gameID)

	return nil
}

// Data Access Object for CRUD
type DAO interface {
	CreateGame(gameID string, playerOneID string, playerTwoID string) error
	FetchGameByID(gameID string) (model.Game, error)
	UpdateGame(gameID string) error
	Delete(gameID string) error
}
