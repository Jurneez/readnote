package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// https://www.topgoer.com/%E6%95%B0%E6%8D%AE%E5%BA%93%E6%93%8D%E4%BD%9C/go%E6%93%8D%E4%BD%9Cmysql/select%E6%93%8D%E4%BD%9C.html
func ConnectMysql() (*sqlx.DB, error) {
	//database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
	database, err := sqlx.Open("mysql", "root@tcp(127.0.0.1:3306)/test")
	return database, err
}
