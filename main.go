package main

import (
	"encoding/json"
	"net/http"
	"tool/readnote/config"
	"tool/readnote/response"
)

func init() {
	config.LoadConfig("./config/config.toml")

	// mysqlTest := config.Conf.Mysql["test"]
	// testDB := mysql.New(mysqlTest.Host, mysqlTest.Port, mysqlTest.UserName, mysqlTest.Password, "test")
	// testDB.Execx("./sql/file.sql")
}

func login(w http.ResponseWriter, r *http.Request) {
	result := response.Common{
		Code:    1,
		Message: "login",
	}
	// 将数据json化，返回json格式的数据
	res, err := json.Marshal(result)
	if err != nil {
		return
	}
	w.Write(res)
}
func main() {
	http.HandleFunc("/login", login)
	http.ListenAndServe(":9091", nil)
}
