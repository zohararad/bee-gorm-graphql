package routers

import (
  "strings"
  "github.com/astaxie/beego/context"
)

var filterLoggedInUser = func(ctx *context.Context) {
  url := ctx.Input.URL()
  if strings.HasPrefix(url, "/session") || strings.HasPrefix(url, "/graphql"){
    return
  }
  _, ok := ctx.Input.Session("current_user").(map[string]string)
  if !ok {
    ctx.Redirect(302, "/session/login")
  }
}

