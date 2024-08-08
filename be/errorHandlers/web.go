package errorhandlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

func NormalErrorRest(w http.ResponseWriter, r *http.Request) {
	if err := recover(); err != nil {
		log.Println("Error", err)
		w.WriteHeader(500)
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Authorization", r.Header.Get("Authorization"))
		res, _ := json.Marshal(map[string]interface{}{"message": fmt.Sprintf("%s", err)})
		w.Write(res)
		log.Println(string(debug.Stack()))
	}
}

func NormalError() {
	if err := recover(); err != nil {
		log.Println("Error", err)
		log.Println(string(debug.Stack()))
	}
}
