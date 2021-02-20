package admin

import (
	// "database/sql"
	defaultOrm "github.com/cn-joyconn/goadmin/models/defaultOrm"
	gologs "github.com/cn-joyconn/gologs"
	strtool "github.com/cn-joyconn/goutils/strtool"

	// gorm "gorm.io/gorm"
	// array "github.com/cn-joyconn/goutils/array"
	"time"
)

// func (adminUser *AdminUser) TableName() string {
// 	return "admin_user"
// }

//AdminUser 实体类
type AdminUser struct {
	AdminUserBasic
	Password    string    `gorm:"type:varchar(30);column:f_password;comment:密码"`
	Status      uint8     `gorm:"default:1;column:f_status;comment:状态1=正常0=禁用"`
	Username    string    `gorm:"type:varchar(60);column:f_user_name;comment:用户名";`
	UsernameMd5 string    `gorm:"type:char(32);column:f_user_name_md5;unique;comment:用户名md5"`
	Phone       string    `gorm:"type:varchar(20);column:f_phone;comment:手机号码"`
	PhoneMd5    string    `gorm:"type:char(32);column:f_phone_md5;unique;comment:手机号码md5"`
	Email       string    `gorm:"type:varchar(200);column:f_email;comment:邮箱"`
	EmailMD5    string    `gorm:"type:char(32);column:f_email_md5;unique;comment:邮箱md5"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime;type:datetime;column:f_update_time;comment:更新时间"`
	RealName    string    `gorm:"type:varchar(20);column:f_real_name;comment:真实姓名"`
	Description string    `gorm:"type:varchar(60);column:f_description;comment:描述"`
	Remarks     string    `gorm:"type:varchar(100);column:f_remarks;comment:备注"`
	PRoles      string    `gorm:"type:varchar(5000);column:f_roles;comment:角色列表,<角色、过期时间>数组的序列化字符串"`
}

var pwdsalt = "\fjoyadmin"

//GetSaltPwd 密码加盐
func (adminUser *AdminUser) GetSaltPwd(password string) string {
	return strtool.Md5(password + pwdsalt)
}

//Insert 插入一条
func (adminUser *AdminUser) Insert() int {
	// if len(adminUser.Password)>0{
	// 	adminUser.Password
	// }
	toMd5 := adminUser.Phone + adminUser.Email + adminUser.Username
	if !strtool.IsBlank(adminUser.Phone) {
		adminUser.PhoneMd5 = strtool.Md5(adminUser.Phone)
	} else {
		adminUser.PhoneMd5 = strtool.Md5(toMd5)
	}
	if !strtool.IsBlank(adminUser.Email) {
		adminUser.EmailMD5 = strtool.Md5(adminUser.Email)
	} else {
		adminUser.EmailMD5 = strtool.Md5(toMd5)
	}
	if !strtool.IsBlank(adminUser.Username) {
		adminUser.UsernameMd5 = strtool.Md5(adminUser.Username)
	} else {
		adminUser.UsernameMd5 = strtool.Md5(toMd5)
	}
	result := defaultOrm.DB.Model(&AdminUser{}).Create(adminUser)
	// err := defaultOrm.DB.Model(&AdminUser{}).Where("f_id = @id",  sql.Named("id", adminUser.ID)).FirstOrCreate(&adminUser)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
	}
	if adminUser.ID > 0 {
		return adminUser.ID
	} else {
		return 0
	}
}

//UpdateInfo 修改基本信息
func (adminUser *AdminUser) UpdateInfo() int64 {
	result := defaultOrm.DB.Model(&adminUser).Select("Alias", "Sex", "HeadPortrait", "UserCD", "RealName", "Description", "Remarks").Updates(adminUser)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
		return 0
	}
	return result.RowsAffected
}

//UpdateState 修改一个用户的认证状态
//  userID 用户ID
//  state 状态
//
func (adminUser *AdminUser) UpdateState(userID int, state uint8) int64 {
	updateObj := &AdminUser{}
	updateObj.ID = userID
	updateObj.Status = state
	result := defaultOrm.DB.Model(&adminUser).Select("Status").Updates(updateObj)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
		return 0
	}
	return result.RowsAffected
}

//UpdatePassword 修改用户的认证密码
//  userID 用户ID
//  password 密码
//
func (adminUser *AdminUser) UpdatePassword(userID int, password string) int64 {
	updateObj := &AdminUser{}
	updateObj.ID = userID
	updateObj.Password = password
	result := defaultOrm.DB.Model(&adminUser).Select("Password").Updates(updateObj)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
		return 0
	}
	return result.RowsAffected
}

//UpdateUserName 修改用户名
//  userID  用户ID
//  username 新用户名
//返回修改结果
func (adminUser *AdminUser) UpdateUserName(userID int, username string) int64 {
	if strtool.IsBlank(username) {
		return 0
	}
	updateObj := &AdminUser{}
	updateObj.ID = userID
	updateObj.Username = username
	updateObj.UsernameMd5 = strtool.Md5(username)
	result := defaultOrm.DB.Model(&adminUser).Select("Username", "UsernameMd5").Updates(updateObj)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
		return 0
	}
	return result.RowsAffected
}

//UpdateEmail 修改邮箱
//  userID  用户ID
//  email 新邮箱
//返回修改结果
func (adminUser *AdminUser) UpdateEmail(userID int, email string) int64 {
	if strtool.IsBlank(email) {
		return 0
	}
	updateObj := &AdminUser{}
	updateObj.ID = userID
	updateObj.Email = email
	updateObj.EmailMD5 = strtool.Md5(email)
	result := defaultOrm.DB.Model(&adminUser).Select("Email", "EmailMD5").Updates(updateObj)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
		return 0
	}
	return result.RowsAffected
}

//UpdatePhone  修改手机号
//  userID  用户ID
//  phone 新手机号
//返回修改结果
func (adminUser *AdminUser) UpdatePhone(userID int, phone string) int64 {
	if strtool.IsBlank(phone) {
		return 0
	}
	updateObj := &AdminUser{}
	updateObj.ID = userID
	updateObj.Phone = phone
	updateObj.PhoneMd5 = strtool.Md5(phone)
	result := defaultOrm.DB.Model(&adminUser).Select("Phone", "PhoneMd5").Updates(updateObj)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
		return 0
	}
	return result.RowsAffected
}

//UpdateLoginValue  修改登录账号
//  userID  用户ID
//  phone 手机号
//  email 邮箱
//  username 用户名
//返回修改结果
func (adminUser *AdminUser) UpdateLoginValue(userID int, phone string, email string, username string) int64 {

	updateObj := &AdminUser{}
	updateObj.ID = userID
	updateObj.Username = username
	updateObj.Phone = phone
	updateObj.Email = email
	toMd5 := adminUser.Phone + adminUser.Email + adminUser.Username
	if !strtool.IsBlank(adminUser.Phone) {
		updateObj.PhoneMd5 = strtool.Md5(adminUser.Phone)
	} else {
		updateObj.PhoneMd5 = strtool.Md5(toMd5)
	}
	if !strtool.IsBlank(adminUser.Email) {
		updateObj.EmailMD5 = strtool.Md5(adminUser.Email)
	} else {
		updateObj.EmailMD5 = strtool.Md5(toMd5)
	}
	if !strtool.IsBlank(adminUser.Username) {
		updateObj.UsernameMd5 = strtool.Md5(adminUser.Username)
	} else {
		updateObj.UsernameMd5 = strtool.Md5(toMd5)
	}
	result := defaultOrm.DB.Model(&adminUser).Select("UserName", "UsernameMd5", "Phone", "PhoneMd5", "Email", "EmailMD5").Updates(updateObj)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
		return 0
	}
	return result.RowsAffected

}

//deleteByUserID 删除一个用户的认证信息
//  userID 用户ID
func (adminUser *AdminUser) DeleteByUserID(userID int) int64 {
	updateObj := &AdminUser{}
	updateObj.ID = userID
	result := defaultOrm.DB.Delete(&updateObj)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
		return 0
	}
	return result.RowsAffected
}

/**
* 通过用户ID获取一个用户认证信息
*  pUserID  用户ID
*  用户认证信息 没有结果到时返回null
 */
func (adminUser *AdminUser) SelectByUserID(userID int) *AdminUser {

	var result *AdminUser
	defaultOrm.DB.First(&result, userID)
	return result
}

/**
* 通过用户ID获取一个用户公共信息
*  pUserID  用户ID
*  用户认证信息 没有结果到时返回null
 */
func (adminUser *AdminUser) GetPubInfoByUserID(userID int) *AdminUserBasic {

	var result *AdminUserBasic
	defaultOrm.DB.Select("ID","Alias","Sex","HeadPortrait","CreatedAt","UserCD").First(&result, userID)
	return result
}
/**
* 通过用户ID获取一个用户公共信息
*  pUserID  用户ID
*  用户认证信息 没有结果到时返回null
 */
 func (adminUser *AdminUser) GetPubInfoByUserIDs(userID []int) []*AdminUserBasic {
	var result []*AdminUserBasic
	defaultOrm.DB.Select("ID","Alias","Sex","HeadPortrait","CreatedAt","UserCD").Find(&result, userID)
	return result
}
/**
* 获取用户列表
 */
 func (adminUser *AdminUser) GetAdminUserIDs(pageSize int,pageIndex int)( []*AdminUserBasic,int64 ){
	var result []*AdminUserBasic
	var count int64
	defaultOrm.DB.Model(&AdminUserBasic{}).Count(&count)
	defaultOrm.DB.Order("ID desc").Limit(pageIndex).Offset((pageIndex-1)*pageSize).Find(&result)
	return result,count
}
/**
* 通过用户名获取一个用户认证信息
*  pUsername  用户名
*  用户认证信息 没有结果到时返回null
 */
func (adminUser *AdminUser) SelectByUserName(username string) *AdminUser {
	var result *AdminUser
	defaultOrm.DB.Where("f_user_name_md5 = ?", strtool.Md5(username)).First(&result)
	return result
}

/**
* 通过手机号获取一个用户认证信息
*  pPhone  手机号
*  用户认证信息 没有结果到时返回null
 */
func (adminUser *AdminUser) SelectByPhone(phone string) *AdminUser {
	var result *AdminUser
	defaultOrm.DB.Where("f_phone_md5 = ?", strtool.Md5(phone)).First(&result)
	return result
}

/**
* 通过email获取一个用户认证信息
*  pEmail  email
*  用户认证信息 没有结果到时返回null
 */
func (adminUser *AdminUser) SelectByEmail(email string) *AdminUser {
	var result *AdminUser
	defaultOrm.DB.Where("f_email_md5 = ?", strtool.Md5(email)).First(&result)
	return result

}
