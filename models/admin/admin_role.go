package admin
func (adminRole *AdminRole) TableName() string {
	return "admin_role"
}

//AdminUser实体类
type AdminRole struct{
	PId	uint32 `orm:"column(f_id);auto";description:"角色ID";`
	PState uint8 `orm:"default(1);column(f_state);";description:"状态1正常 0禁用";json:"pState"`
	PCreatuserid uint64 `orm:"size(32);index;column(f_creat_user_id)";description:"添加人（用户id）";json:"pCreatuserid"`
	PName string `orm:"size(50);column(f_name);unique";description:"角色名称";json:"pName"`
	PDesc string `orm:"size(60);column(f_desc);";description:"描述";json:"pDesc"`
}
