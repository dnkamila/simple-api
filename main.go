package main

import (
	_ "github.com/joho/godotenv/autoload"
	"net"
	"net/http"
	"os"
	"simple-api/application"
)

func main() {
	println("main.main()")
	address := net.JoinHostPort(os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))

	app := application.NewApp()
	app.InitRouter()

	http.ListenAndServe(address, app.Router)
}
