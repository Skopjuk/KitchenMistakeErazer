package main

import (
	"KitchenMistakeErazer/configs"
	"KitchenMistakeErazer/container"
	"KitchenMistakeErazer/server"
	"github.com/sirupsen/logrus"
)

func main() {
	logging := logrus.New()
	logging.SetReportCaller(true)

	config, err := configs.NewConfig()

	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	containerInstance := container.NewContainer(config, logging)

	logrus.Info("http server starting")
	err = server.Run(config.Port, *containerInstance)
	if err != nil {
		logrus.Fatalf("error occured while running http server: %s, address: %s", err.Error(), config.Port)
	}
}
