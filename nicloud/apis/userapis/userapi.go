package userapis

import (
  "github.com/gin-gonic/gin"
  "nicloud/users"
)

func Login(c *gin.Context) {
  username := c.PostForm("username")
  passwd := c.PostForm("passwd")

  res := make(map[string]interface{})
  t, u,  err := users.Login(username, passwd)

  res["token"] = t
  res["username"] = u
  res["err"] = err
  c.JSON(200, res)
}
