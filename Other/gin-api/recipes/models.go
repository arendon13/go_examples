package recipes

// Recipe represents a recipes
type Recipe struct {
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
}

// Ingredient represents individual ingredients
type Ingredient struct {
	Name string `json:"name"`
}
