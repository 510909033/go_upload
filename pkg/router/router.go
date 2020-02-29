package router

import (
	"github.com/510909033/bgf_log"
	"github.com/510909033/bgf_mvc"
	"github.com/510909033/go_upload/pkg/controller/api"
)

var log = bgf_log.GetLogger("filter.apiAuth")

func InitRouter(app *bgf_mvc.Application) {

	// api demo
	app.RegisterDynamicRouter(&api.DemoController{})

}
