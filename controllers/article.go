package controllers

import (
	"demobeego/models"
	"demobeego/models/mh"
	"encoding/json"
)

type ArticleController struct {
	AuthController
}

// @Description 创建文章
// @Param UserToken header string true ""
// @Param - body models.ArticleForm true ""
// @Success 200 {object} mh.Ret
// @router /create [post]
func (this *ArticleController) Create() {
	this.AuthToken()
	var err error
	var article models.Article
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &article)
	this.IfErrorThenStop(err)

	user := this.User
	result, err := models.CreateArticle(&article, &user)
	this.IfErrorThenStop(err)

	this.Data["json"] = mh.NewRet().SetData(result)
	this.ServeJSON()
}

// @Description 修改文章
// @Param UserToken header string true ""
// @Param - body models.ArticleForm true ""
// @Success 200 {object} mh.Ret
// @router /update [post]
func (this *ArticleController) Update() {
	this.AuthToken()
	var err error
	var article models.Article
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &article)
	this.IfErrorThenStop(err)

	user := this.User
	result, err := models.UpdateArticle(&article, &user, false)
	this.IfErrorThenStop(err)

	this.Data["json"] = mh.NewRet().SetData(result)
	this.ServeJSON()
}

// @Description 删除文章
// @Param UserToken header string true ""
// @Param id query integer true "文章ID"
// @Success 200 {object} mh.Ret
// @router /delete [post]
func (this *ArticleController) Delete() {
	this.AuthToken()
	var err error
	user := this.User
	id, err := this.GetInt64("id")
	this.IfErrorThenStop(err)

	err = models.DeleteArticle(id, &user, false)
	this.IfErrorThenStop(err)

	this.Data["json"] = mh.NewRet()
	this.ServeJSON()
}

// @Description 查看文章
// @Param id query integer true "文章ID"
// @Success 200 {object} mh.Ret
// @router /get [get]
func (this *ArticleController) Get() {
	var err error
	id, err := this.GetInt64("id")
	this.IfErrorThenStop(err)

	result, err := models.GetArticle(id, &models.User{}, true)
	this.IfErrorThenStop(err)

	this.Data["json"] = mh.NewRet().SetData(result)
	this.ServeJSON()
}
