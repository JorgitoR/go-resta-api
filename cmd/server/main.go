package main

import "fmt"

type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting Up Our APP")
	return nil
}

func main() {

	fmt.Println("Go REST API")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error Starting up our REST API")
		fmt.Println(err)
	}

}
