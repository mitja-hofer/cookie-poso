package globals

import (
	"CookiePoso/database"
)

var DB = database.GetConnection()
