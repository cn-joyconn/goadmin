package admin


// func (adminRoleResource *AdminRoleResource) TableName() string {
// 	return "admin_role_resource"
// }

//AdminUser实体类
type AdminRoleResource struct {
	PId       uint32 `gorm:"primaryKey;autoIncrement;column:f_id;description:自增ID;";json:"pId"`
	PRoleid   uint32 `gorm:"column:f_role_id;description:角色ID;index;";json:"pRoleid"`
	PResource uint32 `gorm:"column:f_resource_id;description:功能资源ID;index;";json:"pResource"`
}

