package admin

import "time"

//AdminUserBasic 实体类
type AdminUserBasic struct {
	ID           int       `gorm:"primaryKey;autoIncrement;column:f_id;comment:用户ID"`
	Alias        string    `gorm:"type:varchar(30);column:f_alias;comment:昵称"`
	Sex          uint8     `gorm:"default:0;column:f_sex;comment:性别0=未知1=男2=女"`
	HeadPortrait string    `gorm:"type:char(32);column:f_head_portrait;comment:用户头像"`
	CreatedAt    time.Time `gorm:"autoCreateTime;type:datetime;column:f_created_time;comment:创建时间"`
	UserCD       string    `gorm:"type:varchar(20);column:f_user_cd;comment:自定义编号"`
}

//根据用户名和密码获取单条数据
// func  (adminUser *AdminUser) AdminUserGetUserOneByNameAndPwd(username,password string) (*AdminUser, error){
// 	m := AdminUser{}
// 	err := orm.NewOrm().QueryTable(adminUser.TableName()).Filter("f_user_name",username).Filter("f_password",password).One(&m)
// 	if err != nil{
// 		return nil, err
// 	}
// 	return &m, nil
// }
