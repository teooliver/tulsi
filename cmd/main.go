package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"github.com/teooliver/kanban/internal/bootstrap"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()
	defer fmt.Println("Bye")

	config, err := bootstrap.Config(".env")
	if err != nil {
		// TODO: Better error handling
		log.Fatal("Error loading .env file")
	}

	deps, err := bootstrap.Deps(ctx, config)

	// CHI
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Route("/tasks", func(r chi.Router) {
		// TODO: Add Pagination
		r.Get("/", deps.Handlers.TaskHandler.ListTasks)
		r.Post("/", deps.Handlers.TaskHandler.CreateTask)
		r.Delete("/{id}", deps.Handlers.TaskHandler.DeleteTask)
		r.Put("/{id}", deps.Handlers.TaskHandler.UpdateTask)
		r.Post("/seed", deps.Handlers.TaskHandler.SeedTasks)

	})

	http.ListenAndServe(":3000", r)
	println("Listenning at :3000")
}
