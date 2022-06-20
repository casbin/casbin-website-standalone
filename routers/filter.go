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
		http.Redirect(ctx.ResponseWriter, ctx.Request, "/zh", 301)
		return
	} else {
		path += urlPath
	}

	tokens := strings.Split(path, "/")
	if len(tokens) > 0 && !strings.Contains(tokens[len(tokens)-1], ".") {
		path += "/index.html"
	}

	if util.FileExist(path) {
		http.ServeFile(ctx.ResponseWriter, ctx.Request, path)
	} else {
		http.ServeFile(ctx.ResponseWriter, ctx.Request, fmt.Sprintf("../%s/zh/index.html", repo))
	}
}
