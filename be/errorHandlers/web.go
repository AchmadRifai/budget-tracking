package errorhandlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

func NormalErrorRest(w http.ResponseWriter, r *http.Request) {
	errorRest(w, r, 500)
}

func NormalError() {
	if err := recover(); err != nil {
		log.Println("Error", err)
		log.Println(string(debug.Stack()))
	}
}

func AuthErrorRest(w http.ResponseWriter, r *http.Request) {
	errorRest(w, r, 401)
}

func BadReqRest(w http.ResponseWriter, r *http.Request) {
	errorRest(w, r, 400)
}

func errorRest(w http.ResponseWriter, r *http.Request, statusCode int) {
	if err := recover(); err != nil {
		log.Println("Error", err)
		w.WriteHeader(statusCode)
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Authorization", r.Header.Get("Authorization"))
		res, _ := json.Marshal(map[string]interface{}{"message": fmt.Sprintf("%s", err)})
		w.Write(res)
		log.Println(string(debug.Stack()))
	}
}
