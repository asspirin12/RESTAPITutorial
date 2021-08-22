package main

import (
	"fmt"
	"net/http"

	"github.com/asspirin12/RESTAPITutorial/internal/comment"
	"github.com/asspirin12/RESTAPITutorial/internal/database"
	transportHTTP "github.com/asspirin12/RESTAPITutorial/internal/transport/http"
)

// App struct which contains pointers to database connections
type App struct{}

// Run sets up our application
func (a *App) Run() error {
	fmt.Println("Setting up our API")

	var err error
	db, err := database.NewDataBase()
	if err != nil {
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		return err
	}
	fmt.Println("Failed to set up server")
	return nil
}

func main() {
	fmt.Println("Hello, GoLand")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
