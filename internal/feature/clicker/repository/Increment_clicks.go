package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (r *ClickerRepository) IncrementClicks(ctx context.Context, id string) (int64, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return 0, ErrUserNotFound
	}

	var doc userDocument

	err = r.collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": objectID},
		bson.M{
			"$inc": bson.M{
				"clicks": 1,
			},
		},
		options.FindOneAndUpdate().
			SetReturnDocument(options.After),
	).Decode(&doc)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return 0, ErrUserNotFound
		}

		return 0, err
	}

	return doc.Clicks, nil
}
