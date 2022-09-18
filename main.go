package main

import (
	"github.com/spf13/viper"
	"goKafka/app"
	"goKafka/conf"
)

func main() {

	err := conf.LoadAppConfiguration()
	if err != nil {
		conf.LOG.Sugar().Errorf("error while loading configuration %s", err)
	}

	router := app.NewRouter()
	router.ProcessAllRoutes()
	app.StartWebServer(viper.GetString("dev.port"), router)

}
