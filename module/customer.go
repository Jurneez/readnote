package module

type Customer struct {
	id           uint   `db:"id"`
	nickname     string `db:"nickname"`
	age          uint   `db:"age"`
	bornDate     string `db:"born_date"`
	status       uint   `db:"status"`        // 0-注册 1-注销
	registerTime uint   `db:"register_time"` // 注册时间
	isJu         uint   `db:"is_ju"`         //  是否游客用户，0-注册用户，1-游客用户

}
