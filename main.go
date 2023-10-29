package main

import (
	"os"

	"github.com/Prabhdeep52/golang-restaurent-management/database"
	"github.com/Prabhdeep52/golang-restaurent-management/routes"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodConnection *mongo.Collection = database.openCollection(database.Client, "food")

func main() {
	Collection(database.Client, "food")
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)

}
