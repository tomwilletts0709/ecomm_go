package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"
	"production_api/internal/products"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	logger *log.Logger
	db     *sql.DB
}

//run
//mount

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	//middleware
	r.Use(middleware.RequestID) // rate limiting
	r.Use(middleware.RealIP)    // fior rate limiting and analytics
	r.Use(middleware.Logger)    // logging
	r.Use(middleware.Recoverer) // recovery

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("working"))

	})
	
	productHandler := products.NewHandler(nil)
	r.Get('/products', productHandler.ListProducts)
	

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Println("server has started on port", app.config.addr)

	return srv.ListenAndServe()
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
