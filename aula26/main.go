package main

import (
	"sync"

	"github.com/kardianos/service"
)

var onceApp sync.Once
var instanceApp *App

type App struct {
	CoreApi *CoreApi
}

func (app *App) Execute() error {
	svcConfig := &service.Config{
		Name:        "ApiSrv",
		DisplayName: "ApiSrv Service",
		Description: "Minha api",
	}
	service, err := service.New(app, svcConfig)
	if err != nil {
		return err
	}

	return service.Run()
}

func (app *App) Start(s service.Service) error {
	go app.Run()
	return nil
}

func (app *App) Stop(s service.Service) error {
	return nil //nao fazer nada
}
func (app *App) Run() error {
	CoreApi := NewCoreApi()
	go CoreApi.Start()
	return nil
}

func NewApp() *App {
	onceApp.Do(func() {
		instanceApp = &App{}
	})
	return instanceApp

}

func main() {
	NewApp().Execute()
}
