package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func New(cfg Config) (*DB, error) {
	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%d",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
	)

	client, err := mongo.Connect(
		options.Client().ApplyURI(uri),
	)
	if err != nil {
		return nil, err
	}

	return &DB{
		client: client,
	}, nil
}

func (db *DB) Ping(ctx context.Context) error {
	return db.client.Ping(ctx, nil)
}

func (db *DB) Database(name string) *mongo.Database {
	return db.client.Database(name)
}

func (db *DB) Init(ctx context.Context, database string) error {
	users := db.client.Database(database).Collection("users")

	_, err := users.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: map[string]any{
			"username": 1,
		},
		Options: options.Index().
			SetName("username_unique").
			SetUnique(true),
	})

	return err
}
