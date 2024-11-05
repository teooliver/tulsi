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
	server := &http.Server{Addr: "0.0.0.0:3001", Handler: routes.Router(deps)}

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
