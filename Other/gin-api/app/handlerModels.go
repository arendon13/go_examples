package main

import "github.com/go_examples/gin-api/sample/recipes"

type RecipesHandler struct {
	store recipeStore
}

// recipeStore is an interface for the data store
type recipeStore interface {
	Add(name string, recipe recipes.Recipe) error
	Get(name string) (recipes.Recipe, error)
	List() (map[string]recipes.Recipe, error)
	Update(name string, recipe recipes.Recipe) error
	Remove(name string) error
}

func NewRecipesHandler(rs recipeStore) *RecipesHandler {
	return &RecipesHandler{
		store: rs,
	}
}
