package upload

import (
	"net/http"
)

// Result 上传结果基础数据
type Result struct {
	Id int64 `json:"id"`
}

// Upload 上传接口
type Upload interface {
	// UploadHandler 标准上传
	UploadHandler(w http.ResponseWriter, r *http.Request, name string) (err error, res *Result)
	// Detail 获取上传的详细信息
	Detail(res *Result) map[string]interface{}
}
