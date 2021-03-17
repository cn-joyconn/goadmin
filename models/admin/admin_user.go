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
	// ID           int       `json:"pUserID"          gorm:"primaryKey;autoIncrement;column:f_id;comment:用户ID"`
	// Alias        string    `json:"pAlias"           gorm:"type:varchar(30);column:f_alias;comment:昵称"`
	// Sex          int       `json:"pSex"             gorm:"default:0;column:f_sex;comment:性别0（未知） 1（男） 2（女）"`
	// HeadPortrait string    `json:"pHeadPortrait"    gorm:"type:varchar(200);column:f_head_portrait;comment:用户头像"`
	// CreatedAt    time.Time `json:"pCreatedAt"       gorm:"autoCreateTime;type:datetime;column:f_created_time;comment:创建时间"`
	// UserCD       string    `json:"pUserCD"          gorm:"type:varchar(20);column:f_user_cd;comment:自定义编号"`
	Password     string    `json:"pPassword"        gorm:"type:char(32);column:f_password;comment:密码"`
	Status       int       `json:"pStatus"          gorm:"default:1;column:f_status;comment:状态1（正常）  0（禁用）"`
	Username     string    `json:"pUsername"        gorm:"type:varchar(60);column:f_user_name;comment:用户名"`
	UsernameMd5  string    `json:"pUsernameMd5"     gorm:"type:char(32);column:f_user_name_md5;unique;comment:用户名md5"`
	Phone        string    `json:"pPhone"           gorm:"type:varchar(20);column:f_phone;comment:手机号码"`
	PhoneMd5     string    `json:"pPhoneMd5"        gorm:"type:char(32);column:f_phone_md5;unique;comment:手机号码md5"`
	Email        string    `json:"pEmail"           gorm:"type:varchar(200);column:f_email;comment:邮箱"`
	EmailMD5     string    `json:"pEmailMd5"        gorm:"type:char(32);column:f_email_md5;unique;comment:邮箱md5"`
	UpdatedAt    time.Time `json:"pUpdatedAt"       gorm:"autoUpdateTime;type:datetime;column:f_update_time;comment:更新时间"`
	RealName     string    `json:"pRealName"        gorm:"type:varchar(20);column:f_real_name;comment:真实姓名"`
	Description  string    `json:"pDescription"     gorm:"type:varchar(60);column:f_description;comment:描述"`
	Remarks      string    `json:"pRemarks"         gorm:"type:varchar(100);column:f_remarks;comment:备注"`
	PRoles       string    `json:"pRoles"           gorm:"type:varchar(5000);column:f_roles;comment:角色列表,《角色、过期时间》数组的序列化字符串"`
}
