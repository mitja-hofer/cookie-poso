package helpers

import (
	"CookiePoso/globals"
	"CookiePoso/types"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
)

func CheckUserPass(username, password string) (int64, bool) {
	var account types.Account
	res := globals.DB.QueryRow("SELECT id, username, password, email FROM account WHERE username = ?", username)

	err := res.Scan(&account.Id, &account.Username, &account.Password, &account.Email)
	if err != nil {
		/*
			if err == sql.ErrNoRows {
				return false, fmt.Errorf("no such user: %s", username)
			}
		*/
		log.Fatal(err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil {
		log.Println("wrong password")
		return 0, false
	}
	return account.Id, true
}

func AddUser(account types.Account) (int64, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(account.Password), 8)
	if err != nil {
		return 0, fmt.Errorf("Add user: %v", err)
	}
	res, err := globals.DB.Exec("INSERT INTO account (username, password, email) VALUES (?, ?, ?)", account.Username, hashed, account.Email)
	if err != nil {
		return 0, fmt.Errorf("Add user: %v", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Add user: %v", err)
	}
	return id, nil
}

func AddRecipe(recipe types.Recipe) (int64, error) {
	res, err := globals.DB.Exec("INSERT INTO recipe (userId, name, text) VALUES (?, ?, ?)", recipe.UserId, recipe.Name, recipe.Text)
	if err != nil {
		return 0, fmt.Errorf("Add recipe: %v", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Add recipe: %v", err)
	}
	recipe.Id = id
	log.Println(recipe)
	return id, nil

}

func EmptyUserPass(username, password string) bool {
	return strings.Trim(username, " ") == "" || strings.Trim(password, " ") == ""
}

func EmptyUserPassEmail(username, password, email string) bool {
	return strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" || strings.Trim(email, " ") == ""
}
