package models

import (
	"demobeego/ground/encrypt"
	"demobeego/ground/timetools"
	"demobeego/models/merrors"
	uuid "github.com/satori/go.uuid"
	"log"
	"strings"
	"time"
)

var (
	secret = "this is demobeego"
)

type User struct {
	IdPart
	TimePart
	Name     string `json:"name" gorm:"size:100;not null;"`
	Password string `json:"-" gorm:"size:64"`
	Email    string `json:"email" gorm:"type:varchar(127);"`
	Mobile   string `json:"mobile" gorm:"type:varchar(31);"`
	Sex      uint   `json:"sex" gorm:"default:0"`
	Status   uint   `json:"status" gorm:"default:0"`
	Profile  UserInfo
}

type UserInfo struct {
	IdPart
	TimePart
	UserId   int64
	Position string    `json:"position" gorm:"size:255"`
	Company  string    `json:"company" gorm:"size:255"`
	Address  string    `json:"address" gorm:"size:255"`
	Birthday time.Time `json:"birthday" gorm:"type:datetime;"`
}

type UserToken struct {
	IdPart
	UserId int64
	Token  string `gorm:"size:32"`
	Expire time.Time
}

func EncryptPass(before string) string {
	return encrypt.Sha1([]byte(secret + before))
}

func DeleteToken(token string) {
	if token == "" {
		return
	}
	var tokenobj UserToken
	err := DB.Where("token=?", token).First(&tokenobj).Error
	if err == nil {
		expireTime := timetools.Str2Time("2000-01-01 00:00:00+08:00")
		log.Println(tokenobj.Expire, expireTime)
		if tokenobj.Expire.After(expireTime) {
			//tokenobj.Expire = expireTime
			//DB.Save(&tokenobj)
			DB.Model(&tokenobj).Update("expire", expireTime)
		}
	}
}

func Token2User(token string, user *User) error {
	var tokenobj UserToken
	var err error
	err = DB.Where("token=?", token).First(&tokenobj).Error
	if err != nil {
		return merrors.TokenNotExist
	}
	current := time.Now()
	if tokenobj.Expire.Before(current) {
		return merrors.TokenExpired
	}
	err = DB.Where("id=?", tokenobj.UserId).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}

//token length 32
func (user *User) CreateToken() string {
	var err error
	uid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	tokenobj := UserToken{
		UserId: user.ID,
		Token:  strings.Replace(uid.String(), "-", "", -1),
		Expire: time.Now().Add(time.Hour * 2),
	}
	err = DB.Create(&tokenobj).Error
	if err != nil {
		panic(err)
	}
	return tokenobj.Token
}
