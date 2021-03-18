package admin

// func (adminRoleResource *AdminRoleResource) TableName() string {
// 	return "admin_role_resource"
// }

//AdminUser实体类
type AdminRoleResource struct {
	PId       int `json:"pId"        gorm:"primaryKey;autoIncrement;column:f_id;description:自增ID;"`
	PRoleid   int `json:"pRoleid"    gorm:"column:f_role_id;description:角色ID;index;"`
	PResource int `json:"pResource"  gorm:"column:f_resource_id;description:功能资源ID;index;"`
}
