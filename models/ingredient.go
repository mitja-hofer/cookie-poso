package models

import (
	"CookiePoso/globals"
	"database/sql"
	"fmt"
	"log"
)

type Ingredient struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type IngredientInRecipe struct {
	Id       int64  `json:"id" validate:"omitempty"`
	RecipeId int64  `json:"recipe-id" validate:"omitempty"`
	Name     string `json:"name"`
	Amount   int64  `json:"amount"`
	Unit     string `json:"unit"`
}

func AddIngredient(ingredient Ingredient) (int64, error) {
	res, err := globals.DB.Exec("INSERT INTO ingredient name VALUES ?", ingredient.Name)
	if err != nil {
		return 0, fmt.Errorf("Add recipe: %v", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Add recipe: %v", err)
	}
	ingredient.Id = id
	log.Println(ingredient)
	return id, nil

}

func IngredientsByRecipeId(recipeId int64) ([]IngredientInRecipe, error) {
	var ingredients []IngredientInRecipe

	res, err := globals.DB.Query("SELECT i.id, i.name, ir.amount, ir.unit FROM cookie.ingredientrecipe ir JOIN cookie.recipe r ON ir.recipeId = r.id JOIN cookie.ingredient i ON ir.ingredientId = i.id WHERE r.id = ?", recipeId)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	for res.Next() {
		var ingredient IngredientInRecipe
		if err := res.Scan(&ingredient.Id, &ingredient.Name, &ingredient.Amount, &ingredient.Unit); err != nil {
			return nil, err
		}
		println(ingredient.Name)
		ingredients = append(ingredients, ingredient)
	}
	if err := res.Err(); err != nil {
		return nil, err
	}

	return ingredients, nil
}

func SelectIngredientsByPartialName(name string) ([]Ingredient, error) {
	var ingredients []Ingredient

	res, err := globals.DB.Query("SELECT id, name from ingredient WHERE name like '%?%'", name)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	for res.Next() {
		var ingredient Ingredient
		if err := res.Scan(&ingredient.Id, &ingredient.Name); err != nil {
			return nil, err
		}
		ingredients = append(ingredients, ingredient)
	}
	if err := res.Err(); err != nil {
		return nil, err
	}
	return ingredients, nil
}

func InsertIngredient(ingredient Ingredient) (int64, error) {
	res, err := globals.DB.Exec("INSERT INTO ingredient (name) VALUES (?)", ingredient.Name)
	if err != nil {
		return 0, fmt.Errorf("Add ingredient: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Add ingredient: %v", err)
	}
	return id, nil
}

func SelectIngredientById(id int64) (*Ingredient, error) {
	var ingredient Ingredient
	err := globals.DB.QueryRow("SELECT id, name FROM ingredient WHERE id = ?", id).Scan(&ingredient.Id, &ingredient.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &ingredient, nil
}

func SelectIngredientByName(name string) (*Ingredient, error) {
	var ingredient Ingredient
	err := globals.DB.QueryRow("SELECT id, name FROM ingredient WHERE name = ?", name).Scan(&ingredient.Id, &ingredient.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &ingredient, nil
}

func InsertIngredientRecipe(ingredient IngredientInRecipe, recipeId int64) (int64, error) {
	log.Println("inserting ir table", ingredient, recipeId)
	res, err := globals.DB.Exec("INSERT INTO ingredientrecipe (`ingredientId`, `recipeId`, `amount`, `unit`) VALUES (?, ?, ?, ?);", ingredient.Id, recipeId, ingredient.Amount, ingredient.Unit)
	log.Println(res, err)
	if err != nil {
		log.Println(err)
		return 0, fmt.Errorf("Add ingredientrecipe: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Add ingredientrecipe: %v", err)
	}

	return id, nil
}

func InsertIngredients(ingredients []IngredientInRecipe) ([]IngredientInRecipe, error) {
	var ingredientsWithIds []IngredientInRecipe
	for _, ingredient := range ingredients {
		if ingredient.Id != 0 {
			i, err := SelectIngredientById(ingredient.Id)
			if err != nil {
				log.Println(err)
				return nil, err
			}

			if i.Name == ingredient.Name {
				// ingredient exists and has correct id
				ingredientsWithIds = append(ingredientsWithIds, ingredient)

			} else {
				return nil, fmt.Errorf("ingredient with id %d exists, but has a different name. Given: %s, actual: %s", ingredient.Id, ingredient.Name, i.Name)
			}
		} else {
			i, err := SelectIngredientByName(ingredient.Name)
			if err != nil {
				return nil, err
			}
			if i != nil {
				ingredient.Id = i.Id
				ingredientsWithIds = append(ingredientsWithIds, ingredient)
			} else {
				id, err := InsertIngredient(Ingredient{0, ingredient.Name})
				if err != nil {
					return nil, err
				}
				ingredient.Id = id
				ingredientsWithIds = append(ingredientsWithIds, ingredient)
			}
		}
	}
	return ingredientsWithIds, nil
}
