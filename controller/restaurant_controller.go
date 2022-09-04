package controller

import (
	"net/http"
	"swarch/poptum/restaurants/model"
	"swarch/poptum/restaurants/service"

	"github.com/gin-gonic/gin"
)

// RestaurantController : represent the restaurant's controller contract
type RestaurantController interface {
	GetRestaurant(*gin.Context)
	GetRestaurants(*gin.Context)
	CreateRestaurant(*gin.Context)
	EditRestaurant(*gin.Context)
	DeleteRestaurant(*gin.Context)
}

type restaurantController struct {
	restaurantService service.RestaurantService
}

// NewRestaurantController -> returns new restaurant controller
func NewRestaurantController(s service.RestaurantService) RestaurantController {
	return restaurantController{
		restaurantService: s,
	}
}

func (r restaurantController) GetRestaurants(c *gin.Context) {
	res, err := r.restaurantService.GetRestaurants()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	c.JSON(http.StatusOK,
		model.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}},
	)
}

func (r restaurantController) GetRestaurant(c *gin.Context) {
	id := c.Param("id")

	res, err := r.restaurantService.GetRestaurant(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	c.JSON(http.StatusOK,
		model.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}},
	)
}

func (r restaurantController) CreateRestaurant(c *gin.Context) {
	var newRestaurant model.Restaurant
	if err := c.BindJSON(&newRestaurant); err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	res, err := r.restaurantService.CreateRestaurant(newRestaurant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusCreated, model.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": res}})
}

func (r restaurantController) EditRestaurant(c *gin.Context) {
	id := c.Param("id")
	var restaurant model.Restaurant
	if err := c.BindJSON(&restaurant); err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	res, err := r.restaurantService.EditRestaurant(id, &restaurant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	c.JSON(http.StatusOK,
		model.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}},
	)
}

func (r restaurantController) DeleteRestaurant(c *gin.Context) {
	id := c.Param("id")

	res, err := r.restaurantService.DeleteRestaurant(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	c.JSON(http.StatusOK,
		model.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": res}},
	)
}
