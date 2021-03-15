package admin

import (
	// "encoding/json"
	// "crypto/md5"
	"strconv"

	adminModel "github.com/cn-joyconn/goadmin/models/admin"
	defaultOrm "github.com/cn-joyconn/goadmin/models/defaultOrm"
	global "github.com/cn-joyconn/goadmin/models/global"
	gocache "github.com/cn-joyconn/gocache"
	gologs "github.com/cn-joyconn/gologs"
	strtool "github.com/cn-joyconn/goutils/strtool"
)

var menuCacheObj *gocache.Cache

type AdminMenuService struct {
}

func init() {
	menuCacheObj = &gocache.Cache{
		Catalog:   global.AdminCatalog,
		CacheName: global.AdminCacheName,
	}
}

// 获取缓存用的键
// 返回值 缓存key
func (service *AdminMenuService) getMenuCacheKey(pId int) string {
	return "menu_" + strconv.Itoa(pId)
}
func (service *AdminMenuService) removeMenuCache(pId int) {
	menuCacheObj.Delete(service.getMenuCacheKey(pId))
}

/**
*修改菜单状态菜单
* @param pId 菜单id
* @param pState 状态 1正常 0禁用
* @return 结果 1:成功 小于1:失败
 */
func (service *AdminMenuService) UpdateStateByPrimaryKey(pId int, pState int) int64 {
	record := &adminModel.AdminMenu{}
	record.PId = pId
	record.PState = pState
	result := defaultOrm.DB.Model(&record).Select("PState").Updates(record)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
		return 0
	}
	if result.RowsAffected > 0 {
		service.removeMenuCache(pId)
	}
	return result.RowsAffected
}

/**
*添加菜单
* @param record 菜单实例
* @return 结果 1:成功 小于1:失败
 */
func (service *AdminMenuService) Insert(record *adminModel.AdminMenu) int {
	result := defaultOrm.DB.Model(&record).Create(record)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
	}
	if record.PId > 0 {
		return record.PId
	} else {
		return -1
	}
}

/**
*获取整个菜单
* @param pId 菜单id
* @return  未找到时返回null
 */
func (service *AdminMenuService) SelectMenuByID(pId int) *adminModel.AdminMenu {
	var result *adminModel.AdminMenu
	defaultOrm.DB.Where("PId = ?", pId).First(&result)
	return result
}

/**
*获取整个菜单
* @param menuId 菜单id
* @return  未找到时返回null
 */
func (service *AdminMenuService) SelectMenuByMenuID(menuId int) []*adminModel.AdminMenu {
	cacheKey := service.getMenuCacheKey(menuId)
	var result []*adminModel.AdminMenu
	err := menuCacheObj.Get(cacheKey, &result)
	if err != nil || result == nil {
		defaultOrm.DB.Where("PMenuID = ? and PState=1", menuId).Find(&result)
		if result != nil {
			menuCacheObj.Put(cacheKey, result, 1000*60*60*24)
		}
	}
	return result
}

/**
*查询菜单
* @param pAppid 应用id
* @param offset 偏移量
* @param limit 查询结果条数
* @return  未找到时返回null
 */
func (service *AdminMenuService) SelectRootByPage(creatUser string, pageIndex int, pageSize int) ([]*adminModel.AdminMenu, int64) {
	var result []*adminModel.AdminMenu
	var count int64
	db := defaultOrm.DB
	if !strtool.IsBlank(creatUser) {
		db = db.Where("PCreatuserid = ?", creatUser)
	}
	db.Model(&adminModel.AdminMenu{}).Count(&count)
	db.Order("PId desc").Limit(pageIndex).Offset((pageIndex - 1) * pageSize).Find(&result)
	return result, count
}

/**
*修改菜单
* @param record 菜单实例
* @return 结果 1:成功 小于1:失败
 */
func (service *AdminMenuService) UpdateByPrimaryKey(record *adminModel.AdminMenu) int64 {
	result := defaultOrm.DB.Model(&record).Select("PName", "PDesc", "PURL", "PIcon", "PPermission", "PType", "PSort", "PLevel", "PParams").Updates(record)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
		return 0
	}
	if result.RowsAffected > 0 {
		service.removeMenuCache(record.PId)
	}
	return result.RowsAffected
}

/**
*删除菜单
* @param menuID 菜单id
* @return 结果 1:成功 小于1:失败
 */
func (service *AdminMenuService) DeleteByMenuID(menuID int) int64 {
	result := defaultOrm.DB.Where("PMenuID = ?", menuID).Delete(&adminModel.AdminRoleResource{})
	if result.RowsAffected > 0 {
		service.removeMenuCache(menuID)
	}
	return result.RowsAffected
}

/**
*删除菜单节点
* @param menuID 菜单id
* @param pId 菜单节点id
* @return 结果 1:成功 小于1:失败
 */
func (service *AdminMenuService) DeleteByPID(menuID int, pId int) int64 {
	result := defaultOrm.DB.Where("PId = ? AND PMenuID = ? ", pId, menuID).Delete(&adminModel.AdminRoleResource{})
	if result.RowsAffected > 0 {
		service.removeMenuCache(menuID)
	}
	return result.RowsAffected
}
