package main

import (
	"fmt"

	"github.com/AlvinSetyaPranata/ZENITH/backend/internal/config"
)

const DEFAULT_PORT = "8000"

func main() {
	app := config.CreateFiber()
	db := config.NewDatabase()

	config.LogInit(false)

	defer config.Log.Sync()

	config.Log.Debug("Boostrapping")
	config.Boostrap(&config.BoostrapConfig{
		App: app,
		DB:  db,
		Log: config.Log,
	})

	config.Log.Debug("Successfully intialize server")
	config.Log.Debugf("Server is on air, listening to http://127.0.0.1:%s", DEFAULT_PORT)

	err := app.Listen(fmt.Sprintf(":%s", DEFAULT_PORT))

	if err != nil {
		config.Log.Panic("Error during server runtime: ", err.Error())
	}

}
