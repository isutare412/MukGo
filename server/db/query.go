package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func queryUserAdd(
	ctx context.Context,
	db *mongo.Database,
	userID string,
	name string,
	exp int64,
) error {
	coll := db.Collection(CNUser)
	_, err := coll.InsertOne(
		ctx,
		User{
			UserID: userID,
			Name:   name,
			Exp:    exp,
		})

	if err != nil {
		return fmt.Errorf("on queryUserAdd: %v", err)
	}
	return nil
}

func queryUserGet(
	ctx context.Context,
	db *mongo.Database,
	userID string,
) (*User, error) {
	coll := db.Collection(CNUser)
	cursor := coll.FindOne(
		ctx,
		bson.M{
			"userid": userID,
		})

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("on queryUserGet: %v", err)
	}

	var user User
	if err := cursor.Decode(&user); err != nil {
		return nil, fmt.Errorf("on queryUserGet: %v", err)
	}
	return &user, nil
}

func queryRestaurantAdd(
	ctx context.Context,
	db *mongo.Database,
	name string,
	latitude, longitude float64,
) error {
	coll := db.Collection(CNRestaurant)
	_, err := coll.InsertOne(
		ctx,
		Restaurant{
			Name:      name,
			Latitude:  latitude,
			Longitude: longitude,
		})

	if err != nil {
		return fmt.Errorf("on queryRestaurantAdd: %v", err)
	}
	return nil
}

func queryRestaurantsGet(
	ctx context.Context,
	db *mongo.Database,
	minLat, maxLat float64,
	minLon, maxLon float64,
) ([]Restaurant, error) {
	coll := db.Collection(CNRestaurant)
	cursor, err := coll.Find(
		ctx,
		bson.M{
			"latitude": bson.M{
				"$gt": minLat,
				"$lt": maxLat,
			},
			"longitude": bson.M{
				"$gt": minLon,
				"$lt": maxLon,
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("on queryRestaurantsGet: %v", err)
	}

	restaurants := make([]Restaurant, 0)
	err = cursor.All(ctx, &restaurants)
	if err != nil {
		return nil, fmt.Errorf("on queryRestaurantsGet: %v", err)
	}

	return restaurants, nil
}

func queryReviewAdd(
	ctx context.Context,
	db *mongo.Database,
	userID string,
	score int,
	comment string,
) error {
	coll := db.Collection(CNReview)
	_, err := coll.InsertOne(
		ctx,
		Review{
			UserID:  userID,
			Score:   score,
			Comment: comment,
		})

	if err != nil {
		return fmt.Errorf("on queryReviewAdd: %v", err)
	}
	return nil
}
