package main

import (
	"dictionary/pkg"
	"dictionary/pkg/delivery"
	"dictionary/pkg/service"
	"dictionary/pkg/storage"
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error init configuration, %s", err.Error())
	}

	dbConn, err := storage.NewConnect(viper.GetString("dbLink"))
	if err != nil {
		logrus.Fatalf("connect db error: %s", err.Error())
	}

	stor := storage.NewStorage(dbConn)
	serv := service.NewService(stor)
	handler := delivery.NewHandler(serv)

	server := pkg.NewServer(viper.GetString("bindPort"), handler.InitRoutes())

	if err := server.Run(); err != nil {
		logrus.Fatalf("error http server running, %s", err.Error())
	}
	logrus.Info("api server starting")

}

func initConfig() error {
	var configPath, configFile string

	flag.StringVar(&configPath, "path", "..\\configs", "Path to config file")
	flag.StringVar(&configFile, "config", "config_template", "Name of config file")
	flag.StringVar(&configPath, "p", "..\\configs", "Path to config file")
	flag.StringVar(&configFile, "c", "config_template", "Name of config file")
	flag.Parse()

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configFile)
	return viper.ReadInConfig()
}
