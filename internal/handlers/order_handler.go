package handlers

import (
	"net/http"
	"makerspace-api/models"
	"github.com/gin-gonic/gin"
)

var orders = []models.Order{

	{ID: "1", Name: "Bre", File: "bru"},
	{ID: "2", Name: "Bre2", File: "bru2"},
	{ID: "3", Name: "Bre3", File: "bru3"},

}

// function to return all orders from temp var made above :)
func GetOrders(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, orders)

}

func GetOrderByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of orders, looking for
    // an order whose ID value matches the parameter.
    for _, a := range orders {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Order not found"})
}

func AddOrder(c *gin.Context) {
	var newOrder models.Order

	if err := c.BindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	orders = append(orders, newOrder)
	c.JSON(http.StatusCreated, newOrder)
}
