package models

import (
	"demobeego/models/merrors"
)

type Article struct {
	IdPart
	TimePart
	UserId      int64  `json:"user_id"`
	Status      uint   `json:"status" gorm:"default:0"`
	Title       string `json:"title" gorm:"size:255"`
	Content     string `json:"content" gorm:"type:longtext"`
	Attachments string `json:"-" gorm:"type:text"`
	Stype       uint   `json:"stype" gorm:"default:0"`
	Dtype       uint   `json:"dtype" gorm:"default:0"`
}

type Comment struct {
	IdPart
	TimePart
	UserId    int64  `json:"user_id"`
	ArticleId int64  `json:"article_id"`
	Content   string `json:"content" gorm:"size:255"`
}

type ArticleForm struct {
	Title       string `json:"title" gorm:"size:255"`
	Content     string `json:"content" gorm:"type:longtext"`
	Attachments string `json:"-" gorm:"type:text"`
	Stype       uint   `json:"stype" gorm:"default:0"`
	Dtype       uint   `json:"dtype" gorm:"default:0"`
}

type CommentForm struct {
	ArticleId int64  `json:"article_id"`
	Content   string `json:"content" gorm:"size:255"`
}

func GetArticle(id int64, user *User, ignore bool) (*Article, error) {
	var article Article
	err := DB.Where("article_id=?", id).First(&article).Error
	if !ignore && err == nil {
		if article.UserId != user.ID {
			return nil, merrors.NoPermission
		}
	}
	return &article, err
}

func CreateArticle(article *Article, user *User) (*Article, error) {
	saveobj := Article{
		UserId: user.ID,
		Status: 0,
	}
	saveobj.Title = article.Title
	saveobj.Content = article.Content
	saveobj.Attachments = article.Attachments
	saveobj.Stype = article.Stype
	saveobj.Dtype = article.Dtype
	err := DB.Create(&saveobj).Error
	return &saveobj, err
}

func UpdateArticle(article *Article, user *User, ignore bool) (*Article, error) {
	var saveobj *Article
	var err error
	saveobj, err = GetArticle(article.ID, user, ignore)
	if err != nil {
		return saveobj, err
	}
	var kvalues = Article{
		Title:       article.Title,
		Content:     article.Content,
		Attachments: article.Attachments,
		Stype:       article.Stype,
		Dtype:       article.Dtype,
	}
	err = DB.Model(saveobj).Updates(kvalues).Error
	return saveobj, err
}

func DeleteArticle(id int64, user *User, ignore bool) error {
	var article *Article
	var err error
	article, err = GetArticle(id, user, ignore)
	err = DB.Delete(article).Error
	return err
}
