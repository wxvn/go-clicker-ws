package repository

import (
	"context"
	"errors"

	"github.com/wxvn/go-clicker-ws/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *AuthRepository) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	var doc authDocument

	err := r.collection.FindOne(
		ctx,
		bson.M{"username": username},
	).Decode(&doc)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return toDomainUser(doc), nil
}
