// go_upload
package main

import (
	"github.com/510909033/bgf_mvc"
	"github.com/510909033/bgf_mvc/filter/compatible_php_response_format"
	"github.com/510909033/go_upload/pkg/router"
)

func main() {
	app := bgf_mvc.GetApplication()

	// 过滤器
	app.AddFilter(compatible_php_response_format.FormatJsonFilter)
	//app.AddFilter(api_auth.ApiAuthFilter)

	// path中加入项目前缀
	app.EnableAppNameAsPathPrefix()

	app.EnableStaticFileServer()

	// 初始化路由
	router.InitRouter(app)

	//prometheus监控
	//	prometheus.AddPrometheus()

	app.Start()
}
