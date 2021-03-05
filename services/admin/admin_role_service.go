package admin
import (
	// "encoding/json"
	// "crypto/md5"
	"strconv"

	defaultOrm "github.com/cn-joyconn/goadmin/models/defaultOrm"
	gologs "github.com/cn-joyconn/gologs"
	adminModel "github.com/cn-joyconn/goadmin/models/admin"
	global "github.com/cn-joyconn/goadmin/models/global"
	gocache "github.com/cn-joyconn/gocache"
	strtool "github.com/cn-joyconn/goutils/strtool"
)
var roleCacheObj *gocache.Cache
type AdminRoleService struct{

}
func init(){	
	roleCacheObj = &gocache.Cache{
		Catalog:global.AdminCatalog,
		CacheName:global.AdminCacheName,
	}	
}
// 获取缓存用的键
// 返回值 缓存key
func (service *AdminRoleService) getRoleCacheKey( pId int)string {
	return "roles_" + strconv.Itoa(pId);
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
func (service *AdminRoleService)updateStateByPrimaryKey(pId int,pState uint8)int64{
	record := &adminModel.AdminRole{}
	record.PId = pId
	record.PState = pState
	result := defaultOrm.DB.Model(&record).Select("PState").Updates(record)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
		return 0
	}
	if result.RowsAffected>0 {
		service.removeRoleCache(pId)
	}
	return result.RowsAffected
}

//Insert 添加角色
//结果 1:成功 小于1:失败
func  (service *AdminRoleService)Insert(record *adminModel.AdminRole) int{
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
func (service *AdminRoleService)SelectByPrimaryKey(pId int)*adminModel.AdminRole{
	cacheKey :=service.getRoleCacheKey(pId)
	var result *adminModel.AdminRole
	err :=roleCacheObj.Get(cacheKey,&result)
	if err!=nil || result==nil {		
		defaultOrm.DB.Where("PId = ?",pId).First(&result)
		if result!=nil{
			roleCacheObj.Put(cacheKey, result,1000*60*60*24);
		}
	}	
	
	return result

}
/**
*查询角色
* @param creatUser 创建用户
* @return  未找到时返回null
*/
func (service *AdminRoleService)SelectByPage(creatUser string,pageIndex int,pageSize int)([]*adminModel.AdminRole,int64){
	var result []*adminModel.AdminRole
	var count int64
	db := defaultOrm.DB 
	if !strtool.IsBlank(creatUser){
		db = db.Where("PCreatuserid = ?", creatUser)
	}
	db.Model(&adminModel.AdminRole{}).Count(&count)
	db.Order("PId desc").Limit(pageIndex).Offset((pageIndex-1)*pageSize).Find(&result)
	return result,count
}

/**
*修改角色
* @param record 角色实例
* @return 结果 1:成功 小于1:失败
*/
func (service *AdminRoleService)UpdateByPrimaryKey(record *adminModel.AdminRole)int64{
	result := defaultOrm.DB.Model(&record).Select("PName", "PDesc").Updates(record)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
		return 0
	}
	if result.RowsAffected>0 {
		service.removeRoleCache(record.PId)
	}
	return result.RowsAffected
}