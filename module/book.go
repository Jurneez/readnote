package module

type file struct {
	id         uint   `db:"id"`
	order      uint   `db:"order"`
	name       string `db:"name"`
	ftype      string `db:"ftype"`
	content    string `db:"content"`
	author     string `db:"author"`
	remark     string `db:"remark"`
	createTime string `db:"create_time"`
	updateTime string `db:"update_time"`
}
