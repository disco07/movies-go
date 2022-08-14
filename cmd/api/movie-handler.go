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

	err = app.writeJSON(w, http.StatusOK, movie, "movie")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app apps) getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.models.DB.FindAll()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, movies, "movies")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app apps) insertMovie(w http.ResponseWriter, r *http.Request) {

}

func (app apps) updateMovie(w http.ResponseWriter, r *http.Request) {

}

func (app apps) deleteMovie(w http.ResponseWriter, r *http.Request) {

}

func (app apps) searchMovie(w http.ResponseWriter, r *http.Request) {

}
