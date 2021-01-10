package admin

func (adminRoleResource *AdminRoleResource) TableName() string {
	return "admin_role_resource"
}

//AdminUser实体类
type AdminRoleResource struct {
	PId       uint32 `orm:"column(f_id);auto";description:"自增ID";`
	PRoleid   uint32 `orm:"column(f_role_id)";description:"角色ID";json:"pRoleid"`
	PResource uint32 `orm:"column(f_resource_id)";description:"功能资源ID";json:"pResource"`
}
