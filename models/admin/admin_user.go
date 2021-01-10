package admin

import (
	// "database/sql"
	strtool "github.com/cn-joyconn/goutils/strtool"
	defaultOrm "github.com/cn-joyconn/goadmin/models/defaultOrm"
	gologs "github.com/cn-joyconn/gologs"
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



//Insert 插入一条
func (adminUser *AdminUser)Insert() int {
	// if len(adminUser.Password)>0{
	// 	adminUser.Password
	// }
	toMd5 := adminUser.Phone+adminUser.Email+adminUser.Username
	if !strtool.IsBlank(adminUser.Phone) {
		adminUser.PhoneMd5 = strtool.Md5(adminUser.Phone)
	}else{
		adminUser.PhoneMd5 = strtool.Md5(toMd5)
	}
	if !strtool.IsBlank(adminUser.Email) {
		adminUser.EmailMD5 = strtool.Md5(adminUser.Email)
	}else{
		adminUser.EmailMD5 = strtool.Md5(toMd5)
	}
	if !strtool.IsBlank(adminUser.Username) {
		adminUser.UsernameMd5 = strtool.Md5(adminUser.Username)
	}else{
		adminUser.UsernameMd5 = strtool.Md5(toMd5)
	}
	result :=defaultOrm.DB.Model(&AdminUser{}).Create(adminUser)
	// err := defaultOrm.DB.Model(&AdminUser{}).Where("f_id = @id",  sql.Named("id", adminUser.ID)).FirstOrCreate(&adminUser)
	if result.Error!=nil{
		gologs.GetLogger("orm").Error(result.Error.Error())
	}
	if adminUser.ID>0 {
		return adminUser.ID
	} else {
		return 0
	}
}

//UpdateInfo 修改基本信息
func (adminUser *AdminUser)UpdateInfo() int {
	result  := defaultOrm.DB.Model(&adminUser).Select("Alias", "Sex", "HeadPortrait", "UserCD", "RealName", "Description", "Remarks").Updates( adminUser)
	if result.Error!=nil{
		gologs.GetLogger("orm").Error(result.Error.Error())
		return 0
	}
	return int(result.RowsAffected)
}

// //UpdateState 修改一个用户的认证状态
// //  userID 用户ID
// //  state 状态
// //
// func (adminUser *AdminUser)UpdateState(userID uint32, state uint8) int64 {
// 	adminUser := AdminUser{}
// 	adminUser.ID = userID
// 	adminUser.Status = state
// 	id, err := defaultOrm.Orm.Update(&adminUser, "Status")
// 	if err == nil {
// 		return id
// 	} else {
// 		return 0
// 	}
// }

// //UpdatePassword 修改用户的认证密码
// //  userID 用户ID
// //  password 密码
// //
// func (adminUser *AdminUser)UpdatePassword(userID uint32, password string) int64 {
// 	adminUser := AdminUser{}
// 	adminUser.ID = userID
// 	adminUser.Password = password
// 	id, err := defaultOrm.Orm.Update(&adminUser, "Password")
// 	if err == nil {
// 		return id
// 	} else {
// 		return 0
// 	}
// }

// //UpdateUserName 修改用户名
// //  userID  用户ID
// //  username 新用户名
// //返回修改结果
// func (adminUser *AdminUser)UpdateUserName(userID uint32, username string) int64 {
// 	if strtool.IsBlank(username){
// 		return 0
// 	}
// 	adminUser := AdminUser{}
// 	adminUser.ID = userID
// 	adminUser.Username = username
// 	adminUser.UsernameMd5 = strtool.Md5(username)
// 	id, err := defaultOrm.Orm.Update(&adminUser, "UserName", "UsernameMd5")
// 	if err == nil {
// 		return id
// 	} else {
// 		return 0
// 	}
// }

// //UpdateEmail 修改邮箱
// //  userID  用户ID
// //  email 新邮箱
// //返回修改结果
// func (adminUser *AdminUser)UpdateEmail(userID uint32, email string) int64 {
// 	if strtool.IsBlank(email){
// 		return 0
// 	}
// 	adminUser := AdminUser{}
// 	adminUser.ID = userID
// 	adminUser.Email = email
// 	adminUser.EmailMD5 = strtool.Md5(email)
// 	id, err := defaultOrm.Orm.Update(&adminUser, "Email", "EmailMD5")
// 	if err == nil {
// 		return id
// 	} else {
// 		return 0
// 	}
// }

// //UpdatePhone  修改手机号
// //  userID  用户ID
// //  phone 新手机号
// //返回修改结果
// func (adminUser *AdminUser)UpdatePhone(userID uint32, phone string) int64 {
// 	if strtool.IsBlank(phone){
// 		return 0
// 	}
// 	adminUser := AdminUser{}
// 	adminUser.ID = userID
// 	adminUser.Phone = phone
// 	adminUser.PhoneMd5 = strtool.Md5(phone)
// 	id, err := defaultOrm.Orm.Update(&adminUser, "Phone", "PhoneMd5")
// 	if err == nil {
// 		return id
// 	} else {
// 		return 0
// 	}
// }
// //UpdateLoginValue  修改登录账号
// //  userID  用户ID
// //  phone 手机号
// //  email 邮箱
// //  username 用户名
// //返回修改结果
// func (adminUser *AdminUser)UpdateLoginValue(userID uint32, phone string, email string, username string) int64 {

// 	adminUser := AdminUser{}
// 	adminUser.ID = userID
// 	adminUser.Username = username
// 	adminUser.Phone = phone
// 	adminUser.Email = email
// 	toMd5 := adminUser.Phone+adminUser.Email+adminUser.Username
// 	if !strtool.IsBlank(adminUser.Phone) {
// 		adminUser.PhoneMd5 = strtool.Md5(adminUser.Phone)
// 	}else{
// 		adminUser.PhoneMd5 = strtool.Md5(toMd5)
// 	}
// 	if !strtool.IsBlank(adminUser.Email) {
// 		adminUser.EmailMD5 = strtool.Md5(adminUser.Email)
// 	}else{
// 		adminUser.EmailMD5 = strtool.Md5(toMd5)
// 	}
// 	if !strtool.IsBlank(adminUser.Username) {
// 		adminUser.UsernameMd5 = strtool.Md5(adminUser.Username)
// 	}else{
// 		adminUser.UsernameMd5 = strtool.Md5(toMd5)
// 	}
// 	id, err := defaultOrm.Orm.Update(&adminUser, "UserName", "UsernameMd5", "Phone", "PhoneMd5", "Email", "EmailMD5")
// 	if err == nil {
// 		return id
// 	} else {
// 		return 0
// 	}
// }

// //deleteByUserID 删除一个用户的认证信息
// //  userID 用户ID
// func (adminUser *AdminUser)DeleteByUserID(userID uint32) int64 {
// 	adminUser := AdminUser{}
// 	adminUser.ID = userID
// 	id, err := defaultOrm.Orm.Delete(&adminUser)

// 	if err == nil {
// 		return id
// 	} else {
// 		return 0
// 	}
// }

// /**
// * 通过用户ID获取一个用户认证信息
// *  pUserID  用户ID
// *  用户认证信息 没有结果到时返回null
//  */
// func (adminUser *AdminUser)SelectByUserID(userID uint32)(adminUser *AdminUser) {
// 	qs := defaultOrm.Orm.QueryTable(new(AdminUser))
// 	qs.Filter("ID", userID)
// 	qs.One(&adminUser)
// 	return nil
// }

// /**
// * 通过用户名获取一个用户认证信息
// *  pUsername  用户名
// *  用户认证信息 没有结果到时返回null
//  */
// func (adminUser *AdminUser)SelectByUserName(username string)(adminUser *AdminUser) {
// 	qs := defaultOrm.Orm.QueryTable(new(AdminUser))
// 	qs.Filter("UserName", username)
// 	qs.One(&adminUser)
// 	return nil
// }

// /**
// * 通过手机号获取一个用户认证信息
// *  pPhone  手机号
// *  用户认证信息 没有结果到时返回null
//  */
// func (adminUser *AdminUser)SelectByPhone(phone string)(adminUser *AdminUser) {
// 	qs := defaultOrm.Orm.QueryTable(new(AdminUser))
// 	qs.Filter("Phone", phone)
// 	qs.One(&adminUser)
// 	return nil
// }

// /**
// * 通过email获取一个用户认证信息
// *  pEmail  email
// *  用户认证信息 没有结果到时返回null
//  */
// func (adminUser *AdminUser)SelectByEmail(email string) (adminUser *AdminUser) {
// 	qs := defaultOrm.Orm.QueryTable(new(AdminUser))
// 	qs.Filter("Email", email)
// 	qs.One(&adminUser)
// 	return nil

// }
