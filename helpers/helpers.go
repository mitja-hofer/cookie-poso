package helpers

import (
	"CookiePoso/globals"
	"CookiePoso/types"
	"log"
	"strings"
)

func CheckUserPass(username, password string) bool {
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

	return account.Password == password
}

func EmptyUserPass(username, password string) bool {
	return strings.Trim(username, " ") == "" || strings.Trim(password, " ") == ""
}
