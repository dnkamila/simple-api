package application

import (
	"simple-api/controllers"
	. "simple-api/database"
	. "simple-api/routers"
)

type App struct {
	Router Router
	DB     DB
}

func NewApp() *App {
	return &App{GetRouter(), GetDB()}
}

func (app *App) InitRouter() {
	app.Router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	app.Router.HandleFunc("/user/id/{id}", controllers.GetUserById).Methods("GET")
	app.Router.HandleFunc("/user/username/{username}", controllers.GetUserByUsername).Methods("GET")
	app.Router.HandleFunc("/user/id", controllers.UpdateUserPasswordById).Methods("PUT")
	app.Router.HandleFunc("/user/username", controllers.UpdateUserPasswordByUsername).Methods("PUT")
	app.Router.HandleFunc("/user/id/{id}", controllers.DeleteUserById).Methods("DELETE")
	app.Router.HandleFunc("/user/username/{username}", controllers.DeleteUserByUsername).Methods("DELETE")
}
