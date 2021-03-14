package admin

import (
	"strconv"

	adminModel "github.com/cn-joyconn/goadmin/models/admin"
	// adminServices "github.com/cn-joyconn/goadmin/services/admin"
	global "github.com/cn-joyconn/goadmin/models/global"
	// gologs "github.com/cn-joyconn/gologs"
	// joyarray "github.com/cn-joyconn/goutils/array"
	strtool "github.com/cn-joyconn/goutils/strtool"
)




type AdminUserPermissionService struct {
}


/**
* 获取用户的权限列表
*
* @param uid
* @return
*/
func (service *AdminUserPermissionService)GetUserRoles(uid int) []*adminModel.AdminRole{
	roleids :=service.GetUserRoleIDs(uid)	
	if roleids != nil && len(roleids) > 0 {
		adminRoleService := &AdminRoleService{}
		result := adminRoleService.SelectByRoleIds(roleids)
		return result
	}else{
		return nil
	}
}

/**
* 获取用户的权限列表
*
* @param uid
* @return
*/
func (service *AdminUserPermissionService)GetUserRoleIDs(uid int) []int {
	if uid<0{
		return nil;
	}
	// if global.IsSuperAdmin(uid) {
	// 	return this.selectAllRole("", appid,1,Integer.MAX_VALUE);
	// }
	result := make([]int, 0)
	adminUserService := &AdminUserService{}
	userRolesModels := adminUserService.GetUserRolesByUid(strconv.Itoa(uid));
	if userRolesModels != nil && len(userRolesModels) > 0 {
		for  _,roleLimitTimeModel :=range userRolesModels {
			if roleLimitTimeModel!=nil && roleLimitTimeModel.IsEffectiveTime() {
				result = append(result,roleLimitTimeModel.Role )
			}
		}
	}
	return result;
}
	
/**
* 获取用户的权限列表
*
* @param uid
* @return
*/
func  (service *AdminUserPermissionService)GetUserResources(userid string) []*adminModel.AdminResource{
	if strtool.IsBlank(userid){
		return nil;
	}
	uid,err:=strconv.Atoi(userid)
	if err!=nil{
		return nil
	}
	adminResourceService :=&AdminResourceService{}
	if global.IsSuperAdmin(uid) {
		return adminResourceService.SelectAll()
	}
	var result []*adminModel.AdminResource
	roleIDs := service.GetUserRoleIDs(uid)
	if roleIDs != nil && len(roleIDs) > 0 {
		adminRoleResourceService :=&AdminRoleResourceService{}
		resourceList := make([]int, 0)
		var resourceIds []int
		for _,roleID := range roleIDs {
			if roleID!=0 {
				resourceIds = adminRoleResourceService.SelectByRoleID(roleID)
				if resourceIds!=nil{					
					for _,resourceId:=range resourceIds{
						resourceList = append(resourceList, resourceId)
					}
				}
			}
		}
		if len(resourceList)>0{
			result = adminResourceService.SelectByPrimaryKeys(resourceList)
		}
		
	}
	return result;
}
	
/**
* 获取用户的权限列表数据
*
* @param uid
* @return
*/
func  (service *AdminUserPermissionService)GetUserResourcesList(uid string) []*adminModel.AdminResource{
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
	return list;
}
	
/**
* 获取用户的权限列表数据
*
* @param uid
* @return
*/
func (service *AdminUserPermissionService)GetUserResourcesTreeForMenu(uid string) []*adminModel.AdminResource{
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
	
	menuResourcesList :=make([]*adminModel.AdminResource, 0)
	for _,adminResourceModel :=range list{
		if adminResourceModel.PType>0{
			menuResourcesList = append(menuResourcesList, adminResourceModel)
		}

	}
	// return list2Tree(result, menuResourcesList);
	return menuResourcesList
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
func   (service *AdminUserPermissionService)PathPermissin(userid string, resuorces []string) bool {
	if strtool.IsBlank(userid){
		return false;
	}
	uid,err:=strconv.Atoi(userid)
	if err!=nil{
		return false
	}
	if global.IsSuperAdmin(uid) {
		return true;
	}
	if resuorces==nil || len(resuorces)==0 {
		return false;
	}
	// JoyConnAuthenticatePermissionResourceModel joyConnAuthenticatePermissionResourceModel = null;
	roleIDs := service.GetUserRoleIDs(uid);
	if roleIDs==nil || len(roleIDs)==0 {
		return false
	}
	adminResourceService :=&AdminResourceService{}
	resourceModels := adminResourceService.SelectBypPermissions(resuorces)
	if resourceModels==nil || len(resourceModels)==0 {
		return true
	}
	adminRoleResourceService :=&AdminRoleResourceService{}
	roleResouceCacheObjs := adminRoleResourceService.SelectByRoleIDs(roleIDs)

	for _,resource:=range resourceModels{
		if resource !=nil  {
			// List<Integer> resourceIDs =null;
			for _,roleResouceCacheObj:= range roleResouceCacheObjs{
				if roleResouceCacheObj.PResource==resource.PId{
					return true;
				}
			}
		}
	}

	return false;
}


/**
	* 获取用户是否具有该路径的访问权限
	*
	* @param uid
	* @param resuorces
	* @return
	*/
func   (service *AdminUserPermissionService)HasPathPermissin(userid string, resuorces []string)[]string {
	permission := make([]string,0)
	if strtool.IsBlank(userid){
		return permission;
	}
	uid,err:=strconv.Atoi(userid)
	if err!=nil{
		return permission
	}
	if global.IsSuperAdmin(uid) {
		return resuorces;
	}
	if resuorces==nil || len(resuorces)==0 {
		return permission;
	}
	roleIDs := service.GetUserRoleIDs(uid);
	if roleIDs==nil || len(roleIDs)==0 {
		return permission
	}
	adminResourceService := &AdminResourceService{}
	resourceModels := adminResourceService.SelectBypPermissions(resuorces)
	if resourceModels==nil || len(resourceModels)==0 {
		return permission
	}
	adminRoleResourceService := &AdminRoleResourceService{}
	roleResouceCacheObjs := adminRoleResourceService.SelectByRoleIDs(roleIDs)
	

	for _,resource:=range resourceModels{
		if resource !=nil  {
			// List<Integer> resourceIDs =null;
			for _,roleResouceCacheObj:= range roleResouceCacheObjs{
				if roleResouceCacheObj.PResource==resource.PId{
					permission = append(permission, resource.PPermission)
					break;
				}
			}
		}
	}
	return permission;
}


// public static JoyConnAuthenticatePermissionResourceModel list2Tree(JoyConnAuthenticatePermissionResourceModel root, List<JoyConnAuthenticatePermissionResourceModel> ResourceList) {
// 	ResourceList.sort((a, b)-> {if(a.getPLevel().equals(b.getPLevel())){return a.getpSort()-b.getpSort();}else{return a.getPLevel()-b.getPLevel();}});
// 	for (JoyConnAuthenticatePermissionResourceModel node1 : ResourceList) {
// 		if (node1 != null) {
// 			if (node1.getPPid().equals(0)) {//一级节点
// 				root.getChildren().add(node1);
// 			} else {//非一级节点
// 				for (JoyConnAuthenticatePermissionResourceModel node2 : ResourceList) {
// 					if (node2 != null) {
// 						if (node2.getPId().equals(node1.getPPid())) {
// 							if (node2.getChildren() == null) {
// 								node2.setChildren(new ArrayList<>());
// 							}
// 							if (!node2.getChildren().contains(node1)) {
// 								node2.getChildren().add(node1);
// 							}
// 						}
// 					}
// 				}
// 			}

// 		}
// 	}
// 	return root;
// }


// public  static JoyConnAuthenticatePermissionMenuModel list2Tree(JoyConnAuthenticatePermissionMenuModel root, List<JoyConnAuthenticatePermissionMenuModel> menuList) {
// 	menuList.sort((a, b)-> a.getpSort()-b.getpSort());
// 	for (JoyConnAuthenticatePermissionMenuModel node1 : menuList) {
// 		if (node1 != null) {
// 			if (node1.getPPid().equals(root.getPId())) {//一级节点
// 				if(root.getChildren()==null){
// 					root.setChildren(new ArrayList<>());
// 				}
// 				root.getChildren().add(node1);
// 			} else {//非一级节点
// 				for (JoyConnAuthenticatePermissionMenuModel node2 : menuList) {
// 					if (node2 != null) {
// 						if (node2.getPId().equals(node1.getPPid())) {
// 							if (node2.getChildren() == null) {
// 								node2.setChildren(new ArrayList<>());
// 							}
// 							if (!node2.getChildren().contains(node1)) {
// 								node2.getChildren().add(node1);
// 							}
// 						}
// 					}
// 				}
// 			}

// 		}
// 	}
// 	return root;
// }
// func  (service *AdminUserPermissionService)removeEmptyNode(model *adminModel.AdminMenu){
// 	if model.Children !=nil && len(model.Children)>0{
// 		JoyConnAuthenticatePermissionMenuModel childModel;
// 		for(int i=0;i<model.getChildren().size();){
// 			childModel= model.getChildren().get(i);
// 			if(childModel.getpType()==1){
// 					removeEmptyNode(childModel);
// 			}
// 			if(childModel.getpType()==2||(childModel.getChildren()!=null&&childModel.getChildren().size()>0)){
// 				i++;
// 			}else{
// 				model.getChildren().remove(i);
// 			}
// 		}

// 	}
// }

