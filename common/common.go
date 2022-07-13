package common

import (
	"fmt"
	"time"
	"tool/readnote/mysql"
)

var Test_DB *mysql.DB

// timestemp
func GetFileNo() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}
