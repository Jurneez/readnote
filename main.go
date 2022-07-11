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

	// 读取sql文件中的数据
	// 链接数据库
	var err error
	mysqlTest := config.Conf.Mysql["test"]
	common.Db, err = mysql.ConnectMysql(mysqlTest.Host, mysqlTest.Port, mysqlTest.UserName, mysqlTest.Password, "test")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	mysql.Exec("./sql/file.sql")
}

func main() {
	fmt.Println("x")
	// s := []string{"a", "b", "x", "a"}
	// sb := slice.Deduplicate(s).([]string)
	// fmt.Println(sb)
}
