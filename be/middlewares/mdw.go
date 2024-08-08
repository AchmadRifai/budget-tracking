package middlewares

import (
	"be/db"
	"be/dtos"
	errorhandlers "be/errorHandlers"
	"be/models"
	myutils "be/myUtils"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Println("Started", r.Method, "On", r.RequestURI)
		next.ServeHTTP(w, r)
		log.Println("Completed", r.Method, "On", r.RequestURI, "in", time.Since(start))
	})
}

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer errorhandlers.AuthErrorRest(w, r)
		if strings.HasPrefix(r.RequestURI, "/dashboard/") {
			user := myutils.GetUser(r)
			if strings.HasPrefix(r.RequestURI, "/dashboard/admin/") {
				conn := db.DbConnect()
				var role models.Role
				result := conn.Where("id=?", user.RoleId).First(&role)
				if result.Error != nil {
					panic(result.Error)
				}
				if result.RowsAffected == 0 {
					panic(errors.New("role not found"))
				}
				if role.Name != "Admin" {
					panic(errors.New("not allowed"))
				}
			}
		}
		next.ServeHTTP(w, r)
	})
}

func ReqBodyValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer errorhandlers.BadReqRest(w, r)
		if r.Method == "POST" && r.RequestURI == "/login" {
			login(r)
		} else if r.Method == "POST" && r.RequestURI == "/register" {
			registration(r)
		}
		next.ServeHTTP(w, r)
	})
}

func registration(r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var req dtos.RegRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		panic(err)
	}
	if req.Username == "" {
		panic(errors.New("username is required"))
	}
	if req.Password == "" {
		panic(errors.New("password is required"))
	}
	if req.Email == "" || !myutils.IsValidEmail(req.Email) {
		panic(errors.New("email is required"))
	}
	if req.FullName == "" {
		panic(errors.New("fullname is required"))
	}
}

func login(r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println("Req Login", string(body))
	var req dtos.LoginRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		panic(err)
	}
	if req.Username == "" {
		panic(errors.New("username is required"))
	}
	if req.Password == "" {
		panic(errors.New("password is required"))
	}
}
