package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"tool/readnote/common"
	"tool/readnote/response"
)

func init() {
	// config.LoadConfig("./config/config.toml")

	// mysqlTest := config.Conf.Mysql["test"]
	// common.Test_DB = mysql.New(mysqlTest.Host, mysqlTest.Port, mysqlTest.UserName, mysqlTest.Password, "test")
	// result, err := repository.SearchFileByNo("12345678")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Printf("%+v \n", result)
	// }
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

	ext := filepath.Ext(path)
	if contentType := common.ExtensionToContentType[ext]; contentType != "" {
		w.Header().Set("Content-Type", contentType)
	}
	w.Header().Set("Content-Length", strconv.FormatInt(d.Size(), 10))

	w.Write(data)

}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			response.SetResponseJsonWrite(w, response.Common{
				Code:    -1,
				Message: err.Error(),
			})
		}
		defer file.Close()

		f, err := os.OpenFile("./file_resource/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			response.SetResponseJsonWrite(w, response.Common{
				Code:    -1,
				Message: err.Error(),
			})
		}
		defer f.Close()
		io.Copy(f, file)
		fmt.Println("upload file sucess")

		// into read-page
		// url := fmt.Sprintf("/file_resource/%s", handler.Filename)
		url := "/read"
		http.Redirect(w, r, url, http.StatusFound)
	} else {
		response.SetResponseJsonWrite(w, response.Common{
			Code:    -1,
			Message: "error request method",
		})
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./page/upload.html")
	t.Execute(w, nil)
}

func read(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./page/read.html")
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/login", login)

	http.HandleFunc("/file_resource/", fileHandle)

	http.HandleFunc("/uploadFile", uploadFile)
	http.HandleFunc("/read", read)

	http.HandleFunc("/", upload)

	http.ListenAndServe(":9091", nil)
}
