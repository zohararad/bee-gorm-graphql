package controllers

import (
  "net/http"
  "github.com/astaxie/beego"
  "encoding/json"
  "context"
  "github.com/neelance/graphql-go"
  gql "github.com/zohararad/bee-gorm-graphql/graphql"
)

var schema *graphql.Schema

type GraphqlController struct {
  beego.Controller
}

func init()  {
  var err error
  schema, err = graphql.ParseSchema(gql.Schema, &gql.Resolver{})
  if err != nil {
    panic(err)
  }
}

func (c *GraphqlController) Prepare()  {
  c.EnableXSRF = false
}

func (c *GraphqlController) Index() {
  c.TplName = "graphql/index.html"
}

func (c *GraphqlController) Query() {
  var params struct {
    Query         string                 `json:"query"`
    OperationName string                 `json:"operationName"`
    Variables     map[string]interface{} `json:"variables"`
  }
  if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&params); err != nil {
    c.CustomAbort(http.StatusInternalServerError, err.Error())
  } else {
    currentUser := c.GetSession("current_user")
    rCtx := c.Ctx.Request.Context()
    gqlCtx := context.WithValue(rCtx, "current_user", currentUser)
    response := schema.Exec(gqlCtx, params.Query, params.OperationName, params.Variables)
    responseJSON, err := json.Marshal(response)
    if err != nil {
      c.CustomAbort(http.StatusInternalServerError, err.Error())
      return
    }
    c.Data["json"] = string(responseJSON)
    c.ServeJSON()
  }
}

