package module

type Customer struct {
	Id           uint   `db:"id"`
	Nickname     string `db:"nickname"`      // 昵称
	Age          uint   `db:"age"`           // 年龄
	BornDate     string `db:"born_date"`     // 出生日期
	Status       uint   `db:"status"`        // 0-注册 1-注销
	RegisterTime uint   `db:"register_time"` // 注册时间
	UserName     string `db:"user_name"`     // 用户名
	Password     string `db:"password"`      // 密码
	IsTourist    uint   `db:"is_tourist"`    //  是否游客用户，0-注册用户，1-游客用户

}

type CustomerLoginHistory struct {
	Id         uint   `db:"id"`
	CustomerId uint   `db:"customer_id"` // 用户ID
	StartTime  uint   `db:"start_time"`  // 登陆开始时间
	EndTime    uint   `db:"end_time"`    // 登陆结束时间
	LoginIp    string `db:"login_ip"`    // 在此之间用户所在IP
}
