package initialize

import (
	"time"

	adminModel "github.com/cn-joyconn/goadmin/models/admin"
	adminService "github.com/cn-joyconn/goadmin/services/admin"

	middleware "github.com/cn-joyconn/goadmin/middleware"
	global "github.com/cn-joyconn/goadmin/models/global"

	defaultOrm "github.com/cn-joyconn/goadmin/models/defaultOrm"
	"fmt"

	mysql "gorm.io/driver/mysql"
	postgres "gorm.io/driver/postgres"
	sqlite "gorm.io/driver/sqlite"
	gorm "gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	schema "gorm.io/gorm/schema"
)

// InitDB 初始化数据库连接
// @title    InitDB
// @description   初始化数据库连接
// @auth      eric.zsp         时间（2021/03/04  16:04 ）
// @param     initAdminUser        bool         "是否初始化默认用户"
func InitDB(initAdminUser bool) {
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
		Logger: &middleware.GormLogger{
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

	if initAdminUser{
		_initAdminUser()
	}
	
}

//初始化管理员及默认用户
func _initAdminUser() {
	defaultOrm.DB.AutoMigrate(&adminModel.AdminUser{},
			&adminModel.AdminLog{},
			&adminModel.AdminMenu{},
			&adminModel.AdminResource{},
			&adminModel.AdminRole{},
			&adminModel.AdminRoleResource{})
	defaultOrm.DB.AutoMigrate(&adminModel.AdminUser{})
	var adminUsers = global.GetSuperAdminUsers()
	needSaveAdminConfig := false
	for i, userObj := range global.Admins.Users {
		if userObj.ID < 1 {
			adminService := new(adminService.AdminUserService)
			adminUser := &adminModel.AdminUser{}
			// adminUser.ID=userObj.ID
			adminUser.Alias = userObj.Alias
			adminUser.Username = userObj.UserName
			adminUser.Phone = userObj.Phone
			adminUser.Email = userObj.Email
			adminUser.Password=adminService.GetSaltPwd("admin123")
			uid,adminUser := adminService.InsertUserModel(adminUser)
			if uid > 0 {
				userObj.ID = uid
				adminUsers = append(adminUsers, uid)
				needSaveAdminConfig=true
			}
			global.Admins.Users[i]=userObj

		}
	}
	if needSaveAdminConfig{
		global.SaveAdminConfig()
	}
	global.SetSuperAdminUsers(adminUsers)
}
