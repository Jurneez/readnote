package mysql

import (
	"fmt"
	"io/ioutil"
	"tool/readnote/log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	Db *sqlx.DB
}

func NewDB() *DB {
	return &DB{
		Db: &sqlx.DB{},
	}
}

// https://www.topgoer.com/%E6%95%B0%E6%8D%AE%E5%BA%93%E6%93%8D%E4%BD%9C/go%E6%93%8D%E4%BD%9Cmysql/select%E6%93%8D%E4%BD%9C.html
func (db *DB) Connect(host, port, username, password, dbname string) {
	//database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
	if len(password) > 0 {
		username = fmt.Sprintf("%s:%s", username, password)
	}
	if len(port) > 0 {
		host = fmt.Sprintf("%s:%s", host, port)
	}
	sql := fmt.Sprintf("%s@tcp(%s)/%s", username, host, dbname)
	log.Write("log.txt", fmt.Sprintf("[%s]%s", "ConnectMysql", sql))
	database, err := sqlx.Open("mysql", sql)
	if err != nil {
		panic(err)
	}

	db.Db = database
}

func (db *DB) Exec(path string) {
	sqls, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	log.Write("log.txt", fmt.Sprintf("[%s] %s", "Exec", string(sqls)))

	sql := string(sqls)

	if _, err = db.Db.Exec(sql); err != nil {
		panic(err)
	}
}
