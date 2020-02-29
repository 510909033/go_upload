package upload

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/510909033/bgf_log"
)

var logger = bgf_log.GetLogger("upload_bo")

func Save(name string) (id int64, err error) {
	bo := &UploadBO{}
	newBO, err := NewUploadBO(bo, false)

	h := md5.New()
	h.Write([]byte(name + strconv.FormatInt(time.Now().UnixNano(), 10)))
	rename := hex.EncodeToString(h.Sum(nil))
	newBO.Name = rename

	//扩展名
	newBO.Ext = path.Ext(name)

	id, err = newBO.Insert()
	if err != nil {
		err = fmt.Errorf("Inert err, err=%w\n", err)
		logger.Errorf(err.Error())

		return
	}

	return

}

const (
	ParseMultipartForm = 32 << 20 //M , 文件最大使用的内存量
	UPLOAD_DIR         = "/vdb1/uploads"
)

func UploadHandler(r *http.Request, name string) (err error, res *Result) {
	if r.Method != "POST" {
		err = fmt.Errorf("不是POST请求, method=%s, request=%+v\n", r.Method, *r)
		return
	}
	r.ParseMultipartForm(ParseMultipartForm)
	f, h, err := r.FormFile(name)
	if err != nil {
		err = fmt.Errorf("FormFile err, err=%w\n", err)
		return
	}
	defer f.Close()
	filename := h.Filename
	filepath := UPLOAD_DIR + "/" + filename

	t, err := os.Create(filepath)
	if err != nil {
		err = fmt.Errorf("Create err, err=%w\n", err)
		return
	}
	defer t.Close()
	if _, err = io.Copy(t, f); err != nil {
		err = fmt.Errorf("Copy err, err=%w\n", err)
		return
	}

	id, err := Save(filename)
	if err != nil {
		err = fmt.Errorf("Save err, err=%w\n", err)
		return
	}

	res = &Result{}
	res.Id = id
	//改文件名
	err = os.Rename(filepath, UPLOAD_DIR+"/"+strconv.FormatInt(id, 10))
	if err != nil {
		err = fmt.Errorf("Rename err, err=%w\n", err)
		return
	}

	return
}
