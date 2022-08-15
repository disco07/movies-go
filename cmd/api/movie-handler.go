package main

import (
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"github/disco07/movies-go/models"
	"net/http"
	"strconv"
	"time"
)

type jsResp struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (app apps) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	movie, err := app.models.DB.FindMovie(id)
	if err != nil {
		app.logger.Fatal(err)
	}

	err = app.writeJSON(w, http.StatusOK, movie, "movie")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app apps) getAllMovies(w http.ResponseWriter, _ *http.Request) {
	movies, err := app.models.DB.FindMovieAll()
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

func (app apps) insertOrUpdateMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	defer r.Body.Close()

	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		app.errorJSON(w, err)
		return
	}

	if id == 0 {
		movie.CreatedAt = time.Now()
		movie.UpdateAt = time.Now()

		err = app.models.DB.InsertMovie(movie)
		if err != nil {
			app.errorJSON(w, err)
			return
		}

		ok := jsResp{true, "movie successfully added"}
		err = app.writeJSON(w, http.StatusCreated, ok, "response")
		if err != nil {
			app.errorJSON(w, err)
			return
		}
	} else {
		movie.ID = id
		movie.UpdateAt = time.Now()
		err = app.models.DB.UpdateMovie(movie)
		if err != nil {
			app.errorJSON(w, err)
			return
		}

		ok := jsResp{true, "movie successfully updated"}
		err = app.writeJSON(w, http.StatusCreated, ok, "response")
		if err != nil {
			app.errorJSON(w, err)
			return
		}
	}

}

func (app apps) deleteMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}
	err = app.models.DB.DeleteMovie(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	ok := jsResp{true, "movie successfully updated"}
	err = app.writeJSON(w, http.StatusCreated, ok, "response")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app apps) searchMovie(w http.ResponseWriter, r *http.Request) {

}

func (app apps) getAllGenres(w http.ResponseWriter, _ *http.Request) {
	genres, err := app.models.DB.FindGenresAll()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, genres, "genres")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app apps) getMovieByGenre(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}
	movies, err := app.models.DB.FindMovieAll(id)
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
