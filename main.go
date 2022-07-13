package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
	// if r.Method != "POST" {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	result := response.Common{
		Code:    1,
		Message: "login",
	}
	response.SetResponseJsonWrite(w, result)
}

func fileHandle(w http.ResponseWriter, r *http.Request) {
	path := "." + r.URL.Path
	fmt.Println("path = ", path)

	f, err := os.Open(path)
	if err != nil {
		result := response.Common{
			Code:    -1,
			Message: err.Error(),
		}
		response.SetResponseJsonWrite(w, result)
	}
	defer f.Close()

	d, err := f.Stat()
	if err != nil {
		result := response.Common{
			Code:    -1,
			Message: err.Error(),
		}
		response.SetResponseJsonWrite(w, result)
	}

	// 判断该路径是否是一个文件夹
	if d.IsDir() {
		result := response.Common{
			Code:    -1,
			Message: "error dir",
		}
		response.SetResponseJsonWrite(w, result)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		result := response.Common{
			Code:    -1,
			Message: err.Error(),
		}
		response.SetResponseJsonWrite(w, result)
	}
	w.Write(data)

}

func uploadFile(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/login", login)

	http.HandleFunc("/fileRes/", fileHandle)

	http.HandleFunc("/upload", uploadFile)

	http.ListenAndServe(":9091", nil)
}
