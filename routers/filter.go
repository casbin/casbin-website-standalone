package routers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/casbin/casbin-website-standalone/util"
)

func TransparentStatic(ctx *context.Context) {
	repo := beego.AppConfig.String("repo")

	urlPath := ctx.Request.URL.Path
	if strings.HasPrefix(urlPath, "/api/") {
		return
	}

	path := fmt.Sprintf("../%s", repo)
	if urlPath == "/" {
		path += "/zh-CN/index.html"
	} else {
		path += urlPath
	}

	path = strings.ReplaceAll(path, "/en/", "/zh-CN/")
	//println(path)

	tokens := strings.Split(path, "/")
	if len(tokens) > 0 && !strings.Contains(tokens[len(tokens)-1], ".") {
		path += ".html"
	}

	if util.FileExist(path) {
		http.ServeFile(ctx.ResponseWriter, ctx.Request, path)
	} else {
		http.ServeFile(ctx.ResponseWriter, ctx.Request, fmt.Sprintf("../%s/zh-CN/index.html", repo))
	}
}
