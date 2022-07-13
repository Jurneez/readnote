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

var ExtensionToContentType = map[string]string{
	".html": "text/html; charset=utf-8",
	".css":  "text/css; charset=utf-8",
	".js":   "application/javascript",
	".xml":  "text/xml; charset=utf-8",
	".jpg":  "image/jpeg",
}
