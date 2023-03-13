package main

import (
	postgres_client "jwt_registration_api/pkg/client/postgres"
	"os"

	"github.com/spf13/viper"

	"github.com/sirupsen/logrus"
)

const configPath = "config/app_conf.yaml"

func main() {
	logger := logrus.New()

	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		logger.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = viper.GetString("defaultPort")
	}

	db, err := postgres_client.NewClient(viper.GetString("db.url"))
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
}
