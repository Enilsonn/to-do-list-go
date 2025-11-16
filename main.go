package main

import (
	"context"
	"log"
	"net/http"
	"to-do-list/internal/handlers"
	"to-do-list/internal/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ctx := context.Background()

	db, err := repository.InitSQLiteDB(ctx, "./tasks.db")
	if err != nil {
		log.Fatalf("error to initialize or access database: %v", err)
	}
	defer db.Close()

	repo := repository.NewSQLiteRepository(db)
	h := handlers.NewHandler(repo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", h.List)
		r.Post("/", h.Create)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.GetByID)
			r.Put("/", h.Update)
			r.Delete("/", h.Delete)
		})
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
