package main

import (
	"github.com/sirupsen/logrus"
	"product-service/config"
	"product-service/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Fatal(err)
	}
	app.Run(cfg)
}
