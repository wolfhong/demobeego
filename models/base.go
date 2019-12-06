package models

// gorm doc: http://gorm.book.jasperxu.com/

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	// _ "github.com/jinzhu/gorm/dialects/mssql"
)

var DB *gorm.DB

func init() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "t_" + defaultTableName
	}
}

type TimePart struct {
	CreatedAt time.Time  `json:"created_at" gorm:"type:datetime;"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"type:datetime;"`
	DeletedAt *time.Time `json:"-" gorm:"type:datetime"`
}

type IdPart struct {
	ID int64 `json:"id" gorm:"primary_key" ` //数据库ID,便于使用snowflake算法
}
