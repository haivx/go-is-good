package main

import (
	"config"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"handlers"
	"net/http"
)

func Routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
