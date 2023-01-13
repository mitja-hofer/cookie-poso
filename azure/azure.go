package azure

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"log"
	"os"
)

func GetCredential() *azidentity.DefaultAzureCredential {
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatal("Invalid credentials with error: ", err.Error())
	}
	return credential
}

func GetBlobClient(credential *azidentity.DefaultAzureCredential) *azblob.Client {
	url := os.Getenv("AZURE_STORAGE_ACCOUNT")
	client, err := azblob.NewClient(url, credential, nil)
	if err != nil {
		log.Fatal("Invalid credentials with error: ", err.Error())
	}
	return client
}
