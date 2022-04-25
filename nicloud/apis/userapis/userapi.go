package userapis

import (
  "github.com/gin-gonic/gin"
  "github.com/go-playground/validator/v10"
  "nicloud/users"
  "nicloud/utils"
  vm2 "nicloud/vm"
  "nicloud/vmerror"
  "strconv"
)

type User struct {
}

func DelUser(c *gin.Context) {
  res := make(map[string]interface{})
  id, err := strconv.Atoi(c.Query("id"))

  if err != nil {
    res["err"] = err
    c.JSON(400, res)
    return
  }

  checkuser := vm2.Checkuser(id)
  if checkuser != nil {
    res["err"] = checkuser
    c.JSON(200, res)
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
    res["err"] = vmerror.Error{Message: "参数错误"}
    c.JSON(400, res)
    return
  }

  encryption := utils.Encryption(passwd)

  err = users.Createuser(username, encryption, email, roleobj.Id, mobile)
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

func Changepasswd(c *gin.Context) {
  res := make(map[string]interface{})
  username := c.PostForm("username")
  oldpasswd := c.PostForm("oldpasswd")
  newpasswd1 := c.PostForm("newpasswd1")
  newpasswd2 := c.PostForm("newpasswd2")

  if newpasswd1 != newpasswd2 {
    res["err"] = vmerror.Error{Message: "2次输入的新密码不一致"}
    c.JSON(200, res)
    return
  }

  if len(newpasswd1) < 5 {
    res["err"] = vmerror.Error{Message: "密码长度不能小于5"}
    c.JSON(200, res)
    return
  }

  err := users.ChangePasswd(username, utils.Encryption(oldpasswd), utils.Encryption(newpasswd1))
  if err != nil {
    res["err"] = vmerror.Error{Message: err.Error()}
    c.JSON(200, res)
    return
  }

  res["err"] = nil
  c.JSON(200, res)
}

// @Summary 用户登录接口
// @Accept application/json
// @Produce application/json
// @Success 200 object  users.Vms_users
// @Router /api/user/login [post]
// @Param usernmae path string true "name"
// @Param passwd path string true "name"
func Login(c *gin.Context) {
  username := c.PostForm("username")
  passwd := c.PostForm("passwd")
  res := make(map[string]interface{})

  encryption := utils.Encryption(passwd)

  t, u,  err := users.CheckPWD(username, encryption)
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
