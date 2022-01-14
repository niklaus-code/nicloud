package userapis

import (
  "github.com/gin-gonic/gin"
  "nicloud/users"
  "nicloud/vmerror"
)

func Login(c *gin.Context) {
  username := c.PostForm("username")
  passwd := c.PostForm("passwd")

  res := make(map[string]interface{})
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
