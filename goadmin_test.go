package main

import (
	"encoding/json"
	"strconv"
	"testing"

	// handle "github.com/cn-joyconn/goadmin/handle"
	modles "github.com/cn-joyconn/goadmin/models"
	admin "github.com/cn-joyconn/goadmin/models/admin"
	gologs "github.com/cn-joyconn/gologs"
	gin "github.com/gin-gonic/gin"
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
	modles.InitDB()
}
func TestQueryUser(t *testing.T) {
	modles.InitDB()
	obj :=new(admin.AdminUser)
	obj =obj.SelectByUserID(22)
	logMsg,_:=json.Marshal(obj)
	gologs.GetLogger("test").Info(string(logMsg))
	obj =obj.SelectByUserName("superManage")
	logMsg,_=json.Marshal(obj)
	gologs.GetLogger("test").Info(string(logMsg))
	obj =obj.SelectByPhone("18333660110")
	logMsg,_=json.Marshal(obj)
	gologs.GetLogger("test").Info(string(logMsg))
	obj =obj.SelectByEmail("18333660110@189.cn")
	logMsg,_=json.Marshal(obj)
	gologs.GetLogger("test").Info(string(logMsg))
	
}
func TestUpdateUser(t *testing.T) {
	modles.InitDB()
	obj :=new(admin.AdminUser)
	obj =obj.SelectByUserID(22)
	logMsg,_:=json.Marshal(obj)
	gologs.GetLogger("test").Info(string(logMsg))

	obj.Phone="18333660110"
	obj.Alias="测试用户A"
	obj.Sex=1
	updateResult :=obj.UpdateInfo()
	gologs.GetLogger("test").Info( strconv.FormatInt(updateResult,10))
	
	obj =obj.SelectByUserID(22)
	logMsg,err:=json.Marshal(obj)
	if err!=nil{
		gologs.GetLogger("test").Info(err.Error())
	}
	gologs.GetLogger("test").Info(string(logMsg))

	
	 obj.DeleteByUserID(22)
}
func initFilter() {

	// //过滤器：加日志
	// web.InsertFilter("/admin/*",web.BeforeRouter, sysinit.FilterAddLog)

	// //后台权限过滤
	// web.InsertFilter("/admin/*",web.BeforeRouter, sysinit.FilterAdminPermission)

	// //自定义错误页面
	// web.ErrorController(&controllers.ErrorController{})

}
