package config

import (
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func LogInit(isProd bool) {
	var zapLogger *zap.Logger
	var err error

	if isProd {
		zapLogger, err = zap.NewProduction()
	} else {
		zapLogger, err = zap.NewDevelopment()
	}

	if err != nil {
		panic("cannot initialize zap logger: " + err.Error())
	}

	Log = zapLogger.Sugar()
}
