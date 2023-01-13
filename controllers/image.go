package controllers

import (
	"CookiePoso/globals"
	"bytes"
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func randomString() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(r.Int())
}

func UploadImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		containerName := os.Getenv("AZURE_BLOB_CONTAINER")
		ctx := context.Background()
		contentType := "image/jpg"

		file, header, err := c.Request.FormFile("upload")
		filename := header.Filename
		buf := bytes.NewBuffer(nil)
		_, err = io.Copy(buf, file)
		if err != nil {
			log.Println(err)
		}
		blobName := filename + "-" + randomString()

		options := &azblob.UploadBufferOptions{
			HTTPHeaders: &blob.HTTPHeaders{
				BlobContentType: &contentType,
			},
		}

		_, err = globals.AzBlobClient.UploadBuffer(ctx, containerName, blobName, buf.Bytes(), options)
		if err != nil {
			log.Fatalf("Failure to upload to blob: %+v", err)
		}

	}
}
