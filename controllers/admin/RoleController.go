package admin

import (
	"net/url"
	"strconv"

	controllers "github.com/cn-joyconn/goadmin/controllers"
	// "github.com/cn-joyconn/goadmin/middleware/auth"
	adminModel "github.com/cn-joyconn/goadmin/models/admin"
	"github.com/cn-joyconn/goadmin/models/global"
	adminServices "github.com/cn-joyconn/goadmin/services/admin"
	"github.com/gin-gonic/gin"
)

type RoleController struct {
	controllers.BaseController
}

var adminRoleService *adminServices.AdminRoleService
var adminRoleResourceService *adminServices.AdminRoleResourceService

func init() {
	adminRoleService = &adminServices.AdminRoleService{}
}
func (controller *RoleController) ManagePage(c *gin.Context) {
	// fmt.Printf(c.HandlerName())
	data := gin.H{
		"pageTitle": "菜单管理",
	}

	controller.ResponseHtml(c, "admin/role.html", data)
}

/**
 * 修改角色状态角色
 *
 * @param
 * @return
 */
func (controller *RoleController) UpdateState(c *gin.Context) {
	// Integer Id, Integer state
	fId := c.PostForm("Id")
	id, err1 := strconv.Atoi(fId)
	if err1 != nil {
		controller.ApiErrorCode(c, "参数错误", "", global.ParamsError)
		return
	}
	fstate := c.PostForm("state")
	state, err2 := strconv.Atoi(fstate)
	if err2 != nil {
		controller.ApiErrorCode(c, "参数错误", "", global.ParamsError)
		return
	}
	result := adminRoleService.UpdateStateByPrimaryKey(id, state)
	if result > 0 {
		controller.ApiSuccess(c, "修改成功", result)
	} else {
		controller.ApiError(c, "修改失败", result)
	}
}

/**
 * 添加角色
 *
 * @param record
 * @return
 */
func (controller *RoleController) InsertRole(c *gin.Context) {
	var record adminModel.AdminRole
	c.ShouldBind(&record)
	PDesc, err1 := url.QueryUnescape(record.PDesc)
	if err1 == nil {
		record.PDesc = PDesc
	}
	PName, err2 := url.QueryUnescape(record.PName)
	if err2 == nil {
		record.PName = PName
	}
	result := adminRoleService.Insert(&record)
	if result > 0 {
		result = record.PId
		controller.ApiSuccess(c, "添加成功", result)
	} else {
		controller.ApiError(c, "添加失败", result)
	}
}

/**
 * 修改角色
 *
 * @param record
 * @return
 */
func (controller *RoleController) UpdatetRoleByPrimaryKey(c *gin.Context) {
	var record adminModel.AdminRole
	c.ShouldBind(&record)
	PDesc, err1 := url.QueryUnescape(record.PDesc)
	if err1 == nil {
		record.PDesc = PDesc
	}
	PName, err2 := url.QueryUnescape(record.PName)
	if err2 == nil {
		record.PName = PName
	}
	result := adminRoleService.UpdateByPrimaryKey(&record)
	if result > 0 {
		controller.ApiSuccess(c, "修改成功", result)
	} else {
		controller.ApiError(c, "修改失败", result)
	}
}

/**
 * 根据角色ID获取角色信息
 *
 * @param fId
 * @return
 */
func (controller *RoleController) SelectRoleByPrimaryKey(c *gin.Context) {
	//Integer fId
	fId := c.PostForm("fId")
	id, err1 := strconv.Atoi(fId)
	if err1 != nil {
		controller.ApiErrorCode(c, "参数错误", "", global.ParamsError)
		return
	}
	result := adminRoleService.SelectByPrimaryKey(id)
	if result != nil {
		controller.ApiSuccess(c, "", result)
	} else {
		controller.ApiError(c, "", result)
	}
}

/**
 * 获取角色列表，分页
 *
 * @param
 * @param pageIndex
 * @param pageSize
 * @return
 */
func (controller *RoleController) SelectRoleListByPage(c *gin.Context) {
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
	uid := controller.GetContextUserId(c)
	creatUser := strconv.Itoa(uid)
	if global.IsSuperAdmin(uid) {
		creatUser = ""
	}
	err, models, count := adminRoleService.SelectByPage(creatUser, pageindex, pagesize)
	if err == nil {
		controller.ApiDataList(c, "查询成功", models, count)

	} else {
		controller.ApiErrorCode(c, "查询成功", models, global.NoResult)
	}
}

/**
 * 获取角色对应的功能权限
 *
 * @param roleid
 * @return
 */
func (controller *RoleController) GetResourceIDsByRoleID(c *gin.Context) {
	//int roleid
	fId := c.PostForm("roleid")
	id, err1 := strconv.Atoi(fId)
	if err1 != nil {
		controller.ApiErrorCode(c, "参数错误", "", global.ParamsError)
		return
	}
	resourceIDs := adminRoleResourceService.SelectByRoleID(id)
	if resourceIDs != nil {
		controller.ApiSuccess(c, "", resourceIDs)
	} else {
		controller.ApiError(c, "", resourceIDs)
	}

}

/**
 * 添加角色对应的功能权限
 *
 * @return
 */
func (controller *RoleController) UpdateRoleResource(c *gin.Context) {
	//Integer roleid,int appid,@RequestParam(value = "resourceids[]") int[] resourceids) {
	fId := c.PostForm("roleid")
	roleid, err1 := strconv.Atoi(fId)
	if err1 != nil {
		controller.ApiErrorCode(c, "参数错误", "", global.ParamsError)
		return
	}
	resourceidStrs, exisit := c.GetPostFormArray("resourceids[]")
	if !exisit {
		controller.ApiErrorCode(c, "参数错误", "", global.ParamsError)
		return
	}
	var resourceids = make([]int, len(resourceidStrs))
	for i, id := range resourceidStrs {
		resourceid, err2 := strconv.Atoi(id)
		if err2 == nil {
			resourceids[i] = resourceid
		}
	}
	adminRoleResourceService.DeleteByPrimaryKey(roleid, nil)
	insertResult := adminRoleResourceService.Inserts(roleid, resourceids)
	if insertResult > 0 {
		controller.ApiSuccess(c, "", insertResult)
	} else {
		controller.ApiError(c, "", insertResult)
	}
}
