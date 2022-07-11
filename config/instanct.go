package config

import (
	"fmt"
	"sync"

	"github.com/BurntSushi/toml"
)

var Conf Config

func InitConf() {
	var execute sync.Once
	execute.Do(func() {
		if _, err := toml.DecodeFile("./config.toml", &Conf); err != nil {
			fmt.Println(err.Error())
		}
	})
}
