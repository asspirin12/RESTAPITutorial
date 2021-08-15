package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/TutorialEdge/go-rest-api-course/utils/transport/http"
	"github.com/asspirin12/RESTAPITutorial/utils/database"
)

// App struct which contains pointers to database connections
type App struct{}

// Run sets up our application
func (a *App) Run() error {
	fmt.Println("Setting up our API")

	var err error
	_, err = database.NewDataBase()
	if err != nil {
		return err
	}

	handler := transportHTTP.NewHandler()
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
