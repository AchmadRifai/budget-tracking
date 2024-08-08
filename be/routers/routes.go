package routers

import (
	"be/controllers"

	"github.com/gorilla/mux"
)

func Routing(r *mux.Router) {
	r.HandleFunc("/hello", controllers.Hello).Methods("GET")
}
