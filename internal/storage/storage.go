package storage

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/minio/minio-go"
)

// Client holds data about the connection to the S3 storage.
type Client struct {
	actor   *minio.Client
	buckets []*Bucket
}

// Bucket holds the data about the bucket to do actions on it.
type Bucket struct {
	name    string
	storage Client
}

// New creates a new client to access the s3 storage.
func New() *Client {
	err := checkRequiredEnv()
	if err != nil {
		log.Fatalln(err)
	}

	endpoint := os.Getenv("STORAGE_ENDPOINT")
	accessKeyID := os.Getenv("STORAGE_ACCESS_KEY_ID")
	accessSecret := os.Getenv("STORAGE_ACCESS_SECRET")
	useSSL := os.Getenv("STORAGE_USE_SSL") == "true" || os.Getenv("STORAGE_USE_SSL") == "on"

	s3Client, err := minio.New(endpoint, accessKeyID, accessSecret, useSSL)

	if err != nil {
		log.Fatalln(err)
	}

	return &Client{
		actor: s3Client,
	}
}

// RequiredEnv contains all required env-variables for this package.
var RequiredEnv = [...]string{
	"STORAGE_ENDPOINT",
	"STORAGE_ACCESS_KEY_ID",
	"STORAGE_ACCESS_SECRET",
	"STORAGE_USE_SSL",
}

func checkRequiredEnv() error {
	var missingEnv error
	for _, value := range RequiredEnv {
		_, exists := os.LookupEnv(value)
		if !exists {
			errorMessage := fmt.Sprintf("The Env variable %s is missing, please define it.", value)
			missingEnv = errors.New(errorMessage)
			break
		}
	}

	return missingEnv
}
