package admin

// func (adminResource *AdminResource) TableName() string {
// 	return "admin_resource"
// }

//AdminUser实体类
type AdminResource struct {
	PId         int              `json:"pId"         gorm:"column:f_id;primaryKey;autoIncrement;description:功能ID"`
	PName       string           `json:"pName"       gorm:"type:varchar(50);column:f_name;unique;description:名称"`
	PPid        int              `json:"pPid"        gorm:"column:f_pid;description:父级ID"`
	PLevel      uint32           `json:"pLevel"      gorm:"column:f_level;description:层级"`
	PSort       uint32           `json:"pSort"       gorm:"column:f_sort;description:排序"`
	PType       uint8            `json:"pType"       gorm:"default:1;column:f_type;description:资源类型(0页面内功能   1页面  2分类 )"`
	PDesc       string           `json:"pDesc"       gorm:"type:varchar(60);column:f_desc;description:描述"`
	PPermission string           `json:"pPermission" gorm:"type:varchar(200);column:f_permission;description:功能对应的权限标识"`
	Children    []AdminResource `json:"children"    gorm:"-"`
}
