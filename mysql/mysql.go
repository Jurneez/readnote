package mysql

import (
	"fmt"
	"io/ioutil"
	"tool/readnote/common"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// https://www.topgoer.com/%E6%95%B0%E6%8D%AE%E5%BA%93%E6%93%8D%E4%BD%9C/go%E6%93%8D%E4%BD%9Cmysql/select%E6%93%8D%E4%BD%9C.html
func ConnectMysql(host, port, username, password, dbname string) (*sqlx.DB, error) {
	//database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
	sql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)
	database, err := sqlx.Open("mysql", sql)
	return database, err
}

func Exec(path string) error {
	sqls, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(sqls))
	// if len(sqls) > 0 {
	// 	fmt.Println(string(sqls))
	// }
	sql := string(sqls)
	_, err = common.Db.Exec(sql)
	if err != nil {
		fmt.Println("exec failed, ", err)
		panic(err)
	}
	// id, err := r.LastInsertId()
	// if err != nil {
	// 	fmt.Println("exec failed, ", err)
	// 	return
	// }

	// fmt.Println("insert succ:", id)
	return nil
}
