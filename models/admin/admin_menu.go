package admin

// func (adminMenu *AdminMenu) TableName() string {
// 	return "admin_menu"
// }

//AdminUser实体类
type AdminMenu struct {
	PId          int         `json:"pId"              gorm:"column:f_id;primaryKey;autoIncrement;description:菜单ID"`
	PMenuID      int         `json:"pMenuID"          gorm:"index:idx_menu,priority:1;column:f_menu_id;description:菜单ID,根节点此值为0"`
	PPid         int         `json:"pPid"             gorm:"column:f_pid;description:父级ID"`
	PState       int         `json:"pState"           gorm:"index:idx_menu,priority:2;default:1;column:f_state;description:状态1正常 0禁用"`
	PCreatuserid Juint64     `json:"pCreatuserid"     gorm:"index;type:bigint;column:f_creat_user_id;description:添加人（用户id）"`
	PName        string      `json:"pName"            gorm:"type:varchar(50);column:f_name;description:名称"`
	PDesc        string      `json:"pDesc"            gorm:"type:varchar(60);column:f_desc;description:描述"`
	PURL         string      `json:"pUrl"             gorm:"type:varchar(200);column:f_url;description:请求url"`
	PIcon        string      `json:"pIcon"            gorm:"type:varchar(30);column:f_icon;description:图标"`
	PPermission  string      `json:"pPermission"      gorm:"type:varchar(200);column:f_permission;description:功能对应的权限标识"`
	PType        int         `json:"pType"            gorm:"default:1;column:f_type;description:资源类型(1菜单  2页面)"`
	PSort        int         `json:"pSort"            gorm:"column:f_sort;description:排序"`
	PLevel       int         `json:"pLevel"           gorm:"column:f_level;description:层级"`
	PParams      string      `json:"pParams"          gorm:"type:varchar(100);column:f_params;description:自定义参数"`
	Children     []AdminMenu `json:"children"         gorm:"-"`
}

type AdminMenus struct {
	Arr []AdminMenu
	By  func(p, q *AdminMenu) bool
}

func (pw AdminMenus) Len() int { // 重写 Len() 方法
	return len(pw.Arr)
}
func (pw AdminMenus) Swap(i, j int) { // 重写 Swap() 方法
	pw.Arr[i], pw.Arr[j] = pw.Arr[j], pw.Arr[i]
}
func (pw AdminMenus) Less(i, j int) bool { // 重写 Less() 方法
	return pw.By(&pw.Arr[i], &pw.Arr[j])
}
