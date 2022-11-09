package main

import (
	"dictionary/pkg"
	"dictionary/pkg/delivery"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig("../configs", "config_template"); err != nil {
		logrus.Fatalf("error init configuration, %s", err.Error())
	}

	handler := delivery.NewHandler()

	server := pkg.NewServer(viper.GetString("bindPort"), handler.InitRoutes())

	if err := server.Run(); err != nil {
		logrus.Fatalf("error http server running, %s", err.Error())
	}
	logrus.Info("api server starting")

}

func initConfig(path, file string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(file)
	return viper.ReadInConfig()
}
