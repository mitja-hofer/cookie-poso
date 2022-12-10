package models

import (
	"CookiePoso/globals"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CheckUserPass(username, password string) (int64, bool) {
	var user User
	res := globals.DB.QueryRow("SELECT id, username, password, email FROM user WHERE username = ?", username)

	err := res.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	if err != nil {
		/*
			if err == sql.ErrNoRows {
				return false, fmt.Errorf("no such user: %s", username)
			}
		*/
		log.Fatal(err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println("wrong password")
		return 0, false
	}
	return user.Id, true
}

func AddUser(account User) (int64, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(account.Password), 8)
	if err != nil {
		return 0, fmt.Errorf("Add user: %v", err)
	}
	res, err := globals.DB.Exec("INSERT INTO user (username, password, email) VALUES (?, ?, ?)", account.Username, hashed, account.Email)
	if err != nil {
		return 0, fmt.Errorf("Add user: %v", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Add user: %v", err)
	}
	return id, nil
}

func SelectUserByID(userId int64) (User, error) {
	var user User

	res := globals.DB.QueryRow("SELECT id, username, email FROM user WHERE id = ?", userId)
	err := res.Scan(&user.Id, &user.Username, &user.Email)
	return user, err
}

func SelectUserByUsername(username string) (User, error) {
	var user User

	res := globals.DB.QueryRow("SELECT id, us*ername, email FROM user WHERE username = ?", username)
	err := res.Scan(&user.Id, &user.Username, &user.Email)
	return user, err
}
