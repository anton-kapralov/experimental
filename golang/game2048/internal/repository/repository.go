package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/anton-kapralov/experimental/golang/game2048/internal/game"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	db     *mongo.Client
	dbName string
}

type document struct {
	Id     string    `bson:"_id"`
	Status int       `bson:"status"`
	Score  int       `bson:"score"`
	Board  [4][4]int `bson:"board"`
}

func docFromGameState(state *game.State) *document {
	d := &document{
		Status: int(state.Status),
		Score:  state.Score,
		Board:  [4][4]int{},
	}
	for i := 0; i < 4; i++ {
		copy(d.Board[i][:], state.Board[i][:])
	}
	return d
}

func (d *document) asGameState() *game.State {
	s := &game.State{
		Status: game.Status(d.Status),
		Score:  d.Score,
		Board:  [4][4]int{},
	}
	for i := 0; i < 4; i++ {
		copy(s.Board[i][:], d.Board[i][:])
	}
	return s
}

func New(ctx context.Context, dbConnUri string, dbName string) (*Repository, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(dbConnUri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %v", err)
	}
	return &Repository{
		db:     client,
		dbName: dbName,
	}, nil
}

func (r *Repository) Store(key string, state *game.State) error {
	d := docFromGameState(state)
	d.Id = key
	_, err := r.db.Database(r.dbName).Collection("games").InsertOne(context.TODO(), d)
	if err != nil {
		return fmt.Errorf("failed to insert a new game into db: %v", err)
	}
	return nil
}

func (r *Repository) Load(key string) (*game.State, error) {
	d := &document{}
	err := r.db.Database(r.dbName).Collection("games").
		FindOne(context.TODO(), bson.M{"_id": key}).
		Decode(d)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to fetch a document with the key %s: %v", key, err)
	}
	return d.asGameState(), nil
}

func (r *Repository) CompareAndSwap(
	key string, oldState *game.State, newState *game.State,
) (bool, error) {
	oldDoc := docFromGameState(oldState)
	oldDoc.Id = key
	newDoc := docFromGameState(newState)
	newDoc.Id = key
	err := r.db.Database(r.dbName).Collection("games").
		FindOneAndUpdate(context.TODO(), oldDoc, bson.D{{"$set", newDoc}}).Err()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, fmt.Errorf("failed to update a document with the key %s: %v", key, err)
	}
	return true, nil
}
