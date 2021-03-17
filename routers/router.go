package initialize

import (
	// "net/http"

	controllers "github.com/cn-joyconn/goadmin/controllers"
	adminControllers "github.com/cn-joyconn/goadmin/controllers/admin"
	"github.com/cn-joyconn/goadmin/middleware/auth"

	// filetool "github.com/cn-joyconn/goutils/filetool"
	// gologs "github.com/cn-joyconn/gologs"
	// middleware "github.com/cn-joyconn/goadmin/middleware"
	"github.com/gin-gonic/gin"
)

//通用
func InitCommonRouter(publicGroup *gin.RouterGroup, authGroup *auth.JoyAuthorizeGroup, permissioneGroup *auth.JoyPermissionGroup) {
	controller := &controllers.CommonController{}
	publicGroup.GET("/api/CommApi/authimage", controller.AuthImage)
	authGroup.POST("/api/CommApi/fileUploadApi", "", controller.Upload)
	authGroup.POST("/api/CommApi/filesUploadApi", "", controller.UploadFiles)
}

//首页
func InitHomeRouter(publicGroup *gin.RouterGroup, authGroup *auth.JoyAuthorizeGroup, permissioneGroup *auth.JoyPermissionGroup) {
	controller := &controllers.IndexController{}
	authGroup.GET("", "", controller.Index)
	authGroup.GET("/", "", controller.Index)
	authGroup.GET("/home", "", controller.Index)
	authGroup.GET("/home/index", "", controller.Index)
}

//登录
func InitAccountRouter(publicGroup *gin.RouterGroup, authGroup *auth.JoyAuthorizeGroup, permissioneGroup *auth.JoyPermissionGroup) {
	controller := &controllers.AccountController{}
	publicGroup.GET("/page/account/login", controller.LoginPage)
	publicGroup.POST("/api/joyconn/authorize/AuthenticationApi/dologin", controller.LoginForCookie)
	publicGroup.POST("/api/joyconn/authorize/AuthenticationApi/dologinAuth", controller.LoginForAuth)
}

//用户
func InitAdminUserRouter(publicGroup *gin.RouterGroup, authGroup *auth.JoyAuthorizeGroup, permissioneGroup *auth.JoyPermissionGroup) {
	controller := &adminControllers.UserController{}

	authGroup.GET("/page/system/account/modifyphoto", "", controller.ModifyphotoPage)
	permissioneGroup.GET("/page/system/account/UserManage", "page:system:account:usermanage", controller.ManagePage)

	authGroup.POST("/api/system/account/user/getUsers", "", controller.GetUsers)
	authGroup.GET("/api/system/account/user/getMe", "", controller.GetMe)
	authGroup.POST("/api/system/account/user/modifyPwd", "", controller.ModifyPwd)
	authGroup.POST("/api/system/account/user/modifyUserPhoto", "", controller.ModifyUserPhoto)

	permissioneGroup.GET("/api/system/account/user/getUserList", "api:system:account:user:getUserList", controller.GetUserList)
	permissioneGroup.GET("/api/system/account/user/addUser", "api:system:account:user:addUser", controller.AddUserModel)
	permissioneGroup.GET("/api/system/account/user/modifyUserModel", "api:system:account:user:modifyUserModel", controller.ModifyUserModel)
	permissioneGroup.GET("/api/system/account/user/changeUserStat", "api:system:account:user:changeUserStat", controller.ChangeUserStat)
	permissioneGroup.GET("/api/system/account/user/resetPwd", "api:system:account:user:resetPwd", controller.ResetPwd)

	upcontroller := &adminControllers.UserPermissionController{}
	authGroup.GET("/api/system/authorize/user/permission/getMyRoles", "", upcontroller.GetMyRoles)
	authGroup.GET("/api/system/authorize/user/permission/getrights", "", upcontroller.GetUserRights)
	authGroup.GET("/api/system/authorize/user/permission/getMyMenuByID", "", upcontroller.GetMyMenuByID)
	permissioneGroup.GET("/api/system/authorize/user/permission/getUserRoles", "api:system:authorize:user:permission:getUserRoles", upcontroller.SelectUserRoles)
	permissioneGroup.POST("/api/system/authorize/user/permission/update", "api:system:authorize:user:permission:update", upcontroller.UpdateUserRolesByPrimaryKey)
}

//权限
func InitAdminResourceRouter(publicGroup *gin.RouterGroup, authGroup *auth.JoyAuthorizeGroup, permissioneGroup *auth.JoyPermissionGroup) {
	controller := &adminControllers.ResourceController{}

	permissioneGroup.GET("/page/system/authorize/resource/manage", "page:system:resource:manage", controller.ManagePage)

	permissioneGroup.POST("/api/system/authorize/resource/delete", "api:system:authorize:resource:delete", controller.DeleteByPrimaryKey)
	permissioneGroup.POST("/api/system/authorize/resource/add", "api:system:authorize:resource:add", controller.Insert)
	permissioneGroup.POST("/api/system/authorize/resource/update", "api:system:authorize:resource:update", controller.UpdateByPrimaryKey)
	permissioneGroup.GET("/api/system/authorize/resource/getone", "api:system:authorize:resource:getone", controller.SelectRightByPrimaryKey)
	permissioneGroup.GET("/api/system/authorize/resource/getAllPermission", "api:system:authorize:resource:getAllPermission", controller.GetAllPermission)
}

//角色
func InitAdminRoleRouter(publicGroup *gin.RouterGroup, authGroup *auth.JoyAuthorizeGroup, permissioneGroup *auth.JoyPermissionGroup) {
	controller := &adminControllers.RoleController{}

	permissioneGroup.GET("/page/system/authorize/role/manage", "page:system:role:manage", controller.ManagePage)

	permissioneGroup.POST("/api/system/authorize/role/updateState", "api:system:authorize:role:updateState", controller.UpdateState)
	permissioneGroup.POST("/api/system/authorize/role/add", "api:system:authorize:role:add", controller.InsertRole)
	permissioneGroup.POST("/api/system/authorize/role/update", "api:system:authorize:role:update", controller.UpdatetRoleByPrimaryKey)
	permissioneGroup.GET("/api/system/authorize/role/getone", "api:system:authorize:role:getone", controller.SelectRoleByPrimaryKey)
	permissioneGroup.GET("/api/system/authorize/role/getpage", "api:system:authorize:role:getpage", controller.SelectRoleByPrimaryKey)
	permissioneGroup.GET("/api/system/authorize/role/getResourceIDsByRoleID", "api:system:authorize:role:getResourceIDsByRoleID", controller.GetResourceIDsByRoleID)
	permissioneGroup.POST("/api/system/authorize/role/updateRoleResource", "api:system:authorize:role:updateRoleResource", controller.UpdateRoleResource)

}

//菜单
func InitAdminMenuRouter(publicGroup *gin.RouterGroup, authGroup *auth.JoyAuthorizeGroup, permissioneGroup *auth.JoyPermissionGroup) {
	controller := &adminControllers.MenuController{}

	permissioneGroup.GET("/page/system/authorize/menu/manage", "page:system:menu:manage", controller.ManagePage)

	permissioneGroup.POST("/api/system/authorize/menu/updateState", "api:system:authorize:menu:updateState", controller.UpdateState)
	permissioneGroup.POST("/api/system/authorize/menu/deleteNode", "api:system:authorize:menu:deleteNode", controller.DeleteMenuNode)
	permissioneGroup.POST("/api/system/authorize/menu/deleteMenu", "api:system:authorize:menu:deleteMenu", controller.DeleteMenu)
	permissioneGroup.POST("/api/system/authorize/menu/add", "api:system:authorize:menu:add", controller.InsertMenu)
	permissioneGroup.POST("/api/system/authorize/menu/update", "api:system:authorize:menu:update", controller.UpdatetMenuByPrimaryKey)
	permissioneGroup.GET("/api/system/authorize/menu/getMenu", "api:system:authorize:menu:getMenu", controller.GetMenu)
}
