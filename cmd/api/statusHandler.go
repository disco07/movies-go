package main

import (
	"encoding/json"
	"net/http"
)

func (app *apps) statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	currentStatus := struct {
		Status string `json:"status"`
		Env    string `json:"env"`
	}{
		Status: "Availanble",
		Env:    app.config.env,
	}

	js, err := json.Marshal(currentStatus)
	if err != nil {
		app.logger.Fatal(err)
	}
	w.Write(js)
}
