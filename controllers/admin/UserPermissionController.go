package admin

import (
	// "net/url"
	"strconv"

	controllers "github.com/cn-joyconn/goadmin/controllers"
	"github.com/cn-joyconn/goutils/array"
	// "github.com/cn-joyconn/goadmin/middleware/auth"
	adminModel "github.com/cn-joyconn/goadmin/models/admin"
	"github.com/cn-joyconn/goadmin/models/global"
	adminServices "github.com/cn-joyconn/goadmin/services/admin"
	"github.com/gin-gonic/gin"
)

type UserPermissionController struct {
	controllers.BaseController
}
var AdminUserPermissionService *adminServices.AdminUserPermissionService
func init(){
	AdminUserPermissionService = &adminServices.AdminUserPermissionService{}
}


/**
     * 获取用户的所有角色
     *
     * @param appids
     * @return
     */
	//  @RequestMapping(value = "getMyRoles", method= RequestMethod.GET)
	//  @IAuthorization(needPermission = false)
	 func  (controller *UserPermissionController) getMyRoles(c *gin.Context) {
		 uid := controller.GetContextUserId(c)
		 models:=AdminUserPermissionService.GetUserRoles(uid)
		 controller.ApiSuccess(c, "", models) 
	 }
 
	 /**
	  * 根据用户ID获取用户的角色信息
	  *
	  * @param fUserid
	  * @return
	  */
	//  @RequestMapping(value = "get", method= RequestMethod.GET)
	//  @IAuthorization(needPermission = true,resources={"api:system:authorize:userRole:get"})
	 func  (controller *UserPermissionController) selectUserRoles(c *gin.Context) {
		 //int appid, String fUserid
		 fId := c.PostForm("fUserid")
		 fUserid, err1 := strconv.Atoi(fId)
		 if err1 != nil {
			 controller.ApiErrorCode(c, "参数错误", "", global.ParamsError)
			 return
		 }
		 models:=AdminUserPermissionService.GetUserRoles(fUserid)
		 controller.ApiSuccess(c, "", models) 
	 }
 
	 
	 /**
	  * 修改用户的角色信息
	  *
	  * @param record
	  * @return
	  */
	//  @RequestMapping(value = "update", method= RequestMethod.POST)
	//  @IAuthorization(needPermission = true,resources={"api:system:authorize:userRole:update"})
	 func  (controller *UserPermissionController) updateUserRolesByPrimaryKey(c *gin.Context) {
		 //JoyConnAuthenticatePermissionUserModel record
		 var xAdminUserRoleLimit *adminModel.XAdminUserRoleLimit
		 c.ShouldBind(xAdminUserRoleLimit)
		 result:=adminUserService.UpdateUserRoles(xAdminUserRoleLimit.ID,xAdminUserRoleLimit.PRoleObjs)		
		 if result > 0 {
			controller.ApiSuccess(c, "修改成功", result)
		} else {
			controller.ApiError(c, "修改失败", result)
		}
	 }
 
	
 
 
	 /**
	  * 获取用户的权限列表
	  *
	  * @return
	  */
	//  @RequestMapping(value = "getrights", method= RequestMethod.GET)
	//  @IAuthorization(needPermission = false)
	 func  (controller *UserPermissionController) getUserRights(c *gin.Context) {
		uid := controller.GetContextUserIdStr(c)
		models:=AdminUserPermissionService.GetUserResourcesList(uid)
		controller.ApiSuccess(c, "", models) 
	 }
 
	 
 
	 /**
	  * 获取用户的菜单
	  *
	  * @param request
	  * @return
	  */
	//  @RequestMapping(value = "getMyMenuByID", method= RequestMethod.GET)
	//  @IAuthorization(needPermission = false)
	 func  (controller *UserPermissionController) getMyMenuByID(c *gin.Context){
		 //int menuID, HttpServletRequest request
		 fId := c.PostForm("menuID")
		 menuID, err1 := strconv.Atoi(fId)
		 if err1 != nil {
			 controller.ApiErrorCode(c, "参数错误", "", global.ParamsError)
			 return
		 }
		 uid := controller.GetContextUserIdStr(c)
		permissions := AdminUserPermissionService.GetUserPermissions(uid)
		menuModels := MenuService.SelectMenuByMenuID(menuID)
		result := make([]*adminModel.AdminMenu,0)
		 for _,model := range menuModels {
			 if array.InStrArray(model.PPermission,permissions) {
				 result = append(result, model)
			 }
		 }
		controller.ApiSuccess(c, "", result) 
		
		
	 }
