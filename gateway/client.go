package gateway

import (
	"zhanggongzi/pkg"
)

const DEFAULTCONFIG  = "./config/app.properties"

type GTClient struct {
	Conf        *pkg.Config
	HttpRequest *pkg.HttpRequest
}

// start server
func (gtc *GTClient) Run() error {
	// load default config
	gtc.parse(DEFAULTCONFIG)

	//
	server := new(pkg.HttpRequest)
	err := server.RegisterServers()
	if err != nil {
		return err
	}
	return nil 
}

// run with config path
func (gtc *GTClient) RunWithCfg() error {

	return nil
}

// parse  config and set config
func (gtc *GTClient) parse(path string) error {
	cfg := new(pkg.Config)
	cfg.Parse(path)
	pkg.SetCfg(cfg)
	return nil
}


