package repository

import (
	"context"
	"errors"

	"github.com/wxvn/go-clicker-ws/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *AuthRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	doc := authDocument{
		Username:     user.Username,
		PasswordHash: user.PasswordHash,
		Avatar:       user.Avatar,
		Clicks:       user.Clicks,
		CreatedAt:    bson.NewDateTimeFromTime(user.CreatedAt),
	}

	result, err := r.collection.InsertOne(ctx, doc)
	if err != nil {
		var writeErr mongo.WriteException

		if errors.As(err, &writeErr) {
			for _, item := range writeErr.WriteErrors {
				if item.Code == 11000 {
					return nil, ErrUsernameConflict
				}
			}
		}

		return nil, err
	}

	objectID, ok := result.InsertedID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("invalid inserted user id")
	}

	user.ID = objectID.Hex()

	return user, nil
}
