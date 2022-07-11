package main

import (
	"fmt"
	"tool/readnote/common"
	"tool/readnote/config"
	"tool/readnote/mysql"
	// "github.com/Jurneez/goutils/slice"
)

func init() {
	config.InitConf()

	fmt.Printf("%+v \n", config.Conf)
	var err error
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
