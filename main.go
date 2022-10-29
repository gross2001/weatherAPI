package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gross2001/weather/api"
)

type city struct {
	City string `json:"city"`
}

func main() {
	router := gin.Default()
	router.POST("/city", postCity)
	router.GET("/city/:name", getCityByName)

	router.Run("localhost:8080")
}

func postCity(c *gin.Context) {
	var newCity city
	if err := c.BindJSON(&newCity); err != nil {
		log.Println(err.Error())
		return
	}
	cities, err := api.GetCoordinates(newCity.City)
	if err != nil {
		log.Println(err.Error())
		return
	}
	weather, err := api.GetWeather(cities[0])
	if err != nil {
		log.Println(err.Error())
		return
	}
	c.IndentedJSON(http.StatusCreated, weather)
}

func getCityByName(c *gin.Context) {
	city := c.Param("name")
	log.Println(city)
	cities, err := api.GetCoordinates(city)
	if err != nil {
		log.Println(err.Error())
		return
	}
	weather, err := api.GetWeather(cities[0])
	if err != nil {
		log.Println(err.Error())
		return
	}
	c.IndentedJSON(http.StatusCreated, weather)
}
