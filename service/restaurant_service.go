package service

import (
	"swarch/poptum/restaurants/model"
	"swarch/poptum/restaurants/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RestaurantService interface {
	GetRestaurants() ([]model.Restaurant, error)
	GetRestaurant(string) (*model.Restaurant, error)
	CreateRestaurant(model.Restaurant) (*mongo.InsertOneResult, error)
	EditRestaurant(string, *model.Restaurant) (*model.Restaurant, error)
	DeleteRestaurant(string) (*model.Restaurant, error)
}

type restaurantService struct {
	restaurantRepository repository.RestaurantRepository
}

// NewRestaurantService -> returns new restaurant service
func NewRestaurantService(r repository.RestaurantRepository) RestaurantService {
	return restaurantService{
		restaurantRepository: r,
	}
}

func (r restaurantService) GetRestaurants() ([]model.Restaurant, error) {
	return r.restaurantRepository.GetRestaurants()
}

func (r restaurantService) GetRestaurant(id string) (*model.Restaurant, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return r.restaurantRepository.GetRestaurant(objectId)
}

func (r restaurantService) CreateRestaurant(restaurant model.Restaurant) (*mongo.InsertOneResult, error) {
	result, err := r.restaurantRepository.CreateRestaurant(restaurant)
	return result, err
}

func (r restaurantService) EditRestaurant(id string, restaurant *model.Restaurant) (*model.Restaurant, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return r.restaurantRepository.EditRestaurant(objectId, restaurant)
}

func (r restaurantService) DeleteRestaurant(id string) (*model.Restaurant, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return r.restaurantRepository.DeleteRestaurant(objectId)
}
