package main

// Tutorial: https://www.jetbrains.com/guide/go/tutorials/rest_api_series/gin/

import (
	"github.com/gin-gonic/gin"
	"github.com/go_examples/gin-api/sample/recipes"
	"log"
)

func main() {
	// create gin router
	router := gin.Default()

	// instantiate recipe handler and provide a data store
	store := recipes.NewMemStore()
	recipesHandler := NewRecipesHandler(store)

	// register routes
	router.GET("/", homepage)
	router.GET("/recipes", recipesHandler.ListRecipes)
	router.POST("/recipes", recipesHandler.CreateRecipe)
	router.GET("/recipes/:id", recipesHandler.GetRecipe)
	router.PUT("/recipes/:id", recipesHandler.UpdateRecipe)
	router.DELETE("/recipes/:id", recipesHandler.DeleteRecipe)

	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}
