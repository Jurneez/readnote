package config

type Config struct {
	Mysql map[string]*MYSQLCfg `toml:"MYSQL"`
}

// 定义的TypeName不能和toml文件中的大类一样
// 不可以 type MYSQL struct ,
type MYSQLCfg struct {
	Host     string `toml:"Host"`     // ip地址
	Port     string `toml:"Port"`     // 端口号
	UserName string `toml:"UserName"` // 用户名
	Password string `toml:"Password"` // 密码
}
