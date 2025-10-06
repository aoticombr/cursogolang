package main

import (
	"sync"

	"github.com/kardianos/service"
)

var onceApp sync.Once
var instanceApp *App

type App struct {
}

func (app *App) Execute() error {
	svcConfig := &service.Config{
		Name:        "GoDemoService",
		DisplayName: "Go Demo Service",
		Description: "This is a Go service demo.",
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
	return nil
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
