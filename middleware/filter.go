package middleware

// import (
// 	"beego_admin/logic"
// 	"beego_admin/models"
// 	"beego_admin/utils"
// 	"github.com/astaxie/beego/context"
// 	"github.com/astaxie/beego/orm"
// 	"strconv"
// )

// //过滤器的代码

// //后台权限过滤
// func FilterAdminPermission(ctx *context.Context){
// 	adminUser := ctx.Input.Session("admin_user")
// 	if adminUser == nil{
// 		//不需要权限控制的路由地址
// 		notPermission := []string{"/admin/login","/admin/dologin"}
// 		flag := false
// 		for _,v := range notPermission{
// 			if v == ctx.Request.RequestURI{
// 				flag = true
// 				break
// 			}
// 		}
// 		if !flag{
// 			ctx.Redirect(302, "/admin/login")
// 		}
// 	}else{
// 		//组装用户的所有角色权限
// 		ruleIds := logic.GetSessionAuth(ctx)
// 		//不过滤超级管理员
// 		if ruleIds[0] != "*"{
// 			adminAuthRule, _ := models.AdminAuthRuleGetRuleBySkipUrl(ctx.Request.RequestURI)
// 			//如果是没有设置权限的路由默认跳过验证
// 			if adminAuthRule.Id != 0{
// 				isPermission := utils.Contains(ruleIds,strconv.Itoa(int(adminAuthRule.Id)))
// 				if isPermission == 0{
// 					//没有权限
// 					ctx.WriteString("{\"code\": -1,\"msg\": \"没有权限\",\"data\": \"\"}")
// 					return
// 				}
// 			}
// 		}
// 	}
// }

// //自动添加请求日志
// func FilterAddLog(ctx *context.Context){
// 	adminUserSession := ctx.Input.Session("admin_user")
// 	var adminUserId uint
// 	var username string
// 	if adminUserSession != nil{
// 		adminUser := adminUserSession.(*models.AdminUser)
// 		adminUserId = adminUser.Id
// 		username = adminUser.Username
// 	}

// 	adminLog := models.AdminLog{
// 		AdminUserId :adminUserId,
// 		Username :username,
// 		Url : ctx.Request.RequestURI,
// 		Title :"",
// 		Content :"",
// 		Ip :ctx.Request.RemoteAddr,
// 		Useragent :ctx.Request.UserAgent(),
// 	}
// 	m := orm.NewOrm()
// 	_, _= m.Insert(&adminLog)
// }