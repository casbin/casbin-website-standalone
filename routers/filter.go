// Copyright 2022 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
		if getSessionUser(ctx) == "" {
			setSessionUser(ctx, "redirected")
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/zh", 301)
			return
		}
	}

	path += urlPath
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
