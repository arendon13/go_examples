package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go_examples/gin-api/sample/recipes"
	"github.com/gosimple/slug"
	"net/http"
)

func homepage(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the homepage!")
}

func (h RecipesHandler) CreateRecipe(c *gin.Context) {
	// Get request body and convert it to recipes.Recipe
	var recipe recipes.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create a url friendly name
	id := slug.Make(recipe.Name)

	// add to the store
	err := h.store.Add(id, recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return success payload
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h RecipesHandler) ListRecipes(c *gin.Context) {
	// retrieve list of all recipes
	recipesList, err := h.store.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return list with ok status
	c.JSON(http.StatusOK, recipesList)
}

func (h RecipesHandler) GetRecipe(c *gin.Context) {
	// extract id from query param
	id := c.Param("id")

	// retrieve recipe
	recipe, err := h.store.Get(id)
	if err != nil {
		// return not found status if we cannot find the recipe
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// return recipe
	c.JSON(http.StatusOK, recipe)
}

func (h RecipesHandler) UpdateRecipe(c *gin.Context) {
	// Get request body and convert it to recipes.Recipe
	var recipe recipes.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// extract id from query param
	id := c.Param("id")

	// update recipe
	err := h.store.Update(id, recipe)
	if err != nil {
		if errors.Is(err, recipes.NotFoundErr) {
			// return not found status if we cannot find the recipe
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return success payload
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h RecipesHandler) DeleteRecipe(c *gin.Context) {
	// extract id from query param
	id := c.Param("id")

	// delete recipe
	err := h.store.Remove(id)
	if err != nil {
		if errors.Is(err, recipes.NotFoundErr) {
			// return not found status if we cannot find the recipe
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return success payload
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
