package main

import (
	"tool/readnote/config"
	"tool/readnote/mysql"
)

func init() {
	config.LoadConfig("./config/config.toml")

	mysqlTest := config.Conf.Mysql["test"]
	testDB := mysql.New(mysqlTest.Host, mysqlTest.Port, mysqlTest.UserName, mysqlTest.Password, "test")
	testDB.Execx("./sql/file.sql")
}

func main() {
}
