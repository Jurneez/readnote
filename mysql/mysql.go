package mysql

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"
	"tool/readnote/log"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

func New(host, port, username, password, dbname string) *DB {
	if len(password) > 0 {
		username = fmt.Sprintf("%s:%s", username, password)
	}
	if len(port) > 0 {
		host = fmt.Sprintf("%s:%s", host, port)
	}
	sqlx := fmt.Sprintf("%s@tcp(%s)/%s", username, host, dbname)
	log.Write("log.txt", fmt.Sprintf("[%s]  %s", "ConnectMysql", sqlx))
	database, err := sql.Open("mysql", sqlx)
	if err != nil {
		panic(err)
	}

	e := database.Ping()
	if e != nil {
		panic(e)
	} else {
		fmt.Println("Ping to database successful, connection is still alive")
	}
	return &DB{
		database,
	}
}

func (db *DB) Execx(path string) {
	sqls, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	log.Write("log.txt", fmt.Sprintf("[%s] %s", "Exec", string(sqls)))

	if _, err = db.Exec(string(sqls)); err != nil {
		panic(err)
	}
}

// 插入单条数据
func (db *DB) Insert(tableName string, data map[string]interface{}) (int64, error) {
	if len(data) == 0 {
		return -1, nil
	}
	fields := make([]string, 0)
	args := make([]interface{}, 0)
	ds := make([]string, 0)
	for k, v := range data {
		fields = append(fields, k)
		ds = append(ds, "?")
		args = append(args, v)
	}

	sql := fmt.Sprintf("insert into %s(%s) values(%s)", tableName, strings.Join(fields, ","), strings.Join(ds, ","))
	log.Write("log.txt", fmt.Sprintf("[%s] %s", "Insert", sql))

	r, err := db.Exec(sql, args...)
	if err != nil {
		return -1, err
	}
	lastId, err := r.LastInsertId()

	return lastId, err
}

// 补充增删改查接口
// func (db *DB) Select
