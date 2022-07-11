package main

import (
	"fmt"
	"tool/readnote/config"
	"tool/readnote/mysql"
	// "github.com/Jurneez/goutils/slice"
)

func init() {
	config.LoadConfig("./config/config.toml")
	fmt.Printf("%+v \n", config.Conf.Mysql["test"])

	// 读取sql文件中的数据
	// 链接数据库
	testDB := mysql.NewDB()
	mysqlTest := config.Conf.Mysql["test"]
	testDB.Connect(mysqlTest.Host, mysqlTest.Port, mysqlTest.UserName, mysqlTest.Password, "test")
	testDB.Exec("./sql/file.sql")
}

func main() {
	fmt.Println("x")
	// s := []string{"a", "b", "x", "a"}
	// sb := slice.Deduplicate(s).([]string)
	// fmt.Println(sb)
}
