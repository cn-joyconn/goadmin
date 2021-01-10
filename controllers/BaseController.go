package controllers

import (
	// "beego_xadmin/logic"
	// xadminModels "beego_xadmin/models"
	// beegoWeb "github.com/beego/beego/v2/server/web"
	// "github.com/beego/beego/v2/client/orm"
	// config "github.com/beego/beego/v2/core/config"
	// "strconv"
	// "strings"
)

type BaseController struct {
	// beegoWeb.Controller
	controllerName string             //当前控制名称
	actionName     string             //当前action名称
}

// // 设置模板
// // 第一个参数模板，第二个参数为layout
// func (c *BaseController) SetTpl(template ...string) {
// 	var tplName string
// 	layout := "common/admin_layout.html"
// 	switch {
// 	case len(template) == 1:
// 		tplName = template[0]
// 	case len(template) == 2:
// 		tplName = template[0]
// 		layout = template[1]
// 	default:
// 		//不要Controller这个10个字母
// 		ctrlName := strings.ToLower(c.controllerName[0 : len(c.controllerName)-10])
// 		actionName := strings.ToLower(c.actionName)
// 		tplName = ctrlName + "/" + actionName + ".html"
// 	}

// 	c.Layout = layout
// 	c.TplName = tplName
// }

// //JSON输出
// func (c *BaseController) ApiJson(code int16,msg string,data interface{}){
// 	if data == nil{
// 		data = ""
// 	}
// 	c.Data["json"] = map[string]interface{}{"code": code,"msg": msg,"data": data}
// 	c.ServeJSON()
// 	c.StopRun()
// }
// //返回成功的API成功
// func (c *BaseController) ApiSuccess(msg string,data interface{} ){
// 	c.ApiJson(0,msg,data)
// }
// //返回失败的API请求
// func (c *BaseController) ApiError(msg string,data interface{})  {
// 	c.ApiJson(-1,msg,data)
// }
// //返回失败且带code的API请求
// func (c *BaseController) ApiErrorCode(msg string,data interface{},code int16)  {
// 	c.ApiJson(code,msg,data)
// }

// //获取默认每页显示的数量
// func (c *BaseController) GetDefaultPageSize() int{	
// 	pageSize,_ := config.Int("page_size")
// 	return pageSize
// }

// //计算分页的偏移量
// func (c *BaseController) GetPageOffset(pageNo int,pageSize int) int{
// 	return (pageNo - 1) * pageSize
// }

// //请求出错
// func (c *BaseController) RequestError(msg string,route string){
// 	c.Redirect(route, 302)
// }

// //请求成功
// func (c *BaseController) RequestSuccess(msg string,route string){
// 	c.Redirect(route, 302)
// }

// //获取权限列表
// func (c *BaseController) GetRuleList() ([]models.TreeMenu){
// 	//所有权限列表
// 	adminAuthRules,_ := models.AdminAuthRuleGetRule()
// 	return logic.MakeRuleTree(adminAuthRules,0)
// }




