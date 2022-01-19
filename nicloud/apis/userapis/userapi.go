package userapis

import (
  "github.com/gin-gonic/gin"
  "github.com/go-playground/validator/v10"
  "nicloud/users"
  "nicloud/vmerror"
)

func Login(c *gin.Context) {
  username := c.PostForm("username")
  passwd := c.PostForm("passwd")

  user := users.Vms_users{
    Username: username,
    Passwd: passwd,
  }

  res := make(map[string]interface{})
  validate := validator.New()
  err := validate.Struct(user)
  if err != nil {
    res["err"] = vmerror.Error{Message: "参数错误"}
    c.JSON(400, res)
    return
  }

  t, u,  err := users.Login(username, passwd)
  if err != nil {
    res["err"] = vmerror.Error{Message: "登陆失败"}
    c.JSON(200, res)
    return
  }

  res["token"] = t
  res["username"] = u
  res["err"] = nil
  c.JSON(200, res)
}
