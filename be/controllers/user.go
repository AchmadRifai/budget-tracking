package controllers

import (
	errorhandlers "be/errorHandlers"
	"net/http"
)

func GetChart(w http.ResponseWriter, r *http.Request) {
	defer errorhandlers.NormalErrorRest(w, r)
}
