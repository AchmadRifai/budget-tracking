package myutils

import (
	"encoding/json"
	"net/http"
	"net/mail"
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
