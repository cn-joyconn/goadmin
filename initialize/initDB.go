package initialize

import (
	"time"

	adminModel "github.com/cn-joyconn/goadmin/models/admin"
	adminService "github.com/cn-joyconn/goadmin/services/admin"

	middleware "github.com/cn-joyconn/goadmin/middleware"
	global "github.com/cn-joyconn/goadmin/models/global"

	"fmt"

	defaultOrm "github.com/cn-joyconn/goadmin/models/defaultOrm"

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
		dsn = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s&parseTime=true&loc=Local", dblink.DBUser, dblink.DBPWD, dblink.DBHost, dblink.DBPort, dblink.DBName, dblink.DBCharset)
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
			// LogLevel:      gormlogger.Info,
			LogLevel: gormlogger.Info,
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

	if initAdminUser {
		_initAdminUser()
	}
	_initAdminMenu()

}

//初始化管理员及默认用户
func _initAdminUser() {
	defaultOrm.DB.AutoMigrate(&adminModel.AdminUser{},
		&adminModel.AdminLog{},
		&adminModel.AdminMenu{},
		&adminModel.AdminResource{},
		&adminModel.AdminRole{},
		&adminModel.AdminRoleResource{})
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
			adminUser.Password = adminService.GetSaltPwd("admin123")
			uid, adminUser := adminService.InsertUserModel(adminUser)
			if uid > 0 {
				userObj.ID = uid
				adminUsers = append(adminUsers, uid)
				needSaveAdminConfig = true
			}
			global.Admins.Users[i] = userObj

		}
	}
	if needSaveAdminConfig {
		global.SaveAdminConfig()
	}
	global.SetSuperAdminUsers(adminUsers)
}

//初始化菜单
func _initAdminMenu() {
	var result adminModel.AdminMenu
	err :=defaultOrm.DB.Where(" f_menu_id = ?", 0).First(&result).Error
	if err==nil{
		return
	}
	defualtMenu := &adminModel.AdminMenu{
		PId:          0,            //description:菜单ID"`
		PMenuID:      0,            //description:菜单ID,根节点此值为0"`
		PPid:         0,            //description:父级ID"`
		PState:       1,            //description:状态1正常 0禁用"`
		PCreatuserid: 0,            //description:添加人（用户id）"`
		PName:        "默认菜单",       //description:名称"`
		PDesc:        "",           //description:描述"`
		PURL:         "",           //请求url"`
		PIcon:        "", //description:图标"`
		PPermission:  "",           //description:功能对应的权限标识"`
		PType:        1,            //description:资源类型(1菜单  2页面)"`
		PSort:        1,            //description:排序"`
		PLevel:       0,            //description:层级"`
		PParams:      "",           //description:自定义参数"`
		Children:make([]adminModel.AdminMenu,0),
	}	
	defualtMenu.Children = append(defualtMenu.Children,adminModel.AdminMenu{
		PId:          0,            //description:菜单ID"`
		PMenuID:      0,            //description:菜单ID,根节点此值为0"`
		PPid:         0,            //description:父级ID"`
		PState:       1,            //description:状态1正常 0禁用"`
		PCreatuserid: 0,            //description:添加人（用户id）"`
		PName:        "系统配置",       //description:名称"`
		PDesc:        "",           //description:描述"`
		PURL:         "",           //请求url"`
		PIcon:        "fa  fa-cog", //description:图标"`
		PPermission:  "",           //description:功能对应的权限标识"`
		PType:        1,            //description:资源类型(1菜单  2页面)"`
		PSort:        100,            //description:排序"`
		PLevel:       1,            //description:层级"`
		PParams:      "",           //description:自定义参数"`
		Children:make([]adminModel.AdminMenu,0),
	})
	defualtMenu.Children[0].Children = append(defualtMenu.Children[0].Children, adminModel.AdminMenu{
		PId:          0,            //description:菜单ID"`
		PMenuID:      0,            //description:菜单ID,根节点此值为0"`
		PPid:         0,            //description:父级ID"`
		PState:       1,            //description:状态1正常 0禁用"`
		PCreatuserid: 0,            //description:添加人（用户id）"`
		PName:        "菜单管理",       //description:名称"`
		PDesc:        "菜单配置",           //description:描述"`
		PURL:         "/page/system/authorize/menu/manage",           //请求url"`
		PPermission:  "page:system:menu:manage",           //description:功能对应的权限标识"`
		PIcon:        "", //description:图标"`
		PType:        2,            //description:资源类型(1菜单  2页面)"`
		PSort:        1,            //description:排序"`
		PLevel:       2,            //description:层级"`
		PParams:      "",           //description:自定义参数"`
	})
	defualtMenu.Children[0].Children = append(defualtMenu.Children[0].Children, adminModel.AdminMenu{
		PId:          0,            //description:菜单ID"`
		PMenuID:      0,            //description:菜单ID,根节点此值为0"`
		PPid:         0,            //description:父级ID"`
		PState:       1,            //description:状态1正常 0禁用"`
		PCreatuserid: 0,            //description:添加人（用户id）"`
		PName:        "权限配置",       //description:名称"`
		PDesc:        "权限元数据配置",           //description:描述"`
		PURL:         "/page/system/authorize/resource/manage",           //请求url"`
		PPermission:  "page:system:resource:manage",           //description:功能对应的权限标识"`
		PIcon:        "", //description:图标"`
		PType:        2,            //description:资源类型(1菜单  2页面)"`
		PSort:        2,            //description:排序"`
		PLevel:       2,            //description:层级"`
		PParams:      "",           //description:自定义参数"`
	})
	defualtMenu.Children[0].Children = append(defualtMenu.Children[0].Children, adminModel.AdminMenu{
		PId:          0,            //description:菜单ID"`
		PMenuID:      0,            //description:菜单ID,根节点此值为0"`
		PPid:         0,            //description:父级ID"`
		PState:       1,            //description:状态1正常 0禁用"`
		PCreatuserid: 0,            //description:添加人（用户id）"`
		PName:        "角色管理",       //description:名称"`
		PDesc:        "角色管理",           //description:描述"`
		PURL:         "/page/system/authorize/role/manage",           //请求url"`
		PPermission:  "page:system:role:manage",           //description:功能对应的权限标识"`
		PIcon:        "", //description:图标"`
		PType:        2,            //description:资源类型(1菜单  2页面)"`
		PSort:        3,            //description:排序"`
		PLevel:       2,            //description:层级"`
		PParams:      "",           //description:自定义参数"`
	})
	adminService := new(adminService.AdminMenuService)
	insertResult:=adminService.Insert(defualtMenu)
	if insertResult>0{
		for _,menu:=range defualtMenu.Children{
			(&menu).PMenuID=defualtMenu.PId
			(&menu).PPid=defualtMenu.PId			
			insertResult=adminService.Insert(&menu)
			if insertResult>0{
				for _,menu1:=range (&menu).Children{
					(&menu1).PMenuID=(&menu).PMenuID
					(&menu1).PPid=(&menu).PId				
					insertResult=adminService.Insert(&menu1)
				}
			}
		}
		
	}
}
