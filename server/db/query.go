package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func queryUserAdd(
	ctx context.Context,
	db *mongo.Database,
	userID string,
	name string,
	exp int64,
	reviewCount int32,
	likeCount int32,
) error {
	coll := db.Collection(CNUser)
	_, err := coll.InsertOne(
		ctx,
		User{
			UserID:      userID,
			Name:        name,
			Exp:         exp,
			ReviewCount: reviewCount,
			LikeCount:   likeCount,
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
			"user_id": userID,
		})

	if err := cursor.Err(); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, err
		default:
			return nil, fmt.Errorf("on queryUserGet: %v", err)
		}
	}

	var user User
	if err := cursor.Decode(&user); err != nil {
		return nil, fmt.Errorf("on queryUserGet: %v", err)
	}
	return &user, nil
}

func queryUserUpdate(
	ctx context.Context,
	db *mongo.Database,
	user *User,
) error {
	coll := db.Collection(CNUser)
	cursor := coll.FindOneAndUpdate(
		ctx,
		bson.M{
			"user_id": user.UserID,
		},
		bson.M{
			"$set": bson.M{
				"name":         user.Name,
				"exp":          user.Exp,
				"review_count": user.ReviewCount,
				"like_count":   user.LikeCount,
			},
		},
	)

	if err := cursor.Err(); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return err
		default:
			return fmt.Errorf("on queryUserUpdate: %v", err)
		}
	}
	return nil
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

func queryRestaurantGet(
	ctx context.Context,
	db *mongo.Database,
	id primitive.ObjectID,
) (*Restaurant, error) {
	coll := db.Collection(CNRestaurant)
	cursor := coll.FindOne(
		ctx,
		bson.M{
			"_id": id,
		})

	if err := cursor.Err(); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, err
		default:
			return nil, fmt.Errorf("on queryRestaurantGet: %v", err)
		}
	}

	var rest Restaurant
	if err := cursor.Decode(&rest); err != nil {
		return nil, fmt.Errorf("on queryRestaurantGet: %v", err)
	}
	return &rest, nil
}

func queryRestaurantsGet(
	ctx context.Context,
	db *mongo.Database,
	minLat, maxLat float64,
	minLon, maxLon float64,
) ([]*Restaurant, error) {
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

	restaurants := make([]*Restaurant, 0)
	err = cursor.All(ctx, &restaurants)
	if err != nil {
		return nil, fmt.Errorf("on queryRestaurantsGet: %v", err)
	}

	return restaurants, nil
}

func queryReviewsGet(
	ctx context.Context,
	db *mongo.Database,
	restID primitive.ObjectID,
) ([]*Review, error) {
	coll := db.Collection(CNReview)
	cursor, err := coll.Find(
		ctx,
		bson.M{
			"restaurant_id": restID,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("on queryReviewsGet: %v", err)
	}

	reviews := make([]*Review, 0)
	err = cursor.All(ctx, &reviews)
	if err != nil {
		return nil, fmt.Errorf("on queryReviewsGet: %v", err)
	}

	return reviews, nil
}

func queryReviewAdd(
	ctx context.Context,
	db *mongo.Database,
	userID string,
	restID primitive.ObjectID,
	score int32,
	comment string,
	menus []string,
	wait bool,
	numPeople int32,
) error {
	coll := db.Collection(CNReview)
	_, err := coll.InsertOne(
		ctx,
		Review{
			UserID:    userID,
			RestID:    restID,
			Score:     score,
			Comment:   comment,
			Menus:     menus,
			Wait:      wait,
			NumPeople: numPeople,
		})

	if err != nil {
		return fmt.Errorf("on queryReviewAdd: %v", err)
	}
	return nil
}

func queryUserRankingGet(
	ctx context.Context,
	db *mongo.Database,
	top int64,
) ([]*User, error) {
	fopt := options.Find()
	fopt.SetSort(bson.M{"review_count": -1})
	fopt.SetLimit(top)

	coll := db.Collection(CNUser)
	cursor, err := coll.Find(
		ctx,
		bson.D{},
		fopt,
	)
	if err != nil {
		return nil, fmt.Errorf("on queryUserRankingGet: %v", err)
	}

	users := make([]*User, 0)
	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, fmt.Errorf("on queryUserRankingGet: %v", err)
	}

	return users, nil
}
