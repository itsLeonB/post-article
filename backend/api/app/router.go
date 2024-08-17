package app

import (
	"database/sql"
	"net/http"
	"os"
	"post-api/handler"
	"post-api/middleware"
	"post-api/repository"
	"post-api/service"

	"github.com/gin-gonic/gin"
)

type handlers struct {
	postHandler *handler.PostHandler
}

func SetupHandlers(db *sql.DB) *handlers {
	repo := repository.NewPostRepository(db)
	svc := service.NewPostService(repo)

	return &handlers{
		postHandler: handler.NewPostHandler(svc),
	}
}

func SetupRouter(handlers *handlers) http.Handler {
	gin.SetMode(os.Getenv("APP_ENV"))
	r := gin.Default()
	r.Use(middleware.Error())

	r.GET("/article", handlers.postHandler.GetAll())
	r.GET("/article/:id", handlers.postHandler.GetByID())

	return r
}
