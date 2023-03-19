package main

import (
	"fmt"
	"jwt_registration_api/internal/composites"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"

	"github.com/sirupsen/logrus"
)

const configPath = "../config/app_conf.yaml"

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

	dbComposite, err := composites.NewPostgresComposite(viper.GetString("db.url"))
	if err != nil {
		logger.Fatal(err)
	}
	defer dbComposite.Db.Close()

	signingJwtKey := viper.GetString("signingJwtKey")
	hoursOfJwtAction := viper.GetInt("hoursOfJwtToken")
	userComposite, err := composites.NewUserComposite(dbComposite, signingJwtKey, hoursOfJwtAction, logger)
	if err != nil {
		logger.Fatal(err)
	}

	router := httprouter.New()
	userComposite.Handler.RegisterRoute(router)

	logger.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
