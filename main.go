package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	_ "iHome/models"
	_ "iHome/routers"
	"net/http"
	"strings"
)

func main() {
	ignoreStaticPath()
	beego.Run()
}

// 忽略静态路径
func ignoreStaticPath() {
	// 透明Static
	beego.InsertFilter("/", beego.BeforeRouter, transparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, transparentStatic)
}

func transparentStatic(ctx *context.Context) {
	path := ctx.Request.URL.Path
	beego.Debug("request url: ", path)

	if strings.Index(path, "api") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html/"+ctx.Request.URL.Path)
}
