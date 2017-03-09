package routers

import "github.com/gorilla/mux"

type Router struct {
	*mux.Router
}

var router Router

func init() {
	println("router.init()")
	router = Router{mux.NewRouter()}
}

func GetRouter() Router {
	return router
}
