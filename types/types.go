package types

type Account struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Recipe struct {
	Id     int64  `json:"id"`
	UserId int64  `json:"userId"`
	Name   string `json:"name"`
	Text   string `json:"text"`
}
