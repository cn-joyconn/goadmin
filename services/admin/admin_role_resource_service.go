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
	joyarray "github.com/cn-joyconn/goutils/array"
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

	db := defaultOrm.DB.Where("f_role_id = ?", roleid)
	if resourceids != nil && len(resourceids) > 0 {
		db = db.Where("f_resource_id IN ?", resourceids)
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
	defaultOrm.DB.Where("f_resource_id = ?", pResource).Find(&models)
	if models != nil && len(models) > 0 {
		result := defaultOrm.DB.Where("f_resource_id = ?", pResource).Delete(&adminModel.AdminRoleResource{})
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
func (service *AdminRoleResourceService) Inserts(roleid int, resourceids []int) int {
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
	var result adminModel.AdminRoleResource
	err:=defaultOrm.DB.Where("f_id = ?", pId).Find(&result).Error
	if err==nil{
		return &result
	}else{
		return nil	
	}
}

/**
*查询一条角色对应的所有功能资源ID
* @param roleid 信息id
* @return  未找到时返回null
 */
func (service *AdminRoleResourceService) SelectByRoleID(roleid int) *[]int {
	var result = make([]int, 0)
	cacheKey := service.getRoleResourcesCacheKey(roleid)
	err := resouceCacheObj.Get(cacheKey, &result)
	if err != nil || result == nil {
		var models []adminModel.AdminRoleResource
		err:=defaultOrm.DB.Where("f_role_id = ?", roleid).Find(&models).Error
		if err == nil {
			for _, model := range models {
				result = append(result, (&model).PResource)
			}
		}
		resouceCacheObj.Put(cacheKey, result, 1000*60*60*24)
	}

	return &result
}
func (service *AdminRoleResourceService) SelectByRoleIDs(roleids []int) *[]*adminModel.AdminRoleResource {
	if roleids == nil {
		return nil
	}
	cacheKeyList := make([]string, 0)
	notExisitIDs := make([]int, 0)
	var err error

	result := make([]*adminModel.AdminRoleResource, 0)
	roleids = joyarray.RemoveDuplicateInt(roleids)
	if roleids != nil {
		for _, roleID := range roleids {
			cacheKeyList = append(cacheKeyList, service.getRoleResourcesCacheKey(roleID))
		}
		if len(cacheKeyList) > 0 {
			var cachedModel *adminModel.AdminRoleResource
			for _, key := range cacheKeyList {
				err = roleResouceCacheObj.Get(key, &cachedModel)
				if err == nil {
					result = append(result, cachedModel)
				}
			}
		}
		for _, roleID := range roleids {
			exisit := false
			for _, resourceObj := range result {
				if resourceObj != nil && roleID == int(resourceObj.PRoleid) {
					exisit = true
					break
				}
			}
			if !exisit {
				notExisitIDs = append(notExisitIDs, roleID)
			}
		}

		if notExisitIDs != nil && len(notExisitIDs) > 0 {
			var models []adminModel.AdminRoleResource
			err=defaultOrm.DB.Where("f_role_id in (?)", notExisitIDs).Find(&models).Error
			if err == nil {
				for _, model := range models {
					cacheKey := service.getRoleResourcesCacheKey(int((&model).PRoleid))
					roleResouceCacheObj.Put(cacheKey, &model, 1000*60*60*24)
					result = append(result, &model)

				}

			}
		}
	}
	return &result

}

/**
*查询一条功能资源对应的所有角色ID
* @param pResource 功能资源
* @return  未找到时返回null
 */
func (service *AdminRoleResourceService) selectByResourceID(pResource int) *[]int {
	var models []adminModel.AdminRoleResource
	var result = make([]int, 0)
	err:=defaultOrm.DB.Where("f_resource_id = ?", pResource).Find(&models).Error
	if err == nil {
		for _, model := range models {
			result = append(result, (&model).PRoleid)
		}
	}
	return &result
}
