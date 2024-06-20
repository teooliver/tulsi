package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/teooliver/kanban/internal/bootstrap"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	config, err := bootstrap.Config(".env")
	if err != nil {
		log.Fatal("Error loading .env file %w", err)
		panic("error loading .env file")
	}

	deps, err := bootstrap.Deps(ctx, config)
	if err != nil {
		log.Fatal("Error bootstraping application: %w", err)
		panic("error bootstraping application")
	}

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
	})

	r.Route("/status", func(r chi.Router) {
		// TODO: Add Pagination
		r.Get("/", deps.Handlers.StatusHandler.ListAllStatus)
		r.Post("/", deps.Handlers.StatusHandler.CreateStatus)
		r.Delete("/{id}", deps.Handlers.StatusHandler.DeleteStatus)
		r.Put("/{id}", deps.Handlers.StatusHandler.UpdateStatus)
	})

	r.Route("/user", func(r chi.Router) {
		// TODO: Add Pagination
		r.Get("/", deps.Handlers.UserHandler.ListUsers)
		r.Post("/", deps.Handlers.UserHandler.CreateUser)
		r.Delete("/{id}", deps.Handlers.UserHandler.DeleteUser)
		r.Put("/{id}", deps.Handlers.UserHandler.UpdateUser)
	})

	http.ListenAndServe(":3000", r)
	log.Println("Listenning at :3000")
}
