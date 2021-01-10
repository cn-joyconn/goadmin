package models

import (
	"time"

	admin "github.com/cn-joyconn/goadmin/models/admin"

	handle "github.com/cn-joyconn/goadmin/handle"
	global "github.com/cn-joyconn/goadmin/models/global"

	defaultOrm "github.com/cn-joyconn/goadmin/models/defaultOrm"
	// utils "beego-xadmin/models/utils"
	"fmt"

	// _ "gorm.io/database/sql"
	mysql "gorm.io/driver/mysql"
	postgres "gorm.io/driver/postgres"
	sqlite "gorm.io/driver/sqlite"
	gorm "gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	schema "gorm.io/gorm/schema"
)

func InitDB() {
	var dsn string
	var dblink global.DBlink
	var db *gorm.DB
	var driver gorm.Dialector
	var err error
	if global.DBConf.DBType == "mysql" {
		dblink = global.DBConf.Mysql
		dsn = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", dblink.DBUser, dblink.DBPWD, dblink.DBHost, dblink.DBPort, dblink.DBName, dblink.DBCharset)
		driver = mysql.Open(dsn)
	} else if global.DBConf.DBType == "postgres" {
		dblink = global.DBConf.Postgres
		dsn = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", dblink.DBHost, dblink.DBPort, dblink.DBUser, dblink.DBName, dblink.DBPWD)
		driver = postgres.Open(dsn)
	} else if global.DBConf.DBType == "sqlite3" {
		dblink = global.DBConf.Sqlite3
		dsn = dblink.DBName
		driver = sqlite.Open(dsn)

	} else {
		panic("数据库类型错误")
	}

	var durationseconds time.Duration = time.Duration(global.DBConf.SlowThreshold) * time.Second //慢查询阈值 0不记录   单位秒
	db, err = gorm.Open(driver, &gorm.Config{
		SkipDefaultTransaction: global.DBConf.SkipDefaultTransaction, //禁用事务
		PrepareStmt:            true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   global.DBConf.DBDtPrefix, // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,                     // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
		Logger: &handle.GormLogger{
			SlowThreshold: durationseconds,
			LogLevel:      gormlogger.Info,
		},
	})
	// // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	// db.SetMaxIdleConns(global.DBConf.MaxIdleConns)

	// // SetMaxOpenConns sets the maximum number of open connections to the database.
	// db.SetMaxOpenConns(global.DBConf.MaxOpenConns)

	// // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	// db.SetConnMaxLifetime(global.DBConf.ConnMaxLifetime)

	if err != nil {
		panic(err)
	}

	defaultOrm.DB = db

	// //开启日志
	// if global.DBConf.DBLog {
	// 	db.LogMode(true)
	// }
	//create table
	// orm.RunSyncdb("default", false, true)
	InitAdminUser()
}

//初始化
func InitAdminUser() {
	defaultOrm.DB.AutoMigrate(&admin.AdminUser{})
	var adminUsers = global.GetSuperAdminUsers()
	for i, userObj := range global.Admins.Users {
		if userObj.ID < 1 {
			adminUser := &admin.AdminUser{}
			// adminUser.ID=userObj.ID
			adminUser.Alias = userObj.Alias
			adminUser.Username = userObj.UserName
			adminUser.Phone = userObj.Phone
			adminUser.Email = userObj.Email
			uid := adminUser.Insert()
			if uid > 0 {
				userObj.ID = uid
				adminUsers = append(adminUsers, uid)
			}
			global.Admins.Users[i]=userObj

		}
	}
	global.SaveAdminConfig()
	global.SetSuperAdminUsers(adminUsers)
}
