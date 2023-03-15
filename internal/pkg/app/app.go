package app

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/muratovdias/technodom-test-task/internal/app/cache"
	"github.com/muratovdias/technodom-test-task/internal/app/handler"
	"github.com/muratovdias/technodom-test-task/internal/app/models"
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
	links, _ := dataFromFile()
	db := database.InitDB()
	app.cache = cache.NewCache(1000)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go app.cache.FillCache(links, &wg)
	go database.InsertData(db, links, &wg)
	time.Sleep(1 * time.Second)
	wg.Wait()
	log.Println("ready")
	app.repository = repository.NewRepository(db)
	app.service = service.NewService(app.repository, app.cache)
	app.handler = handler.NewHandler(app.service)
	return app
}

func dataFromFile() ([]models.Link, error) {
	file, err := os.ReadFile("links.json")
	if err != nil {
		log.Fatal(err)
	}
	var links []models.Link
	err = json.Unmarshal(file, &links)
	if err != nil {
		log.Fatal(err)
	}
	return links, nil
}

func (app *App) Run() {
	log.Println("Server started on port :8787")
	if err := http.ListenAndServe(":8787", app.handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
