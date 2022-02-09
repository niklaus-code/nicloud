package userapis

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/go-playground/validator/v10"
  "nicloud/users"
  "nicloud/vmerror"
  "strconv"
)

func DelUser(c *gin.Context) {
  res := make(map[string]interface{})
  id, err := strconv.Atoi(c.Query("id"))
  if err != nil {
    res["err"] = err
    c.JSON(400, res)
    return
  }
  err = users.DelUser(id)
  res["err"] = err
  c.JSON(200, res)
}

func GetAllRoles(c *gin.Context) {
  res := make(map[string]interface{})
  r, err := users.GetrAllRoles()
  res["res"] = r
  res["err"] = err
  c.JSON(200, res)
}

func Createuser(c *gin.Context) {
  username := c.PostForm("username")
  passwd := c.PostForm("passwd")
  email := c.PostForm("email")
  mobile := c.PostForm("mobile")
  role := c.PostForm("role")

  res := make(map[string]interface{})
  roleobj, err := users.Getrolebyrolename(role)
  if err != nil {
    res["err"] = err
    c.JSON(200, res)
    return
  }

  r := users.Vms_users{
    Username: username,
    Passwd: passwd,
    Email: email,
    Role: roleobj.Id,
    Mobile: mobile,
  }

  validate := validator.New()
  err = validate.Struct(r)
  if err != nil {
    fmt.Println(err)
    res["err"] = vmerror.Error{Message: "参数错误"}
    c.JSON(400, res)
    return
  }

  err = users.Createuser(username, passwd, email, roleobj.Id, mobile)
  res["err"] = err
  c.JSON(200, res)
}


func GetUser(c *gin.Context) {
  res := make(map[string]interface{})
  users, err := users.GetUsers()
  if err != nil {
    res["err"] = err
    c.JSON(400, res)
    return
  }

  res["res"] = users
  res["err"] = nil
  c.JSON(200, res)
}

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
