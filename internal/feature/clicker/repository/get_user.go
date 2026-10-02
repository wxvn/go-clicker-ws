package repository

import (
	"context"
	"errors"

	"github.com/wxvn/go-clicker-ws/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *ClickerRepository) GetUserByID(ctx context.Context, id string) (domain.User, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return domain.User{}, ErrUserNotFound
	}

	var doc userDocument

	err = r.collection.FindOne(
		ctx,
		bson.M{"_id": objectID},
	).Decode(&doc)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.User{}, ErrUserNotFound
		}

		return domain.User{}, err
	}

	return domain.User{
		ID:        doc.ID.Hex(),
		Username:  doc.Username,
		Avatar:    doc.Avatar,
		Clicks:    doc.Clicks,
		CreatedAt: doc.CreatedAt.Time(),
	}, nil
}
