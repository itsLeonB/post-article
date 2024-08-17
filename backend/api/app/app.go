package app

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"post-api/database"
)

type app struct {
	db     *sql.DB
	router http.Handler
}

func Init() *app {
	a := new(app)
	a.ConnectDatabase()
	handlers := SetupHandlers(a.db)
	a.router = SetupRouter(handlers)

	return a
}

func (a *app) Serve() {
	defer a.db.Close()
	server := http.Server{
		Addr:    os.Getenv("APP_PORT"),
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("error server listen and serve: %s", err.Error())
	}
}

func (a *app) ConnectDatabase() {
	db, err := database.ConnectMysql()
	if err != nil {
		log.Fatalf("fail connect db: %s", err)
	}

	a.db = db
}
