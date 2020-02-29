package api

import (
	"fmt"
	"net/http"

	"github.com/510909033/bgf_mvc"
	"github.com/510909033/go_upload/pkg/model/bo/upload"
)

type DemoController struct {
	*bgf_mvc.Controller
}

func (c *DemoController) TestAction(ctx *bgf_mvc.Context) {

	var m = make(map[string]interface{})
	m["env"] = bgf_mvc.GetConfig().IsDevEnvironment()
	ctx.Success(m)
}

func (c *DemoController) DemoAction(ctx *bgf_mvc.Context) {
	r := ctx.Request
	w := ctx.Writer
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		//		w.Header().Set("Access-Control-Allow-Headers",
		//			"Action, Module") //有使用自定义头 需要这个,Action, Module是例子
	}

	if r.Method == "OPTIONS" {
		ctx.String(http.StatusOK, "123")
		return
	}

	err, val := upload.UploadHandler(ctx.Request, "upload")

	if err != nil {
		fmt.Println(err)
	}

	//bo.Test1()

	//ctx.String(http.StatusOK, "这是api demo")
	ctx.Success(val)
}
