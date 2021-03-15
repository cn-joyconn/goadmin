package admin

import (
	"net/url"
	"strconv"

	controllers "github.com/cn-joyconn/goadmin/controllers"
	"github.com/cn-joyconn/goadmin/utils/saveFile"
	"github.com/cn-joyconn/goutils/strtool"

	adminModel "github.com/cn-joyconn/goadmin/models/admin"
	"github.com/cn-joyconn/goadmin/models/global"
	adminServices "github.com/cn-joyconn/goadmin/services/admin"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	controllers.BaseController
}

var adminUserService *adminServices.AdminUserService

func init() {
	adminUserService = &adminServices.AdminUserService{}
}

/**
* 获取用户列表
*
* @return
 */
//  @RequestMapping(value = "getUserList", method = RequestMethod.POST)
//  @IAuthorization(needPermission = true,resources = {"api:iotdata:account:user:getUserList"})
func (controller *UserController) getUserList(c *gin.Context) {
	queryID := c.Query("searchID")
	pageIndex := c.DefaultQuery("pageIndex", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	pageindex, err1 := strconv.Atoi(pageIndex)
	if err1 != nil {
		controller.ApiErrorCode(c, "参数错误", "页码参数错误", global.ParamsError)
		return
	}
	pagesize, err2 := strconv.Atoi(pageSize)
	if err2 != nil {
		controller.ApiErrorCode(c, "参数错误", "查询条数参数错误", global.ParamsError)
		return
	}
	if strtool.IsBlank(queryID) {
		models, count := adminUserService.SelectUserList(pagesize, pageindex)
		controller.ApiDataList(c, "查询成功", models, count)
	} else {
		var adminUserModel *adminModel.AdminUser
		if adminUserService.IsPhone(queryID) {
			adminUserModel = adminUserService.GetUserByPhone(queryID)
		} else if adminUserService.IsEmail(queryID) {
			adminUserModel = adminUserService.GetUserByEmail(queryID)
		} else if adminUserService.IsUserName(queryID) {
			adminUserModel = adminUserService.GetUserByUserName(queryID)
		}
		if adminUserModel != nil {
			controller.ApiDataList(c, "查询成功", [...]*adminModel.AdminUser{adminUserModel}, 1)
		} else {
			controller.ApiSuccess(c, "查询成功", gin.H{
				"data":  nil,
				"count": 0,
			})
		}
	}

}

/**
* 获取用户列表
*
* @return
 */
//  @RequestMapping(value = "getUsers", method = RequestMethod.POST)
//  @IAuthorization(needPermission = false)
func (controller *UserController) getUsers(c *gin.Context) {
	var uids []string
	err := c.ShouldBind(&uids)
	if err != nil {
		controller.ApiErrorCode(c, "参数错误", "页码参数错误", global.ParamsError)
		return
	}
	models := adminUserService.GetUserInfoByUserIDS(uids)
	controller.ApiSuccess(c, "查询成功", models)
}

/**
* 获取用户列表
*
* @return
 */
//  @ApiOperation("获取我的用户信息")
//  @RequestMapping(value = "getme", method = RequestMethod.GET)
//  @IAuthorization(needPermission =false)
func (controller *UserController) getme(c *gin.Context) {
	// uid:=controller.GetContextUserId(c)
	userEntity := controller.GetContextUserObj(c)
	controller.ApiSuccess(c, "", userEntity)
}

/**
* 添加一个用户
*
* @param model
* @param request
* @return
 */
//  @RequestMapping(value = "addUserModel", method = RequestMethod.POST)
//  @IAuthorization(needPermission = true,resources = {"api:iotdata:account:user:addUserModel"})
func (controller *UserController) addTConfigUserModel(c *gin.Context) {
	var model *adminModel.AdminUser
	err := c.ShouldBind(&model)
	if err != nil {
		controller.ApiErrorCode(c, "参数错误", "参数错误", global.ParamsError)
		return
	}
	// model:=controller.GetContextUserObj(c)
	Alias, err1 := url.QueryUnescape(model.Alias)
	if err1 == nil {
		model.Alias = Alias
	}
	Description, err2 := url.QueryUnescape(model.Description)
	if err2 == nil {
		model.Description = Description
	}
	Remarks, err2 := url.QueryUnescape(model.Remarks)
	if err2 == nil {
		model.Remarks = Remarks
	}
	RealName, err2 := url.QueryUnescape(model.RealName)
	if err2 == nil {
		model.RealName = RealName
	}
	model.PRoles = ""
	insertResult, result := adminUserService.InsertUserModel(model)
	if insertResult > 0 {
		controller.ApiSuccess(c, "添加成功", result)
	} else {
		controller.ApiError(c, "添加失败", result)
	}

}

/**
* 修改一个后台用户
*
* @param model
* @param request
* @return
 */
//  @RequestMapping(value = "modifyUserModel", method = RequestMethod.POST)
//  @IAuthorization(needPermission = true,resources = {"api:iotdata:account:user:modifyUserModel"})
func (controller *UserController) modifyTConfigUserModel(c *gin.Context) {
	var model *adminModel.AdminUser
	err := c.ShouldBind(&model)
	if err != nil {
		controller.ApiErrorCode(c, "参数错误", "参数错误", global.ParamsError)
		return
	}
	// userEntity:=global.GetContextUserObj(c)
	Alias, err1 := url.QueryUnescape(model.Alias)
	if err1 == nil {
		model.Alias = Alias
	}
	Description, err2 := url.QueryUnescape(model.Description)
	if err2 == nil {
		model.Description = Description
	}
	Remarks, err2 := url.QueryUnescape(model.Remarks)
	if err2 == nil {
		model.Remarks = Remarks
	}
	RealName, err2 := url.QueryUnescape(model.RealName)
	if err2 == nil {
		model.RealName = RealName
	}
	model.PRoles = ""
	insertResult := adminUserService.UpdateUserPubInfo(model)
	controller.ApiSuccess(c, "", insertResult)

}

/**
* 修改用户权限状态
*
* @param uid
* @param stat    lock-锁定 normal-正常
* @param request
* @return
 */
//  @RequestMapping(value = "changeUserStat", method = RequestMethod.POST)
//  @IAuthorization(needPermission = true,resources = {"api:iotdata:account:user:changeUserStat"})
func (controller *UserController) changeUserStat(c *gin.Context) {
	uid := c.PostForm("uid")
	id, err1 := strconv.Atoi(uid)
	stat := c.PostForm("stat")
	state, err2 := strconv.Atoi(stat)
	if err1 != nil || err2 != nil {
		controller.ApiErrorCode(c, "参数错误", "", global.ParamsError)
	}
	updateResult := adminUserService.UpdateUserState(id, uint8(state))
	controller.ApiSuccess(c, "", updateResult)

}

/**
* 修改当前用户密码
*
* @param npwd
* @param pwd
* @return
 */
//  @RequestMapping(value = "modifyPwd", method = RequestMethod.POST)
//  @IAuthorization(needPermission = false)
func (controller *UserController) modifyPwd(c *gin.Context) {
	// uid := c.PostForm("uid")
	// id , err1 := strconv.Atoi(uid)
	pwd := c.PostForm("pwd")
	npwd := c.PostForm("npwd")
	uid := controller.GetContextUserId(c)
	userModel := adminUserService.GetUserByUserID(strconv.Itoa(uid))
	if userModel == nil {
		controller.ApiErrorCode(c, "参数错误", "", global.VerifyFail)
		return
	}
	pwd = adminUserService.GetSaltPwd(pwd)
	if pwd != userModel.Password {
		controller.ApiErrorCode(c, "参数错误", "", global.VerifyFail)
		return
	}
	updateResult := adminUserService.UpdateUserPassword(uid, npwd)
	if updateResult > 0 {
		global.TokenHelper.ClearAuthenticationToken(strconv.Itoa(uid))
	}
	controller.ApiSuccess(c, "", updateResult)

}

/**
* 修改当前用户头像
*
* @return
 */
//  @RequestMapping(value = "modifyUserPhoto", method = RequestMethod.POST)
//  @IAuthorization(needPermission = false)
func (controller *UserController) modifyUserPhoto(c *gin.Context) {
	file, _ := c.FormFile("file")
	newFilePath, returnUrl := saveFile.GetSaveFilePath(file, global.AppConf.Upload)
	err := c.SaveUploadedFile(file, newFilePath)
	// if err == nil {
	// 	controller.ApiSuccess(c, "上传成功", returnUrl)
	// } else {
	// 	controller.ApiSuccess(c, "上传失败", "")
	// }

	if err == nil {
		id := controller.GetContextUserId(c)
		updateResult := adminUserService.UpdateUserHeadPortrait(id, newFilePath)
		if updateResult > 0 {
			controller.ApiSuccess(c, "", returnUrl)
		} else {
			controller.ApiErrorCode(c, "", "", global.ParamsError)
		}

	} else {
		controller.ApiErrorCode(c, "", "", global.ParamsError)
	}
}

/**
* 后台用户重置密码
*
* @param uid
* @param request
* @return
 */
//  @RequestMapping(value = "resetPwd", method = RequestMethod.POST)
//  @IAuthorization(needPermission = true,resources = {"api:iotdata:account:user:resetPwd"})
func (controller *UserController) resetPwd(c *gin.Context) {
	uid := c.PostForm("uid")
	id, _ := strconv.Atoi(uid)
	pwd := c.PostForm("pwd")

	updateResult := adminUserService.UpdateUserPassword(id, adminUserService.GetSaltPwd(pwd))
	if updateResult > 0 {
		global.TokenHelper.ClearAuthenticationToken(uid)
	}
	controller.ApiSuccess(c, "", updateResult)
}
