package main

import (
	"fmt"

	"github.com/AlvinSetyaPranata/ZENITH/backend/internal/config"
)

const DEFAULT_PORT = "8000"

func main() {
	config.LogInit(false)
	app := config.CreateFiber(config.Log)
	db := config.NewDatabase(config.Log)

	defer config.Log.Sync()

	config.Log.Debug("Boostrapping routers")
	configuration_status := config.Boostrap(&config.BoostrapConfig{
		App: app,
		DB:  db,
		Log: config.Log,
	})

	if !configuration_status {
		config.Log.Panic("Server failed to initialize")
		return
	}

	config.Log.Debug("All configurations succesfully configured, yay!")
	config.Log.Debugf("Server is on air, listening to http://127.0.0.1:%s", DEFAULT_PORT)

	err := app.Listen(fmt.Sprintf(":%s", DEFAULT_PORT))

	if err != nil {
		config.Log.Panic("Error during server runtime: ", err.Error())
	}

}
