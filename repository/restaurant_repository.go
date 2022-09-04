package repository

import (
	"context"
	"swarch/poptum/restaurants/config"
	"swarch/poptum/restaurants/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RestaurantRepository interface {
	GetRestaurants() ([]model.Restaurant, error)
	GetRestaurant(primitive.ObjectID) (*model.Restaurant, error)
	CreateRestaurant(model.Restaurant) (*mongo.InsertOneResult, error)
	EditRestaurant(primitive.ObjectID, *model.Restaurant) (*model.Restaurant, error)
	DeleteRestaurant(primitive.ObjectID) (*model.Restaurant, error)
}

type restaurantRepository struct {
	mongoClient          *mongo.Client
	restaurantCollection *mongo.Collection
}

// NewRestaurantRepository -> returns new restaurant Repository
func NewRestaurantRepository(mongoClient *mongo.Client) RestaurantRepository {
	return restaurantRepository{
		mongoClient:          mongoClient,
		restaurantCollection: config.GetCollection(mongoClient, "Restaurants"),
	}
}

// GetRestaurants implements RestaurantRepository
func (r restaurantRepository) GetRestaurants() ([]model.Restaurant, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var restaurants []model.Restaurant

	results, err := r.restaurantCollection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var restaurant model.Restaurant
		if err = results.Decode(&restaurant); err != nil {
			return nil, err
		}

		restaurants = append(restaurants, restaurant)
	}
	return restaurants, nil
}

// GetRestaurant implements RestaurantRepository
func (r restaurantRepository) GetRestaurant(id primitive.ObjectID) (*model.Restaurant, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var restaurant *model.Restaurant
	// fmt.Println("id = %s", id)

	err := r.restaurantCollection.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&restaurant)
	if err != nil {
		return nil, err
	}
	return restaurant, nil
}

// CreateRestaurant implements RestaurantRepository
func (r restaurantRepository) CreateRestaurant(restaurant model.Restaurant) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := r.restaurantCollection.InsertOne(ctx, restaurant)
	return result, err
}

// EditRestaurant implements RestaurantRepository
func (r restaurantRepository) EditRestaurant(id primitive.ObjectID, restaurant *model.Restaurant) (*model.Restaurant, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: restaurant}}

	var result model.Restaurant
	err := r.restaurantCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteRestaurant implements RestaurantRepository
func (r restaurantRepository) DeleteRestaurant(id primitive.ObjectID) (*model.Restaurant, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: id}}

	var result model.Restaurant
	err := r.restaurantCollection.FindOneAndDelete(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
