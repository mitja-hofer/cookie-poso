package globals

import (
	"CookiePoso/database"
)

var Secret = []byte("secret")
var DB = database.GetConnection()

const Userkey = "user"
