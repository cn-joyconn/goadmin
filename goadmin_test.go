package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	// "fmt"

	// handle "github.com/cn-joyconn/goadmin/handle"
	initialize "github.com/cn-joyconn/goadmin/initialize"
	adminService "github.com/cn-joyconn/goadmin/services/admin"

	// config "github.com/cn-joyconn/goadmin/utils/config"
	gologs "github.com/cn-joyconn/gologs"
	gin "github.com/gin-gonic/gin"
	snowflake "github.com/cn-joyconn/goutils/snowflake"
	// filetool "github.com/cn-joyconn/goutils/filetool"
)

func TestGoAdmin(t *testing.T) {
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	// r.Use(handle.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	initFilter()
	//启动beego
	// web.Run()
	router := gin.Default()
	router.Run()
}

func TestInitDB(t *testing.T) {
	initialize.Init(func(e *gin.Engine) bool {
		return true
	})
}
func TestQueryUser(t *testing.T) {
	initialize.Init(func(e *gin.Engine) bool {
		return false
	})
	// model := new(adminModel.AdminUser)
	// obj1, err := model.GetAdminUser("1", 1)
	// logMsg1, _ := json.Marshal(obj1)
	// gologs.GetLogger("test").Info(string(logMsg1))
	// logMsg1, _ = json.Marshal(err)
	// gologs.GetLogger("test").Info(string(logMsg1))
	service := new(adminService.AdminUserService)
	obj, err := service.GetUserByUserID("5")
	logMsg, _ := json.Marshal(obj)
	gologs.GetLogger("test").Info(string(logMsg))
	logMsg, _ = json.Marshal(err)
	ii := err != nil
	gologs.GetLogger("test").Info(string(err.Error()))
	gologs.GetLogger("test").Info(strconv.FormatBool(ii))
	obj, _ = service.GetUserByUserName("superManage")
	logMsg, _ = json.Marshal(obj)
	gologs.GetLogger("test").Info(string(logMsg))
	obj, _ = service.GetUserByPhone("18333660110")
	logMsg, _ = json.Marshal(obj)
	gologs.GetLogger("test").Info(string(logMsg))
	obj, _ = service.GetUserByEmail("18333660110@189.cn")
	logMsg, _ = json.Marshal(obj)
	gologs.GetLogger("test").Info(string(logMsg))

}
func TestUpdateUser(t *testing.T) {
	initialize.Init(func(e *gin.Engine) bool {
		return true
	})
	service := new(adminService.AdminUserService)
	obj, _ := service.GetUserByUserID("22")
	logMsg, _ := json.Marshal(obj)
	gologs.GetLogger("test").Info(string(logMsg))

	obj.Phone = "18333660110"
	obj.Alias = "测试用户A"
	obj.Sex = 1
	updateResult := service.UpdateUserPubInfo(obj)
	gologs.GetLogger("test").Info(strconv.Itoa(updateResult))

	obj, _ = service.GetUserByUserID("22")
	logMsg, err := json.Marshal(obj)
	if err != nil {
		gologs.GetLogger("test").Info(err.Error())
	}
	gologs.GetLogger("test").Info(string(logMsg))

	service.DeleteUser(22)
}
func initFilter() {

	// //过滤器：加日志
	// web.InsertFilter("/admin/*",web.BeforeRouter, sysinit.FilterAddLog)

	// //后台权限过滤
	// web.InsertFilter("/admin/*",web.BeforeRouter, sysinit.FilterAdminPermission)

	// //自定义错误页面
	// web.ErrorController(&controllers.ErrorController{})

}

func TestLogin(t *testing.T) {
	initialize.Init(func(e *gin.Engine) bool {
		return true
	})
	service := new(adminService.AdminUserService)
	userObj, _ := service.GetAdminUser("22", 1)
	fmt.Println(userObj.Alias)
}

// func TestBaseConfiger_DefaultString(t *testing.T) {
// 	bc := &config.YamlConfig{}
// 	selfDir := filetool.SelfDir()
// 	bc.Parse(selfDir + "/conf/app.yml")
// 	fmt.Println(bc.DefaultString("app.name", "world"))
// 	fmt.Println(bc.DefaultString("app.contextpath", "world"))
// }

func TestWeb(t *testing.T) {
	initialize.Init(func(e *gin.Engine) bool {
		id := snowflake.NextId()
		fmt.Println(strconv.FormatUint(id, 16))
		return true
	})

}
