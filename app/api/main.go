package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/mxzinke/crawler-storage-api/internal/storage"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("unable to parse .env file")
	}
}

func main() {
	s3Client := storage.New()

	log.Printf("%v", s3Client)
}
