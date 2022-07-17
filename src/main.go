package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"alteacare/golang-basecode/src/drivers/cloudwatch"
	"alteacare/golang-basecode/src/drivers/gorm"
	"alteacare/golang-basecode/src/interfaces/http"
	http_internal "alteacare/golang-basecode/src/interfaces/http-internal"
)

func main() {
	errEnv := godotenv.Load()
	db, errDb := gorm.Connect()
	appInterface := os.Getenv("INTERFACE")

	if errDb != nil {
		log.Panicln("Failed to Initialized DB:", errDb)
	}

	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	if appInterface == "" {
		log.Fatal("Interace not found")
	}

	if appInterface == "http" {
		h := http.New(&http.Http{
			DB:         db,
			Cloudwatch: cloudwatch.InitCloudwatch(),
		})
		h.Launch()
	}

	if appInterface == "http-internal" {
		hi := http_internal.New(&http_internal.HttpInternal{
			DB:         db,
			Cloudwatch: cloudwatch.InitCloudwatch(),
		})
		hi.Launch()
	}

	if appInterface != "" {
		log.Fatalf(`Interface not found (%v)`, appInterface)
	}
}
