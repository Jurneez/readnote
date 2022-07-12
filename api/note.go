package api

import (
	"tool/readnote/module"
)

func CreateNoteSingle(req *module.Note) error {
	data := make(map[string]interface{})
	if chickOriginC(req.OriginC) {
		data["origin_c"] = req.OriginC
	}
	if chickContent(req.Content) {
		data["content"] = req.Content
	}

	// if req.Content

	return nil
}

// TODO：将笔记内容进行切割
func chickOriginC(originC string) bool {
	if len(originC) == 0 || len(originC) > 1024 {
		return false
	}
	return true
}

// TODO：将笔记内容进行切割
func chickContent(content string) bool {
	if len(content) == 0 || len(content) > 1024 {
		return false
	}
	return true
}
