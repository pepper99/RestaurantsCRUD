package routes

import (
	"swarch/poptum/restaurants/controller"
	"swarch/poptum/restaurants/repository"
	"swarch/poptum/restaurants/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type routes struct {
	router *gin.Engine
}

func SetupRoutes(mongoClient *mongo.Client) {
	httpRouter := routes{
		router: gin.Default(),
	}

	apiRouter := httpRouter.router.Group("/api/v1")
	httpRouter.AddRestaurantRoutes(apiRouter, mongoClient)

	httpRouter.router.Run(":5000")
}

func (r routes) AddRestaurantRoutes(rg *gin.RouterGroup, mongoClient *mongo.Client) {
	restaurantRepository := repository.NewRestaurantRepository(mongoClient)
	restaurantService := service.NewRestaurantService(restaurantRepository)
	restaurantController := controller.NewRestaurantController(restaurantService)
	restaurantRouter := rg.Group("restaurants")

	restaurantRouter.GET("/", restaurantController.GetRestaurants)
	restaurantRouter.GET("/:id", restaurantController.GetRestaurant)
	restaurantRouter.POST("/", restaurantController.CreateRestaurant)
	restaurantRouter.PUT("/:id", restaurantController.EditRestaurant)
	restaurantRouter.DELETE("/:id", restaurantController.DeleteRestaurant)
}
