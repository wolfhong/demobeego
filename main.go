package main

import (
	"demobeego/models"
	_ "demobeego/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"strings"
)

func main() {
	appname := strings.ToLower(beego.BConfig.AppName)
	runmode := strings.ToLower(beego.AppConfig.String("runmode"))

	//静态资源
	if runmode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/"+appname+"/swagger"] = "swagger"
	}
	beego.SetStaticPath("/"+appname+"/static", "static")

	//数据库
	var err error
	database := beego.AppConfig.String("database")
	models.DB, err = gorm.Open("mysql", database)
	if err != nil {
		panic(err)
	}
	models.DB.SingularTable(true) //表名不要自动命名为复数
	defer models.DB.Close()

	// 自动建表 - dev开发模式
	if runmode == "dev" {
		models.DB = models.DB.Debug()
		migrateDatabase()
	}

	//启动命令输入
	if len(os.Args) >= 2 {
		switch strings.ToLower(os.Args[1]) {
		case "version":
			log.Println("version:1.0.0")
			return
		case "initdb":
			log.Println("Migrate Database...")
			migrateDatabase()
			log.Println("Migrate Database Done")
			return
		default:
			log.Println("Command %s not recognized", os.Args[1])
		}
	}
	//运行
	beego.Run()
}

// dev开发环境，自动建表，自动改表结构
// prod生产环境，调用 docker run xx initdb 初始化表结构
func migrateDatabase() {
	models.DB.AutoMigrate(
		new(models.User),
		new(models.UserInfo),
		new(models.UserToken),
		new(models.Article),
		new(models.Comment),
	)
}
