package controllers

import (
	"be/db"
	"be/dtos"
	errorhandlers "be/errorHandlers"
	"be/models"
	myutils "be/myUtils"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"gorm.io/gorm"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	defer errorhandlers.NormalErrorRest(w, r)
	myutils.SendJson(w, map[string]interface{}{"message": "hello world"}, 200)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	defer errorhandlers.NormalErrorRest(w, r)
	myutils.SendJson(w, map[string]interface{}{"message": "Not Found"}, 404)
}

func Login(w http.ResponseWriter, r *http.Request) {
	defer errorhandlers.NormalErrorRest(w, r)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var req dtos.LoginRequest
	if err = json.Unmarshal(body, &req); err != nil {
		panic(err)
	}
	if req.Username == "" {
		panic(errors.New("username is required"))
	}
	if req.Password == "" {
		panic(errors.New("password is required"))
	}
	conn := db.DbConnect()
	var user models.User
	query := "email=? AND password=?"
	if !myutils.IsValidEmail(req.Username) {
		query = "user_name=? AND password=?"
	}
	result := conn.Where(query, req.Username, req.Password).First(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	if result.RowsAffected == 0 {
		panic(errors.New("user not found"))
	}
	var role models.Role
	result = conn.Where("id=?", user.RoleId).First(&role)
	if result.Error != nil {
		panic(result.Error)
	}
	if result.RowsAffected == 0 {
		panic(errors.New("user not found"))
	}
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user.UserName, user.Password)))
	w.Header().Add("Authorization", auth)
	myutils.SendJson(w, map[string]interface{}{"message": "Success", "role": role.Name, "name": user.FullName}, 200)
}

func Register(w http.ResponseWriter, r *http.Request) {
	defer errorhandlers.NormalErrorRest(w, r)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var req dtos.RegRequest
	if err = json.Unmarshal(body, &req); err != nil {
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
	conn := db.DbConnect()
	var user models.User
	var role models.Role
	err = conn.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("email=?", req.Email).First(&user)
		if result.Error == nil {
			return errors.New("email is registered")
		}
		result = tx.Where("user_name=?", req.Username).First(&user)
		if result.Error == nil {
			return errors.New("username is registered")
		}
		result = tx.Where("role_name=?", "User").First(&role)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("role not found")
		}
		user = models.User{UserName: req.Username, Email: req.Email, FullName: req.FullName, Password: req.Password, RoleId: role.ID}
		result = tx.Create(&user)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("error saving data")
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user.UserName, user.Password)))
	w.Header().Add("Authorization", auth)
	myutils.SendJson(w, map[string]interface{}{"message": "Success", "role": role.Name, "name": user.FullName}, 200)
}

func DashboardLogout(w http.ResponseWriter, r *http.Request) {
	defer errorhandlers.NormalErrorRest(w, r)
	myutils.SendJson(w, map[string]interface{}{"message": "Success"}, 200)
}
