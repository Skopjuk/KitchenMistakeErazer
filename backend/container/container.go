package container

import (
	"KitchenMistakeErazer/backend/configs"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Container struct {
	Config  configs.Config
	DB      *sqlx.DB
	Logging *logrus.Logger
}

func NewContainer(config configs.Config, logging *logrus.Logger) *Container {
	db, err := NewPostgresDB(config)

	if err != nil {
		logrus.Fatalf("cannot connect to db: %s", err.Error())
	}

	return &Container{Config: config,
		DB:      db,
		Logging: logging,
	}
}
