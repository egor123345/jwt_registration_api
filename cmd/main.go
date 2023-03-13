package main

import (
	"fmt"
	postgres_client "jwt_registration_api/pkg/client/postgres"

	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Print("start")

	logger := logrus.New()
	db, err := postgres_client.NewClient("Из гуся")
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
}
