package admin

import (
	"time"
)

//AdminUserBasic 实体类
type AdminUserBasic struct {
	ID           Juint64   `json:"pUserID"          gorm:"primaryKey;type:bigint;column:f_id;comment:用户ID"`
	Alias        string    `json:"pAlias"           gorm:"type:varchar(30);column:f_alias;comment:昵称"`
	Sex          int       `json:"pSex"             gorm:"default:0;column:f_sex;comment:性别0（未知） 1（男） 2（女）"`
	HeadPortrait string    `json:"pHeadPortrait"    gorm:"type:varchar(200);column:f_head_portrait;comment:用户头像"`
	CreatedAt    time.Time `json:"pCreatedAt"       gorm:"autoCreateTime;type:datetime;column:f_created_time;comment:创建时间"`
	UserCD       string    `json:"pUserCD"          gorm:"type:varchar(20);column:f_user_cd;comment:自定义编号"`
}

// 查询用户信息(登录)
// searchID 查询ID
// type     查询类型 1用户id 2手机号 3邮箱 4用户名
// 返回 用户信息
// func (service *AdminUser) GetAdminUser(searchID string, searchType int) (*AdminUserBasic, error) {
// 	var result *AdminUserBasic
// 	var err error
// 	switch searchType {
// 	case 1:
// 		//userid
// 		uid, _ := strconv.Atoi(searchID)
// 		err = defaultOrm.DB.Model(&AdminUserBasic{}).First(&result, uid).Error
// 		break
// 	case 2:
// 		//phone
// 		defaultOrm.DB.Where("f_phone_md5 = ?", strtool.Md5(searchID)).First(&result)
// 		break
// 	case 3:
// 		//email
// 		defaultOrm.DB.Where("f_email_md5 = ?", strtool.Md5(searchID)).First(&result)
// 		break
// 	case 4:
// 		//username
// 		searchID = strtool.Md5(searchID)
// 		// defaultOrm.DB.Where("UsernameMd5 = ?", searchID).First(&result)
// 		defaultOrm.DB.Where("UsernameMd5 = ?", searchID).First(&result)
// 		break
// 	default:
// 		result = nil
// 		break
// 	}
// 	return result, err
// }
