package main

import (
	"context"
	"fmt"
	"wails-selfupdate-spike/models"
	"wails-selfupdate-spike/update"

	"wails-selfupdate-spike/db"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) CheckForUpdate() string {
	updateExists, version := update.CheckForUpdate()
	fmt.Println("updateExists:", updateExists)
	fmt.Println("version:", version)
	if updateExists {
		return version
	}
	return ""
}

func (a *App) DoSelfUpdate() {
	update.UpdateSelf(context.Background())
}

func (a *App) GetCurrentVersion() string {
	return update.CurrentVersion
}

func (a *App) TestDb() models.Connection {
	db := db.InitDb()
	fmt.Println("inserting into db...")
	conn := models.Connection{
		Name:           "Test Connection",
		Protocol:       "mqtt",
		Host:           "test.mosquitto.org",
		Port:           1883,
		IsProtoEnabled: true,
		IsCertsEnabled: false,
	}
	db.Create(&conn)
	fmt.Println("inserted conn with id of:", conn.ID)

	var connection models.Connection
	db.Model(&connection).Preload("Subscriptions").First(&connection)
	fmt.Println("connection:", connection)
	return connection
}
