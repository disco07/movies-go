package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *apps) initializeRoutes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movie/:id", app.getOneMovie)
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.getAllMovies)
	return app.enableCORS(router)
}
