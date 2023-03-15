package app

import (
	"fmt"
	"net/http"

	"github.com/muratovdias/technodom-test-task/internal/app/cache"
	"github.com/muratovdias/technodom-test-task/internal/app/handler"
	"github.com/muratovdias/technodom-test-task/internal/app/repository"
	"github.com/muratovdias/technodom-test-task/internal/app/service"
	"github.com/muratovdias/technodom-test-task/pkg/database"
)

type App struct {
	service    *service.Service
	cache      *cache.Cache
	repository *repository.Repository
	handler    *handler.Handler
}

func NewApp() *App {
	app := new(App)
	db := database.InitDB()
	app.cache = cache.NewCache(1000)
	database.InsertData(db)
	app.repository = repository.NewRepository(db)
	app.service = service.NewService(app.repository, app.cache)
	app.handler = handler.NewHandler(app.service)
	return app
}

func (app *App) Run() error {
	fmt.Println("Server started on port :8787")
	if err := http.ListenAndServe(":8787", app.handler.InitRoutes()); err != nil {
		return err
	}
	return nil
}
