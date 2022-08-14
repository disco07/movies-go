package main

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (app apps) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	movie, err := app.models.DB.Find(id)
	if err != nil {
		app.logger.Fatal(err)
	}

	app.writeJSON(w, http.StatusOK, movie, "movie")
}

func (app apps) getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.models.DB.FindAll()
	if err != nil {
		app.logger.Fatal(err)
		app.errorJSON(w, err)
		return
	}
	app.writeJSON(w, http.StatusOK, movies, "movies")
}
