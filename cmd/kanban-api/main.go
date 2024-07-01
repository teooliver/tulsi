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
	"github.com/teooliver/kanban/internal/bootstrap"
)

func main() {
	ctx, initialCancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer initialCancel()
	// defer log.Println("Gracefully shuting down")
	// defer os.Exit(0)

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

	// The HTTP Server
	server := &http.Server{Addr: "0.0.0.0:3000", Handler: router(deps)}
	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	fmt.Println("Listenning at :3000")
	fmt.Println("Waiting for ctrl+c...")

}

func router(deps *bootstrap.AllDeps) http.Handler {
	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

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

	r.Route("/users", func(r chi.Router) {
		// TODO: Add Pagination
		r.Get("/", deps.Handlers.UserHandler.ListUsers)
		r.Post("/", deps.Handlers.UserHandler.CreateUser)
		r.Delete("/{id}", deps.Handlers.UserHandler.DeleteUser)
		r.Put("/{id}", deps.Handlers.UserHandler.UpdateUser)
	})

	return r
}
