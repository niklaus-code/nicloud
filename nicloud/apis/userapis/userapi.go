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

func DelUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		vmerror.REQUESTERROR(c, err)
		return
	}

	checkuser := vm2.Checkuser(id)
	if checkuser != nil {
		vmerror.SERVERERROR(c, checkuser)
		return
	}

	err = users.DelUser(id)
	if err != nil {
		vmerror.SERVERERROR(c, err)
		return
	}
	vmerror.SUCCESS(c, nil)
}

func GetAllRoles(c *gin.Context) {
	r, err := users.GetrAllRoles()
	if err != nil {
		vmerror.SERVERERROR(c, err)
		return
	}
	vmerror.SUCCESS(c, r)
}

func Createuser(c *gin.Context) {
	username := c.PostForm("username")
	passwd := c.PostForm("passwd")
	email := c.PostForm("email")
	mobile := c.PostForm("mobile")
	role := c.PostForm("role")

	roleobj, err := users.Getrolebyrolename(role)
	if err != nil {
		vmerror.SERVERERROR(c, err)
		return
	}

	r := users.Vms_users{
		Username: username,
		Passwd:   passwd,
		Email:    email,
		Role:     roleobj.Id,
		Mobile:   mobile,
	}

	validate := validator.New()
	err = validate.Struct(r)
	if err != nil {
		vmerror.REQUESTERROR(c, err)
		return
	}

	encryption := utils.Encryption(passwd)

	err = users.Createuser(username, encryption, email, roleobj.Id, mobile)
	if err != nil {
		vmerror.SERVERERROR(c, err)
		return
	}
	vmerror.SUCCESS(c, nil)
}

func GetUser(c *gin.Context) {
	users, err := users.GetUsers()
	if err != nil {
		vmerror.SERVERERROR(c, err)
		return
	}

	vmerror.SUCCESS(c, users)
}

func Changepasswd(c *gin.Context) {
	username := c.PostForm("username")
	oldpasswd := c.PostForm("oldpasswd")
	newpasswd1 := c.PostForm("newpasswd1")
	newpasswd2 := c.PostForm("newpasswd2")

	if newpasswd1 != newpasswd2 {
		vmerror.SUCCESS(c, vmerror.Error{Message: "2次输入的新密码不一致"})
		return
	}

	if len(newpasswd1) < 5 {
		vmerror.SUCCESS(c, vmerror.Error{Message: "密码长度不能小于5"})
		return
	}

	err := users.ChangePasswd(username, utils.Encryption(oldpasswd), utils.Encryption(newpasswd1))
	if err != nil {
		vmerror.SERVERERROR(c, vmerror.Error{Message: err.Error()})
		return
	}

	vmerror.SUCCESS(c, nil)
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

	t, u, err := users.CheckPWD(username, encryption)
	if err != nil {
		vmerror.SERVERERROR(c, vmerror.Error{Message: "登陆失败"})
		return
	}

	res["token"] = t
	res["username"] = u
	res["err"] = nil
	c.JSON(200, res)
}
