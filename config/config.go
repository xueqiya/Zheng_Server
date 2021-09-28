package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type config struct {
	Server server
}

type server struct {
	Mode         string
	Addr         string
	ReadTimeout  int
	WriteTimeout int
}

var Cfg config

// init 初始化 Cfg 全局变量。
func init() {
	_, err := toml.DecodeFile("./config.toml", &Cfg)
	if err != nil {
		log.Fatal(err)
	}
}
