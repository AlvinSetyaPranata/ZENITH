package main

import (
	"fmt"

	"github.com/AlvinSetyaPranata/ZENITH/backend/internal/config"
	"github.com/spf13/viper"
)

const DEFAULT_PORT = "8000"

func main() {
	config.LogInit(false)
	defer config.Log.Sync()

	config.Log.Debug("Loading .env")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		config.Log.Fatal("Failed to load .env")
		return
	}

	app := config.CreateFiber(config.Log)
	db := config.NewDatabase(config.Log)

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

	// config.Log.Debug("Configuring middlewares")

	config.Log.Debug("All configurations succesfully configured, yay!")
	config.Log.Debugf("Server is on air, listening to http://127.0.0.1:%s", DEFAULT_PORT)

	err := app.Listen(fmt.Sprintf(":%s", DEFAULT_PORT))

	if err != nil {
		config.Log.Panic("Error during server runtime: ", err.Error())
	}

}
