package admin

// func (adminResource *AdminResource) TableName() string {
// 	return "admin_resource"
// }

//AdminUser实体类
type AdminResource struct {
	PId         int              `gorm:"column:f_id;primaryKey;autoIncrement;description:功能ID";json:"pId"`
	PName       string           `gorm:"type:varchar(50);column:f_name;unique;description:名称";json:"pName"`
	PPid        int              `gorm:"column:f_pid;description:父级ID";json:"pPid"`
	PLevel      uint32           `gorm:"column:f_level;description:层级";json:"pLevel"`
	PSort       uint32           `gorm:"column:f_sort;description:排序";json:"pSort"`
	PType       uint8            `gorm:"default:1;column:f_type;description:资源类型(0页面内功能   1页面  2分类 )";json:"pType"`
	PDesc       string           `gorm:"type:varchar(60);column:f_desc;description:描述";json:"pDesc"`
	PPermission string           `gorm:"type:varchar(200);column:f_permission;description:功能对应的权限标识";json:"pPermission"`
	Children    []*AdminResource `gorm:"-";json:"children"`
}
