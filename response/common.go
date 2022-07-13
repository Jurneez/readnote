package response

import (
	"encoding/json"
	"net/http"
)

type Common struct {
	Code    int32
	Message string
}

// 将数据json化，返回json格式的数据
func SetResponseJsonWrite(w http.ResponseWriter, result interface{}) {
	res, err := json.Marshal(result)
	if err != nil {
		return
	}
	w.Write(res)
}
