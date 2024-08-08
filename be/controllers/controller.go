package controllers

import (
	errorhandlers "be/errorHandlers"
	"encoding/json"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	defer errorhandlers.NormalErrorRest(w, r)
	res, err := json.Marshal(map[string]interface{}{"message": "hello world"})
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(res)
}
