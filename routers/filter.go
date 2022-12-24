package routers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/casbin/casbin-website-standalone/util"
)

func TransparentStatic(ctx *context.Context) {
	urlPath := ctx.Request.URL.Path
	if strings.HasPrefix(urlPath, "/api/") {
		return
	}

	repo := beego.AppConfig.String("repo")
	path := fmt.Sprintf("../%s", repo)
	if urlPath == "/" {
		http.Redirect(ctx.ResponseWriter, ctx.Request, "/zh", 301)
		return
	} else {
		path += urlPath
	}

	tokens := strings.Split(path, "/")
	if len(tokens) > 0 {
		lastToken := tokens[len(tokens)-1]
		if lastToken == "" {
			path = filepath.Join(path, "index.html")
		} else if !strings.Contains(lastToken, ".") {
			path += ".html"
		}
	}

	if util.FileExist(path) {
		http.ServeFile(ctx.ResponseWriter, ctx.Request, path)
	} else {
		http.ServeFile(ctx.ResponseWriter, ctx.Request, fmt.Sprintf("../%s/zh/index.html", repo))
	}
}
