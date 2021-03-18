package admin

import (
	"time"

	adminModel "github.com/cn-joyconn/goadmin/models/admin"
	defaultOrm "github.com/cn-joyconn/goadmin/models/defaultOrm"
	gologs "github.com/cn-joyconn/gologs"
)

type AdminLogService struct {
}

//添加日志
func (service *AdminLogService) Insert(record *adminModel.AdminLog) {
	result := defaultOrm.DB.Model(&record).Create(record)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
	}
}

//查询日志
func (service *AdminLogService) SelectByPage(time time.Time, pageIndex int, pageSize int) (*[]adminModel.AdminLog, int64,error) {
	var result []adminModel.AdminLog
	var count int64
	db := defaultOrm.DB.Where("CreatedAt < ?", time)
	err:=db.Model(&adminModel.AdminLog{}).Count(&count).Error
	if err==nil{

		err=db.Order("Id desc").Limit(pageIndex).Offset((pageIndex - 1) * pageSize).Find(&result).Error
	}
	return &result, count,err
}

//删除日志
func (service *AdminLogService) DeleteByPID(pId []int) int64 {
	result := defaultOrm.DB.Where("Id in ? ", pId).Delete(&adminModel.AdminLog{})
	return result.RowsAffected
}
