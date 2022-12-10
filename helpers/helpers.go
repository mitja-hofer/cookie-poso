package helpers

import (
	"strings"
)

func EmptyUserPass(username, password string) bool {
	return strings.Trim(username, " ") == "" || strings.Trim(password, " ") == ""
}

func EmptyUserPassEmail(username, password, email string) bool {
	return strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" || strings.Trim(email, " ") == ""
}
