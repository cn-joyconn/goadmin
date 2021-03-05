package admin

import (
	"time"
)

// func (adminUser *AdminUser) TableName() string {
// 	return "admin_user"
// }

//AdminUser 实体类
type AdminUser struct {
	AdminUserBasic
	Password    string    `gorm:"type:varchar(30);column:f_password;comment:密码";json:"pPassword`
	Status      uint8     `gorm:"default:1;column:f_status;comment:状态1=正常0=禁用";json:"pStatus`
	Username    string    `gorm:"type:varchar(60);column:f_user_name;comment:用户名";json:"pUsername`
	UsernameMd5 string    `gorm:"type:char(32);column:f_user_name_md5;unique;comment:用户名md5";json:"pUsernameMd5`
	Phone       string    `gorm:"type:varchar(20);column:f_phone;comment:手机号码";json:"pPhone`
	PhoneMd5    string    `gorm:"type:char(32);column:f_phone_md5;unique;comment:手机号码md5";json:"pPhoneMd5`
	Email       string    `gorm:"type:varchar(200);column:f_email;comment:邮箱";json:"pEmail`
	EmailMD5    string    `gorm:"type:char(32);column:f_email_md5;unique;comment:邮箱md5";json:"pEmailMd5`
	UpdatedAt   time.Time `gorm:"autoUpdateTime;type:datetime;column:f_update_time;comment:更新时间";json:"pUpdatedAt`
	RealName    string    `gorm:"type:varchar(20);column:f_real_name;comment:真实姓名";json:"pRealName`
	Description string    `gorm:"type:varchar(60);column:f_description;comment:描述";json:"pDescription`
	Remarks     string    `gorm:"type:varchar(100);column:f_remarks;comment:备注";json:"pRemarks`
	PRoles      string    `gorm:"type:varchar(5000);column:f_roles;comment:角色列表,<角色、过期时间>数组的序列化字符串";json:"pRoles`
}











