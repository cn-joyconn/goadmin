package admin

import (
	"net/url"
	"strconv"

	controllers "github.com/cn-joyconn/goadmin/controllers"
	"github.com/cn-joyconn/goadmin/middleware/auth"
	adminModel "github.com/cn-joyconn/goadmin/models/admin"
	"github.com/cn-joyconn/goadmin/models/global"
	adminServices "github.com/cn-joyconn/goadmin/services/admin"
	"github.com/gin-gonic/gin"
)

type ResourceController struct {
	controllers.BaseController
}

var ResourceService *adminServices.AdminResourceService

func init() {
	ResourceService = &adminServices.AdminResourceService{}
}
func (controller *ResourceController) ManagePage(c *gin.Context) {
	// fmt.Printf(c.HandlerName())
	data := gin.H{
		"pageTitle": "菜单管理",
	}

	controller.ResponseHtml(c, "authorize/Resource", data)
}

/**
* 删除功能权限
*
* @param id
* @return
 */
func (controller *ResourceController) DeleteByPrimaryKey(c *gin.Context) {
	fId := c.PostForm("id")
	id, err := strconv.Atoi(fId)
	if err != nil {
		controller.ApiErrorCode(c, "参数错误", "", global.ParamsError)
		return
	}
	result := ResourceService.DeleteByPrimaryKey(id)
	if result > 0 {
		controller.ApiSuccess(c, "删除成功", result)
	} else {
		controller.ApiError(c, "删除失败", result)
	}
}

/**
* 添加功能权限
*
* @param record
* @return
 */
func (controller *ResourceController) Insert(c *gin.Context) {
	var record adminModel.AdminResource
	c.ShouldBind(&record)
	PDesc, err1 := url.QueryUnescape(record.PDesc)
	if err1 == nil {
		record.PDesc = PDesc
	}
	PName, err2 := url.QueryUnescape(record.PName)
	if err2 == nil {
		record.PName = PName
	}
	result := ResourceService.Insert(&record)
	if result > 0 {
		result = record.PId
		controller.ApiSuccess(c, "添加成功", result)
	} else {
		controller.ApiError(c, "添加失败", result)
	}
}

/**
* 修改功能权限
*
* @param record
* @return
 */
// @RequestMapping(value = "update", method= RequestMethod.POST)
// @IAuthorization(needPermission = true,resources={"api:system:authorize:resource:update"})
func (controller *ResourceController) UpdateByPrimaryKey(c *gin.Context) {
	var record adminModel.AdminResource
	c.ShouldBind(&record)
	PDesc, err1 := url.QueryUnescape(record.PDesc)
	if err1 == nil {
		record.PDesc = PDesc
	}
	PName, err2 := url.QueryUnescape(record.PName)
	if err2 == nil {
		record.PName = PName
	}
	result := ResourceService.UpdateByPrimaryKey(&record)
	if result > 0 {
		controller.ApiSuccess(c, "修改成功", result)
	} else {
		controller.ApiError(c, "修改失败", result)
	}
}

/**
* 根据ID获取功能权限
*
* @param fId
* @return
 */
func (controller *ResourceController) SelectRightByPrimaryKey(c *gin.Context) {
	fId := c.PostForm("fId")
	id, err := strconv.Atoi(fId)
	if err != nil {
		controller.ApiErrorCode(c, "参数错误", "", global.ParamsError)
		return
	}
	result := ResourceService.SelectByPrimaryKey(id)
	if result != nil {
		controller.ApiSuccess(c, "", result)
	} else {
		controller.ApiError(c, "", result)
	}
}

/**
*  获取所有路径
* @param request
* @return
 */
// @RequestMapping(value = "getAllPermission", method = RequestMethod.GET)
// @IAuthorization(needPermission = true,resources={"api:system:authorize:resource:getAllPermission"})
func (controller *ResourceController) GetAllPermission(c *gin.Context) {
	permissionGroup := &auth.JoyPermissionGroup{}
	permissions := permissionGroup.GetAllPermissionName()
	controller.ApiSuccess(c, "", permissions)
}
