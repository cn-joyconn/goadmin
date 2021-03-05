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
)

var roleResouceCacheObj *gocache.Cache

type AdminRoleResourceService struct {
}

func init() {
	roleResouceCacheObj = &gocache.Cache{
		Catalog:   global.AdminCatalog,
		CacheName: global.AdminCacheName,
	}
}

// 获取缓存用的键
// 返回值 缓存key
func (service *AdminRoleResourceService) getRoleResourcesCacheKey(roleID int) string {
	return "roles_resources_" + strconv.Itoa(roleID)
}
func (service *AdminRoleResourceService) removeRoleResourceIDsCache(roleid int) {
	roleResouceCacheObj.Delete(service.getRoleResourcesCacheKey(roleid))
}

/**
*删除一条角色-功能资源对应信息
* @param
* @return 结果 1:成功 小于1:失败
 */
func (service *AdminRoleResourceService) DeleteByPrimaryKey(roleid int, resourceids []int) int64 {

	db := defaultOrm.DB.Where("PRoleid = ?", roleid)
	if resourceids != nil && len(resourceids) > 0 {
		db = db.Where("PResource IN ?", resourceids)
	}

	result := db.Delete(&adminModel.AdminRoleResource{})
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
		return 0
	}
	if result.RowsAffected > 0 {
		service.removeRoleResourceIDsCache(roleid)
	}
	return result.RowsAffected

}

/**
*删除一批角色-功能资源对应信息
* @param pResource 功能资源id
* @return 结果 1:成功 小于1:失败
 */
func (service *AdminRoleResourceService) DeleteByResourceID(pResource int) int64 {
	var models []*adminModel.AdminRoleResource
	defaultOrm.DB.Where("PResource = ?", pResource).Find(&models)
	if models != nil && len(models) > 0 {
		result := defaultOrm.DB.Where("PResource = ?", pResource).Delete(&adminModel.AdminRoleResource{})
		if result.RowsAffected > 0 {
			for _, model := range models {
				service.removeRoleResourceIDsCache(int(model.PRoleid))
			}
		}
		return result.RowsAffected
	}
	return 0
}

/**
*添加一条角色-功能资源对应信息
* @param records 角色-功能资源对应信息实例
* @return 结果 1:成功 小于1:失败
 */
func (service *AdminRoleResourceService) Inserts(roleid uint32, resourceids []uint32) int {
	records := make([]*adminModel.AdminRoleResource, 0)
	if resourceids != nil && len(resourceids) > 0 {
		for _, id := range resourceids {
			model := &adminModel.AdminRoleResource{}
			records = append(records, model)
			model.PResource = id
			model.PRoleid = roleid
		}
		defaultOrm.DB.Create(&records)
		result := 0
		for _, model := range records {
			if model.PId > 0 {
				result++
			}
		}
		return result
	}
	return 0
}

/**
*查询一条角色-功能资源对应信息
* @param pId 信息id
* @return  未找到时返回null
 */
func (service *AdminRoleResourceService) SelectByPrimaryKey(pId int) *adminModel.AdminRoleResource {
	var result *adminModel.AdminRoleResource
	defaultOrm.DB.Where("PId = ?", pId).Find(&result)
	return result
}

/**
*查询一条角色对应的所有功能资源ID
* @param roleid 信息id
* @return  未找到时返回null
 */
func (service *AdminRoleResourceService) SelectByRoleID(roleid int) []uint32 {
	var result = make([]uint32, 0)
	cacheKey := service.getRoleResourcesCacheKey(roleid)
	err := resouceCacheObj.Get(cacheKey, &result)
	if err != nil || result == nil {
		var models []*adminModel.AdminRoleResource
		defaultOrm.DB.Where("PRoleid = ?", roleid).Find(&models)
		if models != nil {
			for _, model := range models {
				result = append(result, model.PResource)
			}
		}
		resouceCacheObj.Put(cacheKey, result, 1000*60*60*24)
	}

	return result
}

/**
*查询一条功能资源对应的所有角色ID
* @param pResource 功能资源
* @return  未找到时返回null
 */
func (service *AdminRoleResourceService) SelectByResourceID(pResource int) []uint32 {
	var models []*adminModel.AdminRoleResource
	var result = make([]uint32, 0)
	defaultOrm.DB.Where("PResource = ?", pResource).Find(&models)
	if models != nil {
		for _, model := range models {
			result = append(result, model.PRoleid)
		}
	}
	return result
}
