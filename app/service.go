package app

import (
	"goKafka/conf"
	"net/http"
)

func StartWebServer(port string, r *Router) {
	conf.LOG.Sugar().Infof("strating http service at: %s", port)
	err := http.ListenAndServe("0.0.0.0:"+port, r)
	if err != nil {
		conf.LOG.Sugar().Errorf("an error occured starting http listener at port: %s", port)
		conf.LOG.Sugar().Errorf("error: %s", err.Error())
	}
}
