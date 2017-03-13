package main

import (
	_ "github.com/joho/godotenv/autoload"
	"log"
	"net"
	"net/http"
	"os"
	"simple-api/application"
	"simple-api/helpers"
)

func main() {
	err := helpers.InitPPKeyResource()
	if err != nil {
		log.Fatal("Could not create key resource. Error: %v\n", err)
	}

	address := net.JoinHostPort(os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))

	app := application.NewApp()
	app.InitRouter()

	http.ListenAndServe(address, app.Router)
}
