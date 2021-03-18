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
	strtool "github.com/cn-joyconn/goutils/strtool"
)

var roleCacheObj *gocache.Cache

type AdminRoleService struct {
}

func init() {
	roleCacheObj = &gocache.Cache{
		Catalog:   global.AdminCatalog,
		CacheName: global.AdminCacheName,
	}
}

// 获取缓存用的键
// 返回值 缓存key
func (service *AdminRoleService) getRoleCacheKey(pId int) string {
	return "roles_" + strconv.Itoa(pId)
}
func (service *AdminRoleService) removeRoleCache(pId int) {
	roleCacheObj.Delete(service.getRoleCacheKey(pId))
}

/**
*修改角色状态角色
* @param pId 角色id
* @param pState 状态 1正常 0禁用
* @return 结果 1:成功 小于1:失败
 */
func (service *AdminRoleService) updateStateByPrimaryKey(pId int, pState uint8) int64 {
	record := &adminModel.AdminRole{}
	record.PId = pId
	record.PState = pState
	result := defaultOrm.DB.Model(&record).Select("PState").Updates(record)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
		return 0
	}
	if result.RowsAffected > 0 {
		service.removeRoleCache(pId)
	}
	return result.RowsAffected
}

//Insert 添加角色
//结果 1:成功 小于1:失败
func (service *AdminRoleService) Insert(record *adminModel.AdminRole) int {
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
*查询角色
* @param pId 角色id
* @return  未找到时返回null
 */
func (service *AdminRoleService) SelectByPrimaryKey(pId int) *adminModel.AdminRole {
	cacheKey := service.getRoleCacheKey(pId)
	var result *adminModel.AdminRole
	err := roleCacheObj.Get(cacheKey, &result)
	if err != nil || result == nil {
		var data adminModel.AdminRole
		err:=defaultOrm.DB.Where("PId = ?", pId).First(&data).Error
		if err == nil {
			roleCacheObj.Put(cacheKey, &data, 1000*60*60*24)
			result = &data
		}

	}

	return result

}

/**
*查询角色
* @param pId 角色id
* @return  未找到时返回null
 */
func (service *AdminRoleService) SelectByRoleIds(pIds []int) *[]*adminModel.AdminRole {
	if pIds == nil {
		return nil
	}
	cacheKeyList := make([]string, 0)
	notExisitPids := make([]int, 0)
	var err error

	result := make([]*adminModel.AdminRole, 0)
	pIds = joyarray.RemoveDuplicateInt(pIds)
	if pIds != nil {
		for _, pid := range pIds {
			cacheKeyList = append(cacheKeyList, service.getRoleCacheKey(pid))
		}
		if len(cacheKeyList) > 0 {
			var cachedModel *adminModel.AdminRole
			for _, key := range cacheKeyList {
				err = roleCacheObj.Get(key, &cachedModel)
				if err == nil {
					result = append(result, cachedModel)
				}
			}
		}
		for _, pid := range pIds {
			exisit := false
			for _, roleObj := range result {
				if roleObj != nil && pid == roleObj.PId {
					exisit = true
					break
				}
			}
			if !exisit {
				notExisitPids = append(notExisitPids, pid)
			}
		}

		if notExisitPids != nil && len(notExisitPids) > 0 {
			var roleObjs []adminModel.AdminRole
			err =defaultOrm.DB.Where("PId in (?)", notExisitPids).Find(&roleObjs).Error
			if err == nil {
				for _, roleObj := range roleObjs {
					cacheKey := service.getRoleCacheKey((&roleObj).PId)
						roleCacheObj.Put(cacheKey, &roleObj, 1000*60*60*24)
						result = append(result, &roleObj)

				}

			}
		}
	}
	return &result
}

/**
*查询角色
* @param creatUser 创建用户
* @return  未找到时返回null
 */
func (service *AdminRoleService) SelectByPage(creatUser string, pageIndex int, pageSize int)(err error, list interface{}, total int64) {
	var result []adminModel.AdminRole
	db := defaultOrm.DB
	if !strtool.IsBlank(creatUser) {
		db = db.Where("PCreatuserid = ?", creatUser)
	}
	err=db.Model(&adminModel.AdminRole{}).Count(&total).Error
	if err==nil{
		err=db.Order("PId desc").Limit(pageIndex).Offset((pageIndex - 1) * pageSize).Find(&result).Error

	}
	return err,result, total
}

/**
*修改角色
* @param record 角色实例
* @return 结果 1:成功 小于1:失败
 */
func (service *AdminRoleService) UpdateByPrimaryKey(record *adminModel.AdminRole) int64 {
	result := defaultOrm.DB.Model(&record).Select("PName", "PDesc").Updates(record)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
		return 0
	}
	if result.RowsAffected > 0 {
		service.removeRoleCache(record.PId)
	}
	return result.RowsAffected
}

/**
*修改角色
* @param record 角色实例
* @return 结果 1:成功 小于1:失败
 */
func (service *AdminRoleService) UpdateStateByPrimaryKey(id int, state int) int64 {
	result := defaultOrm.DB.Model(&adminModel.AdminRole{}).Where("PId = ?", id).Update("PState", state)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
		return 0
	}
	if result.RowsAffected > 0 {
		service.removeRoleCache(id)
	}
	return result.RowsAffected
}
