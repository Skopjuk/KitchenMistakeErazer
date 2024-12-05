package main

import (
	"KitchenMistakeErazer/backend/container"
	"KitchenMistakeErazer/backend/server"
	"KitchenMistakeErazer/configs"
	"context"
	firebase "firebase.google.com/go/v4"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

func main() {
	logging := logrus.New()
	logging.SetReportCaller(true)

	config, err := configs.NewConfig()

	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	containerInstance := container.NewContainer(config, logging)
	opt := option.WithCredentialsFile("./backend/credentials.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		logging.Fatalf("Error initializing app: %v\n", err)
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		logrus.Fatalf("Error creating client: %v\n", err)
	}

	logging.Info("http server starting")
	err = server.Run(config.Port, *containerInstance, *client)
	if err != nil {
		logging.Fatalf("error occured while running http server: %s, address: %s", err.Error(), config.Port)
	}
}
