package main

import (
	"net/http"
	"tool/readnote/config"
)

func init() {
	config.LoadConfig("./config/config.toml")

	// mysqlTest := config.Conf.Mysql["test"]
	// testDB := mysql.New(mysqlTest.Host, mysqlTest.Port, mysqlTest.UserName, mysqlTest.Password, "test")
	// testDB.Execx("./sql/file.sql")
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login"))
}
func main() {
	http.HandleFunc("/login", login)
	http.ListenAndServe(":9091", nil)
}
