package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/teooliver/kanban/internal/bootstrap"
	"github.com/teooliver/kanban/internal/routes"
)

func main() {
	ctx, initialCancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer initialCancel()
	// defer log.Println("Gracefully shuting down")
	// defer os.Exit(0)

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

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
	server := &http.Server{Addr: "0.0.0.0:3000", Handler: routes.Router(deps)}

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	// Run the server
	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	fmt.Println("Listenning at :3000")
	fmt.Println("Waiting for ctrl+c...")

	// Wait for server context to be stopped
	<-serverCtx.Done()

}

// func router(deps *bootstrap.AllDeps) http.Handler {
// 	r := chi.NewRouter()

// 	// A good base middleware stack
// 	r.Use(middleware.RequestID)
// 	r.Use(middleware.RealIP)
// 	r.Use(middleware.Logger)
// 	r.Use(middleware.Recoverer)

// 	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("welcome"))
// 	})

// 	r.Route("/tasks", func(r chi.Router) {
// 		// TODO: Add Pagination

// 		r.With(
// 			httpin.NewInput(postgresutils.PageRequest{}),
// 		).Get("/", deps.Handlers.TaskHandler.ListTasks)
// 		r.Post("/", deps.Handlers.TaskHandler.CreateTask)
// 		r.Delete("/{id}", deps.Handlers.TaskHandler.DeleteTask)
// 		r.Put("/{id}", deps.Handlers.TaskHandler.UpdateTask)
// 	})

// 	r.Route("/status", func(r chi.Router) {
// 		// TODO: Add Pagination
// 		r.Get("/", deps.Handlers.StatusHandler.ListAllStatus)
// 		r.Post("/", deps.Handlers.StatusHandler.CreateStatus)
// 		r.Delete("/{id}", deps.Handlers.StatusHandler.DeleteStatus)
// 		r.Put("/{id}", deps.Handlers.StatusHandler.UpdateStatus)
// 	})

// 	r.Route("/users", func(r chi.Router) {
// 		// TODO: Add Pagination
// 		r.Get("/", deps.Handlers.UserHandler.ListUsers)
// 		r.Post("/", deps.Handlers.UserHandler.CreateUser)
// 		r.Delete("/{id}", deps.Handlers.UserHandler.DeleteUser)
// 		r.Put("/{id}", deps.Handlers.UserHandler.UpdateUser)
// 	})

// 	return r
// }
