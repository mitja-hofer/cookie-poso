package models

import (
	"CookiePoso/globals"
	"database/sql"
	"fmt"
	"log"
)

type Recipe struct {
	Id          int64                `json:"id"`
	UserId      int64                `json:"user-id"`
	Name        string               `json:"name"`
	Text        string               `json:"text"`
	Ingredients []IngredientInRecipe `json:"ingredients"`
}

func AddRecipe(recipe Recipe) (int64, error) {
	res, err := globals.DB.Exec("INSERT INTO recipe (userId, name, text) VALUES (?, ?, ?)", recipe.UserId, recipe.Name, recipe.Text)
	if err != nil {
		return 0, fmt.Errorf("Add recipe: %v", err)
	}
	recipe.Id, err = res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Add recipe: %v", err)
	}
	log.Println("here", recipe.Ingredients)

	recipe.Ingredients, err = InsertIngredients(recipe.Ingredients)
	for _, ingredient := range recipe.Ingredients {
		_, err = InsertIngredientRecipe(ingredient, recipe.Id)
		if err != nil {
			return 0, err
		}
	}

	return recipe.Id, nil

}

func SelectRecipesByUserId(userId int64) ([]Recipe, error) {
	var recipes []Recipe

	res, err := globals.DB.Query("SELECT id, userId, name, text FROM recipe WHERE userId = ?", userId)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	for res.Next() {
		var recipe Recipe
		var ingredients []IngredientInRecipe
		if err := res.Scan(&recipe.Id, &recipe.UserId, &recipe.Name, &recipe.Text); err != nil {
			return nil, err
		}

		ingredients, err = IngredientsByRecipeId(recipe.Id)
		if err != nil {
			println(ingredients)
			return nil, err
		}

		recipe.Ingredients = ingredients
		recipes = append(recipes, recipe)

	}
	if err := res.Err(); err != nil {
		return nil, err
	}

	return recipes, nil
}

func SelectRecipeByID(id int64) (*Recipe, error) {
	var recipe Recipe

	err := globals.DB.QueryRow("SELECT id, userId, name, text FROM recipe WHERE id = ?", id).Scan(&recipe.Id, &recipe.UserId, &recipe.Name, &recipe.Text)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	ingredients, err := IngredientsByRecipeId(recipe.Id)
	if err != nil {
		return nil, err
	}
	recipe.Ingredients = ingredients

	return &recipe, nil
}

func SelectRecipesByIngredient(name string) ([]Recipe, error) {
	var recipes []Recipe
	var id int64
	res, err := globals.DB.Query("SELECT r.id FROM ingredientrecipe ir JOIN recipe r ON ir.recipeId = r.id JOIN ingredient i ON ir.ingredientId = i.id WHERE i.name = ?", name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer res.Close()

	for res.Next() {
		res.Scan(&id)
		recipe, err := SelectRecipeByID(id)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, *recipe)
	}
	if err := res.Err(); err != nil {
		return nil, err
	}
	return recipes, nil
}

func SelectRecipeByName(name string) (*Recipe, error) {
	var recipe Recipe

	err := globals.DB.QueryRow("SELECT id, userId, name, text FROM recipe WHERE name = ?", name).Scan(&recipe.Id, &recipe.UserId, &recipe.Name, &recipe.Text)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	ingredients, err := IngredientsByRecipeId(recipe.Id)
	if err != nil {
		return nil, err
	}
	recipe.Ingredients = ingredients

	return &recipe, nil
}

func SelectRecipesByPartialName(name string) ([]Recipe, error) {
	var recipes []Recipe

	res, err := globals.DB.Query("SELECT id, userId, name, text FROM recipe WHERE name like ?", "%"+name+"%")
	if err != nil {
		log.Println("error executing name like query", err)
		return nil, err
	}
	defer res.Close()

	for res.Next() {
		var recipe Recipe
		var ingredients []IngredientInRecipe
		if err := res.Scan(&recipe.Id, &recipe.UserId, &recipe.Name, &recipe.Text); err != nil {
			log.Println(err)
			return nil, err
		}

		ingredients, err = IngredientsByRecipeId(recipe.Id)
		if err != nil {
			return nil, err
		}

		recipe.Ingredients = ingredients
		log.Println(recipe)
		recipes = append(recipes, recipe)

	}
	if err := res.Err(); err != nil {
		return nil, err
	}

	return recipes, nil
}
