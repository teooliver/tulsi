package main

import (
	"fmt"
	"net/http"

	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

// TODO move to .env
const (
	host     = "localhost"
	port     = 5432
	user     = "db_user"
	password = "12345"
	dbname   = "kanban-go"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	// insert
	// hardcoded
	insertStmt := `insert into "task"("title") values('hello_world')`
	_, e := db.Exec(insertStmt)
	CheckError(e)

	// dynamic
	// insertDynStmt := `insert into "Students"("Name", "Roll") values($1, $2)`
	// _, e = db.Exec(insertDynStmt, "Jane", 2)
	// CheckError(e)

	// CHI
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)
}
