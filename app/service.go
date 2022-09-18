package app

import (
	"goKafka/conf"
	"net/http"
)

func StartWebServer(port string) {
	conf.LOG.Sugar().Infof("strating http service at: %s", port)
	err := http.ListenAndServe("0.0.0.0:"+port, nil)

	if err != nil {
		conf.LOG.Sugar().Errorf("an error occured starting http listener at port: %s", port)
		conf.LOG.Sugar().Errorf("error: %s", err.Error())
	}
}
