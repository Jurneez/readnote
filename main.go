package main

import (
	"fmt"
	"tool/readnote/common"
	"tool/readnote/config"
	"tool/readnote/mysql"
	// "github.com/Jurneez/goutils/slice"
)

func init() {
	config.LoadConfig("./config/config.toml")
	fmt.Printf("%+v \n", config.Conf.Mysql["test"])

	// 链接数据库
	var err error
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
