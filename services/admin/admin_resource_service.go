package admin

import (
	"strconv"

	adminModel "github.com/cn-joyconn/goadmin/models/admin"
	defaultOrm "github.com/cn-joyconn/goadmin/models/defaultOrm"
	global "github.com/cn-joyconn/goadmin/models/global"
	gocache "github.com/cn-joyconn/gocache"
	gologs "github.com/cn-joyconn/gologs"
	"github.com/cn-joyconn/goutils/array"
	joyarray "github.com/cn-joyconn/goutils/array"
	strtool "github.com/cn-joyconn/goutils/strtool"
)

var resouceCacheObj *gocache.Cache

type AdminResourceService struct {
}

func init() {
	resouceCacheObj = &gocache.Cache{
		Catalog:   global.AdminCatalog,
		CacheName: global.AdminCacheName,
	}
}

// 获取缓存用的键
// 返回值 缓存key
func (service *AdminResourceService) getResourcesIDCacheKey(pId int) string {
	return "resource_id_" + strconv.Itoa(pId)
}
func (service *AdminResourceService) getResourcesPermissionCacheKey(permission string) string {
	return "resource_per_" + permission
}
func (service *AdminResourceService) removeResourceCache(model *adminModel.AdminResource) {
	if model != nil {
		resouceCacheObj.Delete(service.getResourcesIDCacheKey(model.PId))
		if !strtool.IsBlank(model.PPermission) {
			resouceCacheObj.Delete(service.getResourcesPermissionCacheKey(model.PPermission))
		}
	}
}
func (service *AdminResourceService) getAllResourcesCacheKey() string {
	return "all_resource_list"
}

/**
* 删除
* @param pId  功能资源id
* @param pAppid  功能模块id
* @return  删除结果
 */
func (service *AdminResourceService) DeleteByPrimaryKey(pId int) int64 {
	obj := service.SelectByPrimaryKey(pId)
	if obj == nil {
		return -1
	}
	result := defaultOrm.DB.Where("PId = ?", pId).Delete(&obj)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
		return 0
	}
	if result.RowsAffected > 0 {
		roleResourceService :=&AdminRoleResourceService{}
		roleResourceService.DeleteByResourceID(pId)
		service.removeResourceCache(obj)
	}
	return result.RowsAffected

}

/**
* 插入一条功能资源信息
* @param record 功能资源实例对象
* @return 插入结果
 */
func (service *AdminResourceService) Insert(record *adminModel.AdminResource) int {
	result := defaultOrm.DB.Model(&adminModel.AdminResource{}).Create(record)
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
* 根据 功能资源id 查询 一条功能资源信息
* @param pId  功能资源id
* @param pAppid  功能模块id
* @return 查询结果
 */
func (service *AdminResourceService) SelectByPrimaryKeys(pIds []int) *[]*adminModel.AdminResource {
	if pIds == nil {
		return nil
	}
	cacheKeyList := make([]string, 0)
	notExisitIDs := make([]int, 0)
	var err error

	result := make([]*adminModel.AdminResource, 0)
	pIds = joyarray.RemoveDuplicateInt(pIds)
	if pIds != nil {
		for _, pid := range pIds {
			cacheKeyList = append(cacheKeyList, service.getResourcesIDCacheKey(pid))
		}
		if len(cacheKeyList) > 0 {
			var cachedModel *adminModel.AdminResource
			for _, key := range cacheKeyList {
				err = resouceCacheObj.Get(key, &cachedModel)
				if err == nil {
					result = append(result, cachedModel)
				}
			}
		}
		for _, pid := range pIds {
			exisit := false
			for _, resourceObj := range result {
				if resourceObj != nil && pid == resourceObj.PId {
					exisit = true
					break
				}
			}
			if !exisit {
				notExisitIDs = append(notExisitIDs, pid)
			}
		}

		if notExisitIDs != nil && len(notExisitIDs) > 0 {
			var resourceObjs []adminModel.AdminResource
			err:=defaultOrm.DB.Find(&resourceObjs, notExisitIDs).Error
			if err == nil {
				for _, resourceObj := range resourceObjs {
					cacheKey := service.getResourcesIDCacheKey((&resourceObj).PId)
						resouceCacheObj.Put(cacheKey, &resourceObj, 1000*60*60*24)
						result = append(result, &resourceObj)

				}

			}
		}
	}
	return &result

}

/**
* 根据 功能资源id 查询 一条功能资源信息
* @param pPermission  功能资源标识
* @param pAppid  功能模块id
* @return 查询结果
 */
func (service *AdminResourceService) SelectByPrimaryKey(pId int) *adminModel.AdminResource {
	cacheKey := service.getResourcesIDCacheKey(pId)
	var result *adminModel.AdminResource
	err := resouceCacheObj.Get(cacheKey, &result)
	if err != nil || result == nil {
		defaultOrm.DB.Where("PId = ?", pId).First(&result)
		if result != nil {
			resouceCacheObj.Put(cacheKey, result, 1000*60*60*24)
		}
	}

	return result
}

/**
* 根据 功能资源id 查询 一条功能资源信息
* @param pPermission  功能资源标识
* @param pAppid  功能模块id
* @return 查询结果
 */
func (service *AdminResourceService) SelectBypPermission(pPermission string) *adminModel.AdminResource {
	cacheKey := service.getResourcesPermissionCacheKey(pPermission)
	var result *adminModel.AdminResource
	err := resouceCacheObj.Get(cacheKey, &result)
	if err != nil || result == nil {
		defaultOrm.DB.Where("PPermission = ?", pPermission).First(&result)
		if result != nil {
			resouceCacheObj.Put(cacheKey, result, 1000*60*60*24)
		}
	}

	return result
}

/**
* 根据 功能资源id 查询 一条功能资源信息
* @param pPermission  功能资源标识
* @param pAppid  功能模块id
* @return 查询结果
 */
func (service *AdminResourceService) SelectBypPermissions(pPermissions []string) *[]*adminModel.AdminResource {
	if pPermissions == nil {
		return nil
	}
	cacheKeyList := make([]string, 0)
	notExisitPermissions := make([]string, 0)
	var err error

	result := make([]*adminModel.AdminResource, 0)
	pPermissions = joyarray.RemoveDuplicateStr(pPermissions)
	if pPermissions != nil {
		for _, permission := range pPermissions {
			cacheKeyList = append(cacheKeyList, service.getResourcesPermissionCacheKey(permission))
		}
		if len(cacheKeyList) > 0 {
			var cachedModel *adminModel.AdminResource
			for _, key := range cacheKeyList {
				err = resouceCacheObj.Get(key, &cachedModel)
				if err == nil {
					result = append(result, cachedModel)
				}
			}
		}
		for _, permission := range pPermissions {
			exisit := false
			for _, resourceObj := range result {
				if resourceObj != nil && permission == resourceObj.PPermission {
					exisit = true
					break
				}
			}
			if !exisit {
				notExisitPermissions = append(notExisitPermissions, permission)
			}
		}

		if notExisitPermissions != nil && len(notExisitPermissions) > 0 {
			var resourceObjs []*adminModel.AdminResource
			defaultOrm.DB.Where("PPermission in (?)", notExisitPermissions).Find(&resourceObjs)
			if resourceObjs != nil {
				for _, resourceObj := range resourceObjs {
					if resourceObj != nil {
						cacheKey := service.getResourcesPermissionCacheKey(resourceObj.PPermission)
						resouceCacheObj.Put(cacheKey, resourceObj, 1000*60*60*24)
						result = append(result, resourceObj)
					}

				}

			}
		}
	}
	return &result
}

/**
* 查询 所有资源信息
* @return 查询结果
 */
func (service *AdminResourceService) SelectAll() *[]*adminModel.AdminResource {
	cacheKey := service.getAllResourcesCacheKey()
	var result []*adminModel.AdminResource
	err := resouceCacheObj.Get(cacheKey, &result)
	if err != nil || result == nil {
		defaultOrm.DB.Find(&result)
		if result != nil {
			resouceCacheObj.Put(cacheKey, result, 1000*60*60*24*30)
		}
	}
	return &result
}

/**
* 根据 功能资源id 查询 一条功能资源信息
* @param pAppid  功能模块id
* @return 查询结果
 */
func (service *AdminResourceService) ResetAllListCache() {
	cacheKey := service.getAllResourcesCacheKey()
	resouceCacheObj.Delete(cacheKey)
}

/**
* 修改一条功能资源信息
* @param record 功能资源实例对象
* @return 修改结果
 */
func (service *AdminResourceService) UpdateByPrimaryKey(record *adminModel.AdminResource) int64 {
	result := defaultOrm.DB.Model(&record).Select("PName", "PSort", "PType", "PDesc", "PPermission").Updates(record)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
		return 0
	}
	if result.RowsAffected > 0 {
		service.removeResourceCache(record)
	}
	return result.RowsAffected
}

/**
* 获取资源的权限名称（全称）
*
* @param resource
* @return
 */
func (service *AdminResourceService) GetPermissinsNames(resources []string) [][]string {
	var result [][]string
	result = make([][]string, 0)
	var namePath []string
	var resourceModel *adminModel.AdminResource
	for _, resource := range resources {
		resourceModel = service.SelectBypPermission(resource)
		namePath = make([]string, 0)
		if resourceModel != nil {
			namePath = append(namePath, resourceModel.PName)
			for {
				resourceModel = service.SelectByPrimaryKey(resourceModel.PPid)
				if resourceModel != nil {
					namePath = append(namePath, resourceModel.PName)
				}
				if resourceModel.PId <= 0 {
					namePath = array.RemoveDuplicateStr(namePath)
					result = append(result, namePath)
					break
				}
			}
		}
	}

	return result
}
