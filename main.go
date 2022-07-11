package main

import (
	"fmt"
	"tool/readnote/common"
	"tool/readnote/config"
	"tool/readnote/mysql"

	"github.com/BurntSushi/toml"
	// "github.com/Jurneez/goutils/slice"
)

func init() {
	// config.InitConf()
	var err error
	var conf config.Config
	_, err = toml.DecodeFile("./config/config.toml", &conf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%+v \n", conf)

	// 链接数据库
	common.Db, err = mysql.ConnectMysql()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func main() {
	fmt.Println("x")
	// s := []string{"a", "b", "x", "a"}
	// sb := slice.Deduplicate(s).([]string)
	// fmt.Println(sb)
}
