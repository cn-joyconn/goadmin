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

type MenuController struct {
	controllers.BaseController
}

var MenuService *adminServices.AdminMenuService

func init() {
	MenuService = &adminServices.AdminMenuService{}
}

func (controller *MenuController) ManagePage(c *gin.Context) {
	// fmt.Printf(c.HandlerName())
	data := gin.H{
		"pageTitle": "菜单管理",
	}

	controller.ResponseHtml(c, "admin/menu.html", data)
}

/**
 * 修改菜单状态菜单
 *
 * @param
 * @return
 */
func (controller *MenuController) UpdateState(c *gin.Context) {
	//Integer Id, Integer state, HttpServletRequest request
	fId := c.PostForm("Id")
	Id, err1 := strconv.Atoi(fId)
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
	//  uid := controller.GetContextUserId(c)

	result := MenuService.UpdateStateByPrimaryKey(Id, state)
	if result > 0 {
		controller.ApiSuccess(c, "修改成功", result)
	} else {
		controller.ApiError(c, "修改失败", result)
	}
}

/**
 * 删除菜单节点
 *
 * @param
 * @return
 */
func (controller *MenuController) DeleteMenuNode(c *gin.Context) {
	//Integer menuID, Integer pId, HttpServletRequest request
	fId := c.PostForm("menuID")
	menuID, err1 := strconv.Atoi(fId)
	if err1 != nil {
		controller.ApiErrorCode(c, "参数错误", "", global.ParamsError)
		return
	}
	fId = c.PostForm("pId")
	pId, err2 := strconv.Atoi(fId)
	if err2 != nil {
		controller.ApiErrorCode(c, "参数错误", "", global.ParamsError)
		return
	}
	//  uid := controller.GetContextUserId(c)

	result := MenuService.DeleteByPID(menuID, pId)
	if result > 0 {
		controller.ApiSuccess(c, "删除成功", result)
	} else {
		controller.ApiError(c, "删除失败", result)
	}

}

/**
 * 删除菜单节点
 *
 * @param
 * @return
 */
func (controller *MenuController) DeleteMenu(c *gin.Context) {
	//Integer menuID,  HttpServletRequest request
	fId := c.PostForm("menuID")
	menuID, err1 := strconv.Atoi(fId)
	if err1 != nil {
		controller.ApiErrorCode(c, "参数错误", "", global.ParamsError)
		return
	}

	result := MenuService.DeleteByMenuID(menuID)
	if result > 0 {
		controller.ApiSuccess(c, "删除成功", result)
	} else {
		controller.ApiError(c, "删除失败", result)
	}

}

/**
 * 添加菜单
 *
 * @param record
 * @return
 */
func (controller *MenuController) InsertMenu(c *gin.Context) {
	var record adminModel.AdminMenu
	c.ShouldBind(&record)
	PDesc, err1 := url.QueryUnescape(record.PDesc)
	if err1 == nil {
		record.PDesc = PDesc
	}
	PName, err2 := url.QueryUnescape(record.PName)
	if err2 == nil {
		record.PName = PName
	}
	record.PCreatuserid = controller.GetContextUserId(c)
	result := MenuService.Insert(&record)
	if result > 0 {
		result = record.PId
		controller.ApiSuccess(c, "添加成功", result)
	} else {
		controller.ApiError(c, "添加失败", result)
	}

}

/**
 * 修改菜单
 *
 * @param record
 * @return
 */
func (controller *MenuController) UpdatetMenuByPrimaryKey(c *gin.Context) {
	//JoyConnAuthenticatePermissionMenuModel record, HttpServletRequest request
	var record adminModel.AdminMenu
	c.ShouldBind(&record)
	PDesc, err1 := url.QueryUnescape(record.PDesc)
	if err1 == nil {
		record.PDesc = PDesc
	}
	PName, err2 := url.QueryUnescape(record.PName)
	if err2 == nil {
		record.PName = PName
	}
	result := MenuService.UpdateByPrimaryKey(&record)
	if result > 0 {
		controller.ApiSuccess(c, "修改成功", result)
	} else {
		controller.ApiError(c, "修改失败", result)
	}

}

/**
 * 根据菜单ID获取菜单
 *
 * @param menuID
 * @return
 */
func (controller *MenuController) GetMenu(c *gin.Context) {
	//Integer menuID, HttpServletRequest request
	fId := c.Query("menuID")
	menuID, err1 := strconv.Atoi(fId)
	if err1 != nil {
		controller.ApiErrorCode(c, "参数错误", "", global.ParamsError)
		return
	}

	result := MenuService.SelectMenuByMenuID(menuID)
	if result != nil {
		controller.ApiSuccess(c, "", result)
	} else {
		controller.ApiError(c, "", result)
	}

}

/**
 * 查询菜单列表
 *
 * @param menuID
 * @return
 */
func (controller *MenuController) GetRootByPage(c *gin.Context) {
	//Integer menuID, HttpServletRequest request
	fId := c.Query("pageIndex")
	pageIndex, err1 := strconv.Atoi(fId)
	if err1 != nil {
		controller.ApiErrorCode(c, "参数错误", "", global.ParamsError)
		return
	}
	fId = c.Query("pageSize")
	pageSize, err2 := strconv.Atoi(fId)
	if err2 != nil {
		controller.ApiErrorCode(c, "参数错误", "", global.ParamsError)
		return
	}
	userID := controller.GetContextUserId(c)
	if global.IsSuperAdmin(userID){
		userID = 0
	}
	err, result, count := MenuService.SelectRootByPage(strconv.Itoa(userID), pageIndex, pageSize)
	if err == nil {
		controller.ApiDataList(c, "", result, count)
	} else {
		controller.ApiErrorCode(c, "", result, global.NoResult)
	}

}
