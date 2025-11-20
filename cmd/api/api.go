package main

import (
	"log"
	"net/http"
	"time"

	"github.com/adedejiosvaldo/gopher_social/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	store  store.Storage
}

type dbConfig struct {
	address      string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type config struct {
	address  string
	dbConfig dbConfig
	env      string
}

// function to return mux
func (app *application) mount() http.Handler {

	r := chi.NewRouter()

	// Recover from panics
	r.Use(middleware.Recoverer)

	// Log IP address
	r.Use(middleware.RealIP)

	// Log request IP address
	r.Use(middleware.RequestID)

	// Timeout for requests
	r.Use(middleware.Timeout((60 * time.Second)))

	// Log requests
	r.Use(middleware.Logger)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheck)

		r.Route("/posts", func(r chi.Router) {
			r.Post("/", app.createPostHandler)

			r.Route("/{postID}", func(r chi.Router) {
				r.Get("/", app.getPostHandler)
			})
		})
	})

	return r
}

func (app *application) run(mux http.Handler) error {

	srv := &http.Server{
		Addr:         app.config.address,
		Handler:      mux,
		WriteTimeout: time.Second * 30, // Timeout for writing response
		ReadTimeout:  time.Second * 10, // Timeout for reading request
		IdleTimeout:  time.Minute,      //if it takes longer than this, the connection will be closed
	}

	log.Printf("Starting application on %s", app.config.address)

	return srv.ListenAndServe()

}
