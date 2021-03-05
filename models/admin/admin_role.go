package admin



// func (adminRole *AdminRole) TableName() string {
// 	return "admin_role"
// }

//AdminUser实体类
type AdminRole struct{
	PId	int `gorm:"primaryKey;autoIncrement;column:f_id;comment:角色ID";json:"pId"`
	PState uint8 `gorm:"default:1;column:f_state;comment:状态1正常 0禁用";json:"pState"`
	PCreatuserid uint64 `gorm:"index;column:f_creat_user_id;description:添加人（用户id）";json:"pCreatuserid"`
	PName string `gorm:"type:varchar(50);column:f_name;unique";description:"角色名称";json:"pName"`
	PDesc string `gorm:"type:varchar(60);column:f_desc;description:描述;json:"pDesc"`
}
