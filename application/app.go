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
	userRouter := app.Router.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/", controllers.CreateUser).Methods("POST")
	userRouter.HandleFunc("/id/{id}", controllers.GetUserById).Methods("GET")
	userRouter.HandleFunc("/username/{username}", controllers.GetUserByUsername).Methods("GET")
	userRouter.HandleFunc("id", controllers.UpdateUserPasswordById).Methods("PUT")
	userRouter.HandleFunc("/username", controllers.UpdateUserPasswordByUsername).Methods("PUT")
	userRouter.HandleFunc("/id/{id}", controllers.DeleteUserById).Methods("DELETE")
	userRouter.HandleFunc("/username/{username}", controllers.DeleteUserByUsername).Methods("DELETE")

	tokenRouter := app.Router.PathPrefix("/token").Subrouter()
	tokenRouter.HandleFunc("/", controllers.UpdateToken).Methods("PUT")

	dummyRouter := app.Router.PathPrefix("/dummy").Subrouter()
	dummyRouter.HandleFunc("/", controllers.Dummy).Methods("GET")
}
