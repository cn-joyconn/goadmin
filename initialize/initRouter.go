package initialize

import (
	"net/http"

	// middleware "github.com/cn-joyconn/goadmin/middleware"
	auth "github.com/cn-joyconn/goadmin/middleware/auth"
	global "github.com/cn-joyconn/goadmin/models/global"
	routers "github.com/cn-joyconn/goadmin/routers"
	gologs "github.com/cn-joyconn/gologs"
	filetool "github.com/cn-joyconn/goutils/filetool"
	"github.com/gin-gonic/gin"
)

//RegistorRouters 初始化总路由
func RegistorRouters(Router *gin.Engine) {
	selfDir := filetool.SelfDir()
	Router.StaticFS(global.AppConf.ContextPath+"/static", http.Dir(selfDir+"/static/")) // 提供静态文件地址

	contextRouter := Router.Group(global.AppConf.ContextPath)
	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := contextRouter.Group("")
	AuthGroup := contextRouter.Group("")
	PermissioneGroup := contextRouter.Group("")
	AuthGroup.Use(auth.Authorize(global.AppConf.Authorize.LoginUrl))
	PermissioneGroup.Use(auth.Permission())
	JoyAuthorizeGroup := &auth.JoyAuthorizeGroup{GinGroup:AuthGroup}
	JoyPermissionGroup := &auth.JoyPermissionGroup{GinGroup:PermissioneGroup}

	routers.InitCommonRouter(PublicGroup, JoyAuthorizeGroup, JoyPermissionGroup)
	routers.InitAccountRouter(PublicGroup, JoyAuthorizeGroup, JoyPermissionGroup)
	routers.InitHomeRouter(PublicGroup, JoyAuthorizeGroup, JoyPermissionGroup)
	routers.InitAdminResourceRouter(PublicGroup, JoyAuthorizeGroup, JoyPermissionGroup)
	routers.InitAdminRoleRouter(PublicGroup, JoyAuthorizeGroup, JoyPermissionGroup)
	routers.InitAdminMenuRouter(PublicGroup, JoyAuthorizeGroup, JoyPermissionGroup)
	routers.InitAdminUserRouter(PublicGroup, JoyAuthorizeGroup, JoyPermissionGroup)

	// {
	// 	router.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	// }
	// PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	// {
	// 	router.InitApiRouter(PrivateGroup)                   // 注册功能api路由
	// 	router.InitJwtRouter(PrivateGroup)                   // jwt相关路由
	// 	router.InitUserRouter(PrivateGroup)                  // 注册用户路由
	// 	router.InitMenuRouter(PrivateGroup)                  // 注册menu路由
	// 	router.InitEmailRouter(PrivateGroup)                 // 邮件相关路由
	// 	router.InitSystemRouter(PrivateGroup)                // system相关路由
	// 	router.InitCasbinRouter(PrivateGroup)                // 权限相关路由
	// 	router.InitCustomerRouter(PrivateGroup)              // 客户路由
	// 	router.InitAutoCodeRouter(PrivateGroup)              // 创建自动化代码
	// 	router.InitAuthorityRouter(PrivateGroup)             // 注册角色路由
	// 	router.InitSimpleUploaderRouter(PrivateGroup)        // 断点续传（插件版）
	// 	router.InitSysDictionaryRouter(PrivateGroup)         // 字典管理
	// 	router.InitSysOperationRecordRouter(PrivateGroup)    // 操作记录
	// 	router.InitSysDictionaryDetailRouter(PrivateGroup)   // 字典详情管理
	// 	router.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由
	// 	router.InitWorkflowProcessRouter(PrivateGroup)       // 工作流相关接口
	// 	router.InitExcelRouter(PrivateGroup)                 // 表格导入导出
	// }
	gologs.GetLogger("").Info("router register success")
}
