package globals

import (
	"CookiePoso/azure"
	"CookiePoso/database"
)

var DB = database.GetConnection()
var AzCredential = azure.GetCredential()
var AzBlobClient = azure.GetBlobClient(AzCredential)
