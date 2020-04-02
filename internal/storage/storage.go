package storage

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/minio/minio-go"
)

// S3Client holds data about the connection to the S3 storage.
type S3Client struct {
	actor *minio.Client
}

// New creates a new client to access the s3 storage.
func New() *S3Client {
	err := checkRequiredEnv()
	if err != nil {
		log.Fatalln(err)
	}

	endpoint := os.Getenv("STORAGE_ENDPOINT")
	accessKeyID := os.Getenv("STORAGE_ACCESS_KEY_ID")
	accessSecret := os.Getenv("STORAGE_ACCESS_SECRET")
	useSSL := os.Getenv("STORAGE_USE_SSL") == "true" || os.Getenv("STORAGE_USE_SSL") == "on"
	bucketName := os.Getenv("STORAGE_BUCKET")

	minioClient, err := minio.New(endpoint, accessKeyID, accessSecret, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	if found, _ := minioClient.BucketExists(bucketName); !found {
		log.Fatalf("Bucket %s not found!", bucketName)
	}

	return &S3Client{
		actor: minioClient,
	}
}

// RequiredEnv contains all required env-variables for this package.
var RequiredEnv = [...]string{
	"STORAGE_ENDPOINT",
	"STORAGE_ACCESS_KEY_ID",
	"STORAGE_ACCESS_SECRET",
	"STORAGE_USE_SSL",
	"STORAGE_BUCKET",
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
