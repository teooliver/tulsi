package routes

import (
	"net/http"

	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/teooliver/kanban/internal/bootstrap"
	"github.com/teooliver/kanban/pkg/postgresutils"
)

func Router(deps *bootstrap.AllDeps) http.Handler {
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
		r.With(
			httpin.NewInput(postgresutils.PageRequest{}),
		).Get("/", deps.Handlers.TaskHandler.ListTasks)
		r.Post("/", deps.Handlers.TaskHandler.CreateTask)
		r.Delete("/{id}", deps.Handlers.TaskHandler.DeleteTask)
		r.Put("/{id}", deps.Handlers.TaskHandler.UpdateTask)
	})

	r.Route("/status", func(r chi.Router) {
		r.With(
			httpin.NewInput(postgresutils.PageRequest{}),
		).Get("/", deps.Handlers.StatusHandler.ListAllStatus)
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
