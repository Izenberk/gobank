package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/Izenberk/gobank/internal/config"
	"github.com/Izenberk/gobank/internal/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	cfg := config.NewConfig()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	connStr := "postgres://postgres:password@localhost:5432/gobank?sslmode=disable"

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewPostgresAccountRepository(db)
	_ = repo

	server := &http.Server{
		Addr: 				cfg.Host + ":" + strconv.Itoa(cfg.Port),
		Handler: 			r,
		ReadTimeout: 	cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout: 	cfg.IdleTimeout,
	}

	log.Printf("Server starting on  %s", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}