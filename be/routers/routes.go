package routers

import (
	"be/controllers"
	"be/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func Routing(r *mux.Router) {
	r.Use(middlewares.Logging)
	r.Use(middlewares.Authorization)
	r.HandleFunc("/hello", controllers.Hello).Methods("GET")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	dashboardRouting(r)
	r.NotFoundHandler = http.HandlerFunc(controllers.NotFound)
}

func dashboardRouting(r *mux.Router) {
	r.HandleFunc("/dashboard/logout", controllers.DashboardLogout).Methods("GET")
	r.HandleFunc("/dashboard/chart", controllers.GetChart).Methods("GET")

	r.HandleFunc("/dashboard/admin/users", controllers.AllUsers).Methods("GET")
	r.HandleFunc("/dashboard/admin/users/{id}", controllers.DelUser).Methods("DELETE")
}
