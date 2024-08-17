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
	trx := repository.NewTransactor(db)
	repo := repository.NewPostRepository(db)
	svc := service.NewPostService(trx, repo)

	return &handlers{
		postHandler: handler.NewPostHandler(svc),
	}
}

func SetupRouter(handlers *handlers) http.Handler {
	gin.SetMode(os.Getenv("APP_ENV"))
	r := gin.Default()
	r.Use(middleware.Error(), middleware.CORS())

	r.GET("/article", handlers.postHandler.GetAll())
	r.GET("/article/:id", handlers.postHandler.GetByID())
	r.POST("/article", handlers.postHandler.Insert())
	r.PUT("/article/:id", handlers.postHandler.Update())
	r.DELETE("/article/:id", handlers.postHandler.Delete())

	return r
}
