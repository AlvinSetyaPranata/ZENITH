package main

import (
	"github.com/AlvinSetyaPranata/ZENITH/backend/config"
)

func main() {
	app := config.CreateFiber()
	// db := config.NewDatabase()

	err := app.Listen(":8000")

	if err != nil {
		config.Log.Critical("Failed to initialize server: ", err.Error())
	}

}
