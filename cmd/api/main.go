package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

type config struct {
	port int
	env  string
}

type app struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "dev", "Application environment (dev|prod)")
	flag.Parse()

	fmt.Println("App's running")

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		currentStatus := struct {
			Status string `json:"status"`
			Env    string `json:"env"`
		}{
			Status: "Availanble",
			Env:    cfg.env,
		}

		js, err := json.Marshal(currentStatus)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(js)
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), nil); err != nil {
		log.Fatal(err)
	}
}
