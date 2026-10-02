package repository

import (
	"errors"

	"github.com/wxvn/go-clicker-ws/internal/domain"
	"github.com/wxvn/go-clicker-ws/internal/mongodb"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type userDocument struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	Username  string        `bson:"username"`
	Avatar    int           `bson:"avatar"`
	Clicks    int64         `bson:"clicks"`
	CreatedAt bson.DateTime `bson:"created_at"`
}

type ClickerRepository struct {
	collection *mongo.Collection
}

func NewClickerRepository(db *mongodb.DB, database string) *ClickerRepository {
	return &ClickerRepository{
		collection: db.Database(database).Collection("users"),
	}
}

var ErrUserNotFound = errors.New("user not found")

func toDomainUser(doc userDocument) *domain.User {
	return &domain.User{
		ID:        doc.ID.Hex(),
		Username:  doc.Username,
		Avatar:    doc.Avatar,
		Clicks:    doc.Clicks,
		CreatedAt: doc.CreatedAt.Time(),
	}
}
