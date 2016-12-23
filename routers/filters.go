package routers

import (
  "strings"
  "github.com/astaxie/beego/context"
)

var filterLoggedInUser = func(ctx *context.Context) {
  if strings.HasPrefix(ctx.Input.URL(), "/session") {
    return
  }
  _, ok := ctx.Input.Session("current_user").(map[string]string)
  if !ok {
    ctx.Redirect(302, "/session/login")
  }
}

