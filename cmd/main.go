package main

import (
	"fmt"
	"net/http"

	"github.com/rafaeldiazmiles/API-rest-tutorialedge/comment"
	"github.com/rafaeldiazmiles/API-rest-tutorialedge/database"
	transportHTTP "github.com/rafaeldiazmiles/API-rest-tutorialedge/transport"
)

// App - the struct which contains things like pointers
// to database connections
type App struct{}

//Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting up our App")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	database.MigrateDB(db)
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to setup server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
