package mysql

import (
	"testing"
	"tool/readnote/config"
)

func Test_Execx(t *testing.T) {
	config.LoadConfig("../config/config.toml")
	mysqlTest := config.Conf.Mysql["test"]
	testDB := New(mysqlTest.Host, mysqlTest.Port, mysqlTest.UserName, mysqlTest.Password, "test")
	testDB.Execx("../sql/file.sql")
}

func Test_Insert_File(t *testing.T) {
	config.LoadConfig("../config/config.toml")
	mysqlTest := config.Conf.Mysql["test"]
	testDB := New(mysqlTest.Host, mysqlTest.Port, mysqlTest.UserName, mysqlTest.Password, "test")
	// testDB.Execx("../sql/dic_tag.sql")
	data := make(map[string]interface{})
	data["no"] = "12345678"
	id, err := testDB.Insert("file", data)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(id)
	}
}
func Test_Insert_DicTag(t *testing.T) {
	config.LoadConfig("../config/config.toml")
	mysqlTest := config.Conf.Mysql["test"]
	testDB := New(mysqlTest.Host, mysqlTest.Port, mysqlTest.UserName, mysqlTest.Password, "test")
	testDB.Execx("../sql/dic_tag.sql")
	data := make(map[string]interface{})
	data["tag"] = "教育"
	id, err := testDB.Insert("dic_tag", data)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(id)
	}
}
