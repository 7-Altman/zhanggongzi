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
	Log map[string]map[string]string
	DataSource map[string]string
	All map[string]string
}


func (cfg *Config) Parse(path string) {
	//  read config file
	cfg.All = make(map[string]string)
	cfg.App = make(map[string]string)
	cfg.Log = make(map[string]map[string]string)
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
			tmp = strings.Trim(tmp," ")
			cfg.App[tmp] = strings.Trim(v," ")
		}
		if strings.Index(k, "log.") == 0 {
			var lg = strings.Split(k, ".")
			if nil == cfg.Log[lg[1]] {
				cfg.Log[lg[1]] = make(map[string]string)
			}
			cfg.Log[lg[1]][strings.Trim(lg[2]," ")] = strings.Trim(v, " ")
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