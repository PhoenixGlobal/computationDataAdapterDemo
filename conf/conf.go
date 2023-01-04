package conf

/**
  Author: xiaoy
  Created: 2022-12-13
*/

import (
	"flag"
	"github.com/BurntSushi/toml"
	"go.uber.org/zap"
	"log"
	"os"
	"path"
	"path/filepath"
)

var (
	confPath string
	// Conf all conf
	Conf = &Config{}
)

// Config is
type Config struct {
	ZapLog	       *Log
	SwagConf       *SwagConf
}


type Log struct {
	LogLevel	string
}

type SwagConf struct {
	Addr	string    `toml:"addr"`
}

func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	filePath := path.Join(dir, "conf", "local.toml")  // local.toml  test.toml prd.toml
	flag.StringVar(&confPath, "Conf", filePath, "config path")
	flag.Parse()
	if err := Init(); err != nil {
		log.Fatal("conf.Init() error(%v)", zap.String("err", err.Error()))
		panic(err)
	}
}


// Init init conf
func Init() error {
	return local()
}

func NewConfig() *Config {
	var err error
	conf := &Config{}
	if confPath != "" {
		_, err = toml.DecodeFile(confPath, conf)
	}
	if err != nil {
		panic("NewConfig panic")
	}
	return conf
}

func local() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

