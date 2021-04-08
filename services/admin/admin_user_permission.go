package admin

import (
	"sort"
	// "strconv"

	adminModel "github.com/cn-joyconn/goadmin/models/admin"
	// adminServices "github.com/cn-joyconn/goadmin/services/admin"
	global "github.com/cn-joyconn/goadmin/models/global"
	// gologs "github.com/cn-joyconn/gologs"
	// joyarray "github.com/cn-joyconn/goutils/array"
	"github.com/cn-joyconn/goutils/array"
	// strtool "github.com/cn-joyconn/goutils/strtool"
)

type AdminUserPermissionService struct {
}

/**
* 获取用户的权限列表
*
* @param uid
* @return
 */
func (service *AdminUserPermissionService) GetUserRoles(uid adminModel.Juint64) *[]*adminModel.AdminRole {
	roleids := service.GetUserRoleIDs(uid)
	if roleids != nil && len(*roleids) > 0 {
		adminRoleService := &AdminRoleService{}
		result := adminRoleService.SelectByRoleIds(*roleids)
		return result
	} else {
		return nil
	}
}

/**
* 获取用户的权限列表
*
* @param uid
* @return
 */
func (service *AdminUserPermissionService) GetUserRoleIDs(uid adminModel.Juint64) *[]int {
	if uid < 0 {
		return nil
	}
	// if global.IsSuperAdmin(uid) {
	// 	return this.selectAllRole("", appid,1,Integer.MAX_VALUE);
	// }
	result := make([]int, 0)
	adminUserService := &AdminUserService{}
	userRolesModels := adminUserService.GetUserRolesByUid(uid.ToString())
	if userRolesModels != nil && len(*userRolesModels) > 0 {
		for _, roleLimitTimeModel := range *userRolesModels {
			if roleLimitTimeModel != nil && roleLimitTimeModel.IsEffectiveTime() {
				result = append(result, roleLimitTimeModel.Role)
			}
		}
	}
	return &result
}

/**
* 获取用户的权限列表
*
* @param uid
* @return
 */
func (service *AdminUserPermissionService) GetUserResources(userid adminModel.Juint64) *[]*adminModel.AdminResource {
	// if strtool.IsBlank(userid) {
	// 	return nil
	// }
	// uid, err := strconv.Atoi(userid)
	// if err != nil {
	// 	return nil
	// }
	adminResourceService := &AdminResourceService{}
	if global.IsSuperAdmin(uint64(userid)) {
		return (adminResourceService.SelectAll())
	}
	var result []*adminModel.AdminResource
	roleIDs := service.GetUserRoleIDs(userid)
	if roleIDs != nil && len(*roleIDs) > 0 {
		adminRoleResourceService := &AdminRoleResourceService{}
		resourceList := make([]int, 0)
		var resourceIds []int
		for _, roleID := range *roleIDs {
			if roleID != 0 {
				resourceIds = *(adminRoleResourceService.SelectByRoleID(roleID))
				if resourceIds != nil {
					for _, resourceId := range resourceIds {
						resourceList = append(resourceList, resourceId)
					}
				}
			}
		}
		if len(resourceList) > 0 {
			result = *(adminResourceService.SelectByPrimaryKeys(resourceList))
		}

	}
	return &result
}

/**
* 获取用户的权限列表
*
* @param uid
* @return
 */
func (service *AdminUserPermissionService) GetUserResourceIDs(userid adminModel.Juint64) *[]int {
	// if strtool.IsBlank(userid) {
	// 	return nil
	// }
	// uid, err := strconv.Atoi(userid)
	// if err != nil {
	// 	return nil
	// }
	result := make([]int, 0)
	adminResourceService := &AdminResourceService{}
	if global.IsSuperAdmin(uint64(userid)) {
		resourceList := adminResourceService.SelectAll()
		for _, resourceObj := range *resourceList {
			if resourceObj != nil {
				result = append(result, resourceObj.PId)
			}
		}
		return &result
	}

	roleIDs := service.GetUserRoleIDs(userid)
	if roleIDs != nil && len(*roleIDs) > 0 {
		adminRoleResourceService := &AdminRoleResourceService{}
		var resourceIds []int
		for _, roleID := range *roleIDs {
			if roleID != 0 {
				resourceIds = *(adminRoleResourceService.SelectByRoleID(roleID))
				if resourceIds != nil {
					for _, resourceId := range resourceIds {
						result = append(result, resourceId)
					}
				}
			}
		}

		return &result

	}
	return nil
}
func (service *AdminUserPermissionService) GetUserPermissions(userid adminModel.Juint64) *[]string {
	models := service.GetUserResources(userid)
	result := make([]string, 0)
	if models != nil {

		for _, model := range *models {
			result = append(result, model.PPermission)
		}
	}
	return &result

}

/**
* 获取用户的权限列表数据
*
* @param uid
* @return
 */
func (service *AdminUserPermissionService) GetUserResourcesList(uid adminModel.Juint64) *[]*adminModel.AdminResource {
	// result :=&adminModel.AdminResource{
	// 	PId: 0,
	// 	PDesc: "",
	// 	PName: "权限管理",
	// 	PPid: -1,
	// 	PType: 3,
	// 	PPermission: "",
	// 	PLevel: 0,
	// 	Children:make([]*adminModel.AdminResource, 0),
	// }
	list := service.GetUserResources(uid)
	return list
}

/**
* 获取用户的权限列表数据
*
* @param uid
* @return
 */
func (service *AdminUserPermissionService) GetUserResourcesListForMenu(uid adminModel.Juint64) *[]*adminModel.AdminResource {
	// result :=&adminModel.AdminResource{
	// 	PId: 0,
	// 	PDesc: "",
	// 	PName: "权限管理",
	// 	PPid: -1,
	// 	PType: 3,
	// 	PPermission: "",
	// 	PLevel: 0,
	// 	Children:make([]*adminModel.AdminResource, 0),
	// }
	list := service.GetUserResources(uid)

	menuResourcesList := make([]*adminModel.AdminResource, 0)
	for _, adminResourceModel := range *list {
		if adminResourceModel.PType > 0 {
			menuResourcesList = append(menuResourcesList, adminResourceModel)
		}

	}
	// return list2Tree(result, menuResourcesList);
	return &menuResourcesList
}

// func  (service *AdminUserPermissionService)GetUserResourcesTree(list []*adminModel.AdminResource) *adminModel.AdminResource{
// 	result :=&adminModel.AdminResource{
// 		PId: 0,
// 		PDesc: "",
// 		PName: "权限管理",
// 		PPid: -1,
// 		PType: 3,
// 		PPermission: "",
// 		PLevel: 0,
// 		Children:make([]*adminModel.AdminResource, 0),
// 	}
// 	return list2Tree(result, list);
// }

/**
* 获取用户是否具有该路径的访问权限
*
* @param uid
* @param resuorces
* @return
 */
func (service *AdminUserPermissionService) PathPermissin(userid adminModel.Juint64, resuorces []string) bool {
	// if strtool.IsBlank(userid) {
	// 	return false
	// }
	// uid, err := strconv.Atoi(userid)
	// if err != nil {
	// 	return false
	// }
	if global.IsSuperAdmin(uint64(userid)) {
		return true
	}
	if resuorces == nil || len(resuorces) == 0 {
		return false
	}
	// JoyConnAuthenticatePermissionResourceModel joyConnAuthenticatePermissionResourceModel = null;
	roleIDs := service.GetUserRoleIDs(userid)
	if roleIDs == nil || len(*roleIDs) == 0 {
		return false
	}
	adminResourceService := &AdminResourceService{}
	resourceModels := adminResourceService.SelectBypPermissions(resuorces)
	if resourceModels == nil || len(*resourceModels) == 0 {
		return true
	}
	adminRoleResourceService := &AdminRoleResourceService{}
	roleResouceCacheObjs := *(adminRoleResourceService.SelectByRoleIDs(*roleIDs))

	for _, resource := range *resourceModels {
		if resource != nil {
			// List<Integer> resourceIDs =null;
			for _, roleResouceCacheObj := range roleResouceCacheObjs {
				if roleResouceCacheObj.PResource == resource.PId {
					return true
				}
			}
		}
	}

	return false
}

/**
* 获取用户是否具有该路径的访问权限
*
* @param uid
* @param resuorces
* @return
 */
func (service *AdminUserPermissionService) HasPathPermissin(userid adminModel.Juint64, resuorces []string) []string {
	permission := make([]string, 0)
	// if userid==nil {
	// 	return permission
	// }
	// uid, err := strconv.Atoi(userid)
	// if err != nil {
	// 	return permission
	// }
	if global.IsSuperAdmin(uint64(userid)) {
		return resuorces
	}
	if resuorces == nil || len(resuorces) == 0 {
		return permission
	}
	roleIDs := service.GetUserRoleIDs(userid)
	if roleIDs == nil || len(*roleIDs) == 0 {
		return permission
	}
	adminResourceService := &AdminResourceService{}
	resourceModels := adminResourceService.SelectBypPermissions(resuorces)
	if resourceModels == nil || len(*resourceModels) == 0 {
		return permission
	}
	adminRoleResourceService := &AdminRoleResourceService{}
	roleResouceCacheObjs := adminRoleResourceService.SelectByRoleIDs(*roleIDs)

	for _, resource := range *resourceModels {
		if resource != nil {
			// List<Integer> resourceIDs =null;
			for _, roleResouceCacheObj := range *roleResouceCacheObjs {
				if roleResouceCacheObj.PResource == resource.PId {
					permission = append(permission, resource.PPermission)
					break
				}
			}
		}
	}
	return permission
}

func (service *AdminUserPermissionService)ListResource2Tree( root adminModel.AdminResource,  ResourceList []adminModel.AdminResource) {
	//var aa = 
	sort.Sort(adminModel.AdminResources{ResourceList,func (a, b *adminModel.AdminResource) bool {
		if a.PLevel==b.PLevel{
			return a.PSort>b.PSort
		}else{
			return a.PLevel>b.PLevel
		}
    }} )
	for _, node1 :=range ResourceList {
		if node1.PPid == 0 {//一级节点
			root.Children = append(root.Children,node1);
	   } else {//非一级节点
		   for _, node2 := range ResourceList {
			if node2.PId==node1.PPid {
				if node2.Children == nil {
					node2.Children=make([]adminModel.AdminResource, 0)
				}
				Contain,err:= array.Contain(node1,node2.Children)
				if err==nil && Contain {
					node2.Children = append(node2.Children,node1);
				}
			}
		   }
	   }
	}
}

func    (service *AdminUserPermissionService)ListMenu2Tree(root adminModel.AdminMenu, menuList []adminModel.AdminMenu)  {
	sort.Sort(adminModel.AdminMenus{menuList, func (a, b *adminModel.AdminMenu) bool {
		if a.PLevel==b.PLevel{
			return a.PSort>b.PSort
		}else{
			return a.PLevel>b.PLevel
		}
    }})
	for _, node1 :=range menuList {
		if node1.PPid == root.PId {//一级节点
			root.Children = append(root.Children,node1);
	   } else {//非一级节点
		   for _, node2 := range menuList {
			if node2.PId==node1.PPid {
				if node2.Children == nil {
					node2.Children=make([]adminModel.AdminMenu, 0)
				}
				Contain,err:= array.Contain(node1,node2.Children)
				if err==nil && Contain {
					node2.Children = append(node2.Children,node1);
				}
			}
		   }
	   }
	}
	
}
func  (service *AdminUserPermissionService)removeEmptyNode(model adminModel.AdminMenu){
	if model.Children !=nil && len(model.Children)>0{
		var  childModel adminModel.AdminMenu
		var length = len(model.Children) 
		for i:=0;i<length;{
			childModel= model.Children[i];
			if childModel.PType==1{
				service.removeEmptyNode(childModel);
			}
			if childModel.PType==2||(childModel.Children!=nil&&len(childModel.Children)>0){
				i++;
			}else{
				model.Children =append(model.Children[:i], model.Children[i+1:]...) 
			}
		}

	}
}
