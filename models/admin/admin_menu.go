package admin

// func (adminMenu *AdminMenu) TableName() string {
// 	return "admin_menu"
// }

//AdminUser实体类
type AdminMenu struct {
	PId          int          `gorm:"column:f_id;primaryKey;autoIncrement;description:菜单ID";`
	PMenuID      int          `gorm:"index:idx_menu,priority:1;column:f_menu_id;description:菜单ID,根节点此值为0";json:"pMenuID"`
	PPid         int          `gorm:"column:f_pid;description:父级ID";json:"pPid"`
	PState       int          `gorm:"index:idx_menu,priority:2;default:1;column:f_state;description:状态1正常 0禁用";json:"pState"`
	PCreatuserid int          `gorm:"index;column:f_creat_user_id;description:添加人（用户id）";json:"pCreatuserid"`
	PName        string       `gorm:"type:varchar(50);column:f_name;description:名称";json:"pName"`
	PDesc        string       `gorm:"type:varchar(60);column:f_desc;description:描述";json:"pDesc"`
	PURL         string       `gorm:"type:varchar(200);column:f_url;description:请求url";json:"pUrl"`
	PIcon        string       `gorm:"type:varchar(30);column:f_icon;description:图标";json:"pIcon"`
	PPermission  string       `gorm:"type:varchar(200);column:f_permission;description:功能对应的权限标识";json:"pPermission"`
	PType        int          `gorm:"default:1;column:f_type;description:资源类型(1菜单  2页面)";json:"pType"`
	PSort        int          `gorm:"column:f_sort;description:排序";json:"pSort"`
	PLevel       int          `gorm:"column:f_level;description:层级";json:"pLevel"`
	PParams      string       `gorm:"type:varchar(100);column:f_params;description:自定义参数";json:"pParams"`
	Children     []*AdminMenu `gorm:"-";json:"children"`
}
