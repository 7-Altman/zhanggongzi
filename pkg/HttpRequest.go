package pkg

import (
	"fmt"
	"net/http"
)

type HttpRequest struct {
	HttpClient *http.Client
	Server     *http.Server
}

func (hq *HttpRequest) RegisterServers() error {
	app := GetCfg().App
	err := http.ListenAndServe(app["addr"]+":"+app["port"], nil)
	if err == nil {
		fmt.Println("xxxxxxxxxxxxxx")
		fmt.Println("RegisterServer on %s port %s", app["addr"], app["port"])
	}
	return err
}
