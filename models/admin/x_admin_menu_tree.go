package admin


//AdminUser实体类
type AdminMenuTree struct{
	AdminMenu
	Children []*AdminMenuTree 
}
