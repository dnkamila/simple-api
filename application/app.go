package application

import (
	"net/http"
	"simple-api/controllers"
	"simple-api/database"
	. "simple-api/database"
	"simple-api/routers"
	. "simple-api/routers"
)

type App struct {
	Router Router
	DB     DB
}

func NewApp() *App {
	return &App{routers.GetRouter(), database.GetDB()}
}

func (app *App) InitDB() {

}

func (app *App) InitRouter() {
	app.Router.HandleFunc("/", index)
	app.Router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	app.Router.HandleFunc("/user/id/{id}", controllers.GetUserById).Methods("GET")
	app.Router.HandleFunc("/user/username/{username}", controllers.GetUserByUsername).Methods("GET")
	app.Router.HandleFunc("/user", controllers.UpdateUserPasswordById).Methods("PUT")
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}
