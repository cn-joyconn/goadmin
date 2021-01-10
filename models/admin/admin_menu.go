package admin
func (adminMenu *AdminMenu) TableName() string {
	return "admin_menu"
}

//AdminUser实体类
type AdminMenu struct{
	PId	uint32 `orm:"column(f_id);auto";description:"菜单ID";`
	PMenuID	uint32 `orm:"column(f_menu_id)";description:"菜单ID,根节点此值为0";json:"pMenuID"`
	PPid	uint32 `orm:"column(f_pid)";description:"父级ID";json:"pPid"`
	PState uint8 `orm:"default(1);column(f_state);";description:"状态1正常 0禁用";json:"pState"`
	PCreatuserid uint64 `orm:"size(32);index;column(f_creat_user_id)";description:"添加人（用户id）";json:"pCreatuserid"`
	PName string `orm:"size(50);column(f_name);unique";description:"名称";json:"pName"`
	PDesc string `orm:"size(60);column(f_desc);";description:"描述";json:"pDesc"`
	PURL string `orm:"size(200);column(f_url);";description:"请求url";json:"pUrl"`
	PIcon string `orm:"size(30);column(f_icon);";description:"图标";json:"pIcon"`
	PPermission string `orm:"size(200);column(f_permission);";description:"功能对应的权限标识";json:"pPermission"`
	PType uint8 `orm:"default(1);column(f_type);";description:"资源类型(1菜单  2页面)";json:"pType"`
	PSort	uint32 `orm:"column(f_sort)";description:"排序";json:"pSort"`
	PLevel	uint32 `orm:"column(f_level)";description:"层级";json:"pLevel"`
	PParams string `orm:"size(100);column(f_params);";description:"自定义参数";json:"pParams"`
}
