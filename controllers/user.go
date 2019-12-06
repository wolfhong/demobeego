package controllers

import (
	"demobeego/models"
	"demobeego/models/merrors"
	"demobeego/models/mh"
	//"log"
)

type UserController struct {
	BaseController
}

// @Title signup
// @Description 注册
// @Param username query string	true "用户名"
// @Param password query string	true "密码"
// @Param email query string true "邮箱"
// @Success 200 {object} mh.Ret
// @router /signup [post]
func (this *UserController) Signup() {
	username := this.GetString("username")
	password := this.GetString("password")
	email := this.GetString("email")
	var err error
	var user models.User
	err = models.DB.Where("name = ?", username).First(&user).Error
	if err == nil {
		this.Data["json"] = mh.NewRet().Error(merrors.AccountExist)
		this.ServeJSON()
		return
	}
	err = models.DB.Where("email = ?", email).First(&user).Error
	if err == nil {
		this.Data["json"] = mh.NewRet().Error(merrors.EmailExist)
		this.ServeJSON()
		return
	}
	user = models.User{
		Name:     username,
		Email:    email,
		Password: models.EncryptPass(password),
	}
	models.DB.Create(&user)
	this.Data["json"] = mh.NewRet()
	this.ServeJSON()
}

// @Title login
// @Description 登录
// @Param username query string	true "用户名"
// @Param password query string	true "密码"
// @Success 200 {object} mh.Ret
// @router /login [post]
func (this *UserController) Login() {
	username := this.GetString("username")
	password := this.GetString("password")
	encrypt_pass := models.EncryptPass(password) //加密后
	var user models.User
	err := models.DB.Where("name = ? and password = ?", username, encrypt_pass).First(&user).Error
	if err != nil {
		this.Data["json"] = mh.NewRet().Error(err)
		this.ServeJSON()
		return
	}
	token := user.CreateToken()
	this.Data["json"] = mh.NewRet().SetData(map[string]interface{}{"token": token})
	this.ServeJSON()
}

// @Title logout
// @Description 登出
// @Param UserToken header string true ""
// @Success 200 {object} mh.Ret
// @router /logout [get,post]
func (this *UserController) Logout() {
	token := this.Ctx.Input.Header("UserToken")
	if token != "" {
		models.DeleteToken(token)
	}
	this.Data["json"] = mh.NewRet()
	this.ServeJSON()
}

// @Description 修改个人信息
// @Success 200 {object} mh.Ret
// @router /profile/update [post]
func (this *UserController) UpdateProfile() {
}
