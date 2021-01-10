package admin

func (adminResource *AdminResource) TableName() string {
	return "admin_resource"
}

//AdminUser实体类
type AdminResource struct {
	PId         uint32 `orm:"column(f_id);auto";description:"功能ID";`
	PName       string `orm:"size(50);column(f_name);unique";description:"名称";json:"pName"`
	PPid        uint32 `orm:"column(f_pid)";description:"父级ID";json:"pPid"`
	PLevel      uint32 `orm:"column(f_level)";description:"层级";json:"pLevel"`
	PSort       uint32 `orm:"column(f_sort)";description:"排序";json:"pSort"`
	PType       uint8  `orm:"default(1);column(f_type);";description:"资源类型(0页面内功能   1页面  2分类 )";json:"pType"`
	PDesc       string `orm:"size(60);column(f_desc);";description:"描述";json:"pDesc"`
	PPermission string `orm:"size(200);column(f_permission);";description:"功能对应的权限标识";json:"pPermission"`
}
