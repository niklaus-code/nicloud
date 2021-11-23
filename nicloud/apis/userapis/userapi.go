package userapis

import (
  "github.com/gin-gonic/gin"
  "nicloud/users"
)

func Login(c *gin.Context) {
  username := c.PostForm("username")
  passwd := c.PostForm("passwd")

  res := make(map[string]interface{})
  r, err := users.Login(username, passwd)

  res["res"] = r
  res["err"] = err
  c.JSON(200, res)
}
