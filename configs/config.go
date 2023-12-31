package configs

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Port       string
	Host       string
	DBPort     string
	Username   string
	Password   string
	DBName     string
	SSLMode    string
	SigningKey string
}

func NewConfig() (Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Errorf("error while reading from config: %s", err)
	}

	return Config{
		Port:       viper.GetString("port"),
		Host:       viper.GetString("db.host"),
		DBPort:     viper.GetString("db.port"),
		Username:   viper.GetString("db.username"),
		Password:   viper.GetString("db.password"),
		DBName:     viper.GetString("db.dbname"),
		SSLMode:    viper.GetString("db.sslmode"),
		SigningKey: viper.GetString("signingKey"),
	}, err
}
