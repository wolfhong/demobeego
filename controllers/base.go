package controllers

import (
	"demobeego/models"
	"demobeego/models/mh"
	"errors"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

type AuthController struct {
	BaseController
	User  models.User
	Token string
}

func (this *BaseController) IfErrorThenStop(err error) {
	if err != nil {
		this.Data["json"] = mh.NewRet().Error(err)
		this.ServeJSON()
		this.StopRun()
	}
}

func (this *BaseController) IfNotErrorThenStop(err error, msg string) {
	if err == nil {
		this.Data["json"] = mh.NewRet().Error(errors.New(msg))
		this.ServeJSON()
		this.StopRun()
	}
}

func (this *AuthController) AuthToken() {
	token := this.Ctx.Input.Header("UserToken")
	var user models.User
	var err error
	err = models.Token2User(token, &user)
	this.IfErrorThenStop(err)
	this.Token = token
	this.User = user
}

func (this *AuthController) Prepare() {
	//this.AuthToken()
}
