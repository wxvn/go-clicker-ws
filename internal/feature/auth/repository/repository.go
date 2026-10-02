package repository

import (
	"errors"

	"github.com/wxvn/go-clicker-ws/internal/domain"
	"github.com/wxvn/go-clicker-ws/internal/mongodb"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type authDocument struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	Username     string        `bson:"username"`
	PasswordHash string        `bson:"password_hash"`
	Avatar       int           `bson:"avatar"`
	Clicks       int64         `bson:"clicks"`
	CreatedAt    bson.DateTime `bson:"created_at"`
}

type AuthRepository struct {
	collection *mongo.Collection
}

func NewAuthRepository(db *mongodb.DB, database string) *AuthRepository {
	return &AuthRepository{
		collection: db.Database(database).Collection("users"),
	}
}

func toDomainUser(doc authDocument) *domain.User {
	return &domain.User{
		ID:           doc.ID.Hex(),
		Username:     doc.Username,
		PasswordHash: doc.PasswordHash,
		Avatar:       doc.Avatar,
		Clicks:       doc.Clicks,
		CreatedAt:    doc.CreatedAt.Time(),
	}
}

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrUsernameConflict = errors.New("username already exists")
)
