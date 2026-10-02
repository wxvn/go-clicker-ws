package repository

import (
	"context"
	"errors"

	"github.com/wxvn/go-clicker-ws/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (r *ClickerRepository) GetLeaderboard(ctx context.Context, limit int, userID string) (domain.LeaderboardResult, error) {
	cursor, err := r.collection.Find(ctx, bson.M{}, options.Find().SetSort(bson.D{{Key: "clicks", Value: -1}, {Key: "_id", Value: 1}}).SetLimit(int64(limit)))
	if err != nil {
		return domain.LeaderboardResult{}, err
	}
	defer cursor.Close(ctx)

	var docs []userDocument
	if err := cursor.All(ctx, &docs); err != nil {
		return domain.LeaderboardResult{}, err
	}

	users := make([]domain.Leaderboard, 0, len(docs))
	for i, doc := range docs {
		position := int64(i + 1)
		users = append(users, domain.Leaderboard{
			ID:       doc.ID.Hex(),
			Username: doc.Username,
			Avatar:   doc.Avatar,
			Clicks:   doc.Clicks,
			Position: &position,
		})
	}

	result := domain.LeaderboardResult{Users: users}
	if userID == "" {
		return result, nil
	}

	objectID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return domain.LeaderboardResult{}, ErrUserNotFound
	}

	var user userDocument
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.LeaderboardResult{}, ErrUserNotFound
		}
		return domain.LeaderboardResult{}, err
	}

	filter := bson.M{"$or": bson.A{
		bson.M{"clicks": bson.M{"$gt": user.Clicks}},
		bson.M{"clicks": user.Clicks, "_id": bson.M{"$lt": user.ID}},
	}}

	position, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return domain.LeaderboardResult{}, err
	}
	position++

	result.CurrentUser = &domain.Leaderboard{
		ID:       user.ID.Hex(),
		Username: user.Username,
		Avatar:   user.Avatar,
		Clicks:   user.Clicks,
		Position: &position,
	}

	return result, nil
}
