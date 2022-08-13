package main

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"github/disco07/movies-go/models"
	"net/http"
	"strconv"
	"time"
)

func (app apps) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}
	movie := models.Movie{
		ID:          id,
		Title:       "Some movie",
		Description: "Some description",
		Year:        2022,
		ReleaseDate: time.Date(2022, 01, 01, 22, 25, 00, 00, time.Local),
		Runtime:     100,
		Rating:      5,
		MPAARating:  "PG-13",
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}
	app.writeJSON(w, http.StatusOK, movie, "movie")
}

func (app apps) getAllMovies(w http.ResponseWriter, r *http.Request) {

}
