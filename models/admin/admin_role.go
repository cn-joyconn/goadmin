package admin

// func (adminRole *AdminRole) TableName() string {
// 	return "admin_role"
// }

//AdminUser实体类
type AdminRole struct {
	PId          int    `json:"pId"             gorm:"primaryKey;autoIncrement;column:f_id;comment:角色ID"`
	PState       uint8  `json:"pState"          gorm:"default:1;column:f_state;comment:状态1正常 0禁用"`
	PCreatuserid uint64 `json:"pCreatuserid"    gorm:"index;column:f_creat_user_id;description:添加人（用户id）"`
	PName        string `json:"pName"           gorm:"type:varchar(50);column:f_name;unique";description:角色名称"`
	PDesc        string `json:"pDesc"           gorm:"type:varchar(60);column:f_desc;description:描述"`
}
