package myutils

import (
	"be/db"
	"be/models"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"net/mail"
	"strings"
)

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func SendJson(w http.ResponseWriter, data any, statusCode int) {
	res, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(res)
}

func GetUser(r *http.Request) models.User {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		panic(errors.New("\"Authorization\" is required"))
	}
	parts := strings.Split(auth, " ")
	if len(parts) != 2 || parts[0] != "Basic" {
		panic(errors.New("invalid basic Auth"))
	}
	payload, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		panic(err)
	}
	pairs := strings.Split(string(payload), ":")
	if len(pairs) != 2 {
		panic(errors.New("user not found"))
	}
	conn := db.DbConnect()
	var user models.User
	query := "email=? AND password=?"
	if !IsValidEmail(pairs[0]) {
		query = "user_name=? AND password=?"
	}
	result := conn.Where(query, pairs[0], pairs[1]).First(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	if result.RowsAffected == 0 {
		panic(errors.New("user not found"))
	}
	return user
}
