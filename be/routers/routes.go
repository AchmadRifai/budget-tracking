package routers

import (
	"be/controllers"
	"be/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func Routing(r *mux.Router) {
	r.Use(mux.CORSMethodMiddleware(r))
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
	r.HandleFunc("/dashboard/budget", controllers.AllBudget).Methods("GET")
	r.HandleFunc("/dashboard/budget", controllers.AddBudget).Methods("POST")
	r.HandleFunc("/dashboard/budget/{id}", controllers.EditBudget).Methods("PUT")
	r.HandleFunc("/dashboard/budget/{id}", controllers.DelBudget).Methods("DELETE")
	r.HandleFunc("/dashboard/category", controllers.AllCategory).Methods("GET")
	r.HandleFunc("/dashboard/category", controllers.AddCategory).Methods("POST")
	r.HandleFunc("/dashboard/category/{id}", controllers.EditCategory).Methods("PUT")
	r.HandleFunc("/dashboard/category/{id}", controllers.DelCategory).Methods("DELETE")
	r.HandleFunc("/dashboard/expenses", controllers.AllExpenses).Methods("GET")
	r.HandleFunc("/dashboard/expenses", controllers.AddExpenses).Methods("POST")
	r.HandleFunc("/dashboard/expenses/{id}", controllers.EditExpenses).Methods("PUT")
	r.HandleFunc("/dashboard/expenses/{id}", controllers.DelExpenses).Methods("DELETE")

	r.HandleFunc("/dashboard/admin/users", controllers.AllUsers).Methods("GET")
	r.HandleFunc("/dashboard/admin/users/{id}", controllers.DelUser).Methods("DELETE")
}
