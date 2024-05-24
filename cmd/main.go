package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"github.com/teooliver/kanban/internal/controller/task"
)

func main() {
	// config, err := bootstrap.Config(".env")
	// if err != nil {
	// 	// TODO: Better error handling
	// 	log.Fatal("Error loading .env file")
	// }
	//

	// insert
	// hardcoded
	// insertStmt := `insert into "task"("title") values('hello_world2')`
	// _, e := db.Exec(insertStmt)
	// CheckError(e)
	// models.ListTasks(db)

	// dynamic
	// insertDynStmt := `insert into "Students"("Name", "Roll") values($1, $2)`
	// _, e = db.Exec(insertDynStmt, "Jane", 2)
	// CheckError(e)
	//
	taskHandler := task.New(taskService)

	// CHI
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Route("/task", func(r chi.Router) {
		// r.With(paginate).Get("/", listArticles)                           // GET /articles
		// r.With(paginate).Get("/{month}-{day}-{year}", listArticlesByDate) // GET /articles/01-16-2017
		r.Get("/", taskHandler.ListTasks)
		r.Post("/", task.CreateTask) // POST /task

	})

	http.ListenAndServe(":3000", r)
}
