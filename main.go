package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tool/readnote/common"
	"tool/readnote/config"
	"tool/readnote/mysql"
	"tool/readnote/repository"
	"tool/readnote/response"
)

func init() {
	config.LoadConfig("./config/config.toml")

	mysqlTest := config.Conf.Mysql["test"]
	common.Test_DB = mysql.New(mysqlTest.Host, mysqlTest.Port, mysqlTest.UserName, mysqlTest.Password, "test")
	result, err := repository.SearchFileByNo("12345678")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%+v \n", result)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	// 限制post类型
	if r.Method != "POST" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
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
