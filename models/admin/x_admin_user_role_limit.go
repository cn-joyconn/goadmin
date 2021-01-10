package admin


//AdminUser实体类
type XAdminUserRoleLimit struct{
	AdminUser
	PRoleObjs []*XAdminRoleLimit 
}
