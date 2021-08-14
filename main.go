package main

import "fmt"

// App struct which contains pointers to database connections
type App struct{}

func (a *App) Run() error {
	fmt.Println("Setting up our API")
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
