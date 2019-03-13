package pkg

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Config struct {
	App map[string]string
	DataSource map[string]string
	All map[string]string
}


func (cfg *Config) Parse(path string) {
	//  read config file
	cfg.All = make(map[string]string)
	cfg.App = make(map[string]string)
	fi, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		panic(err.Error())
		return
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		tmp := strings.TrimLeft(string(a), " ")

		tmp = strings.TrimRight(tmp, " ")
		if len(tmp) == 0 || strings.Index(tmp, "#") == 0 {
			continue
		}

		o := strings.Split(tmp, "=")
		if len(o) == 2 {
			cfg.All[o[0]] = o[1]
		}
	}

	for k, v := range cfg.All {
		if strings.Index(k, "app.") == 0 {
			tmp := strings.TrimPrefix(k, "app.")
			cfg.App[tmp] = v
		}
	}
}

var _cfg *Config = nil

func SetCfg(c *Config) {
	_cfg = c
}
func GetCfg() *Config {
	return _cfg
}