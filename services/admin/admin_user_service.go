package admin

import (
	// "encoding/json"
	// "crypto/md5"
	"regexp"
	"strconv"
	"time"

	adminModel "github.com/cn-joyconn/goadmin/models/admin"
	global "github.com/cn-joyconn/goadmin/models/global"
	gocache "github.com/cn-joyconn/gocache"
	joyarray "github.com/cn-joyconn/goutils/array"
	encrypt "github.com/cn-joyconn/goutils/encrypt"
	strtool "github.com/cn-joyconn/goutils/strtool"
)

var cachePreffix = "user_"
var cacheObj *gocache.Cache
var adminDao *adminModel.AdminUser

func init(){	
	cacheObj = &gocache.Cache{
		Catalog:global.AdminCatalog,
		CacheName:global.AdminCacheName,
	}	
	adminDao = &adminModel.AdminUser{}
}

func  getPassword( userId string, password string)string{
	return encrypt.MakeMD5Str(userId + "\f" + password);
}

// 获取缓存用的键
// userId 用户id
// 返回值 缓存key
func getCachekey(  userId string)string {
	return cachePreffix + userId;
}


// 删除缓存
// userId 用户id
func removeCache( userId string){
	cacheObj.Delete(getCachekey(userId));
}

// 是否是用户名
// userName 用户名
// 返回值 是/否
func isUserName(userName string) bool {
	return (!isEmail(userName)) && (!isPhone(userName));
}

// 是否是邮箱
// email 邮箱
// 返回值 是/否
func  isEmail( email string) bool{
	reg1 := regexp.MustCompile(`[a-zA-Z_]{1,}[0-9]{0,}@(([a-zA-z0-9]-*){1,}\.){1,3}[a-zA-z\-]{1,}`)
    if reg1 == nil {
        //fmt.Println("regexp err")
        return false
    }
    //根据规则提取关键信息
    result1 := reg1.FindStringSubmatch(email)

	// 字符串是否与正则表达式相匹配
	return len(result1)>0;
}

// 是否是手机号
// phone 手机号
// 返回值 是/否
func  isPhone(phone string) bool {
	reg1 := regexp.MustCompile(`^1[3456789]\d{9}$`)
    if reg1 == nil {
        // fmt.Println("regexp err")
        return false
    }
    //根据规则提取关键信息
    result1 := reg1.FindStringSubmatch(phone)
    // 字符串是否与正则表达式相匹配
	return len(result1)>0;

}

// 查询用户信息(登录)
// searchID 查询ID
// type     查询类型 1用户id 2手机号 3邮箱 4用户名
// 返回 用户信息
func  GetAdminUser( searchID string,  searchType int) *adminModel.AdminUser{
	var result *adminModel.AdminUser;
	switch (searchType) {
		case 1:
			//userid			
			uid,_:=strconv.Atoi(searchID)
			result = adminDao.SelectByUserID( uid);
			break;
		case 2:
			//phone
			result = adminDao.SelectByPhone(searchID);
			break;
		case 3:
			//email
			result = adminDao.SelectByEmail(searchID);
			break;
		case 4:
			//username
			result = adminDao.SelectByUserName(searchID);
			break;
			default:
				result= nil;break;
	}
	return result;
}


// 登录验证
// userEntity 用户信息
// password    密码
// 返回 验证结果 
func  validationLogin(userEntity *adminModel.AdminUser, password string) (int){
	if userEntity==nil{

		// resultObject.setCode(ResultCode.LoginIdError);
		// resultObject.setErrorMsg("用户不存在");
		return global.LoginIdError
	}
	if userEntity.Status<1 {
		// resultObject.setCode(ResultCode.UserLocck);
		// resultObject.setErrorMsg("用户已被锁定");
		return global.LoginSucess
	}else if password != userEntity.Password {
		// resultObject.setCode(ResultCode.LoginPassError);
		// resultObject.setErrorMsg("密码错误");
		return global.LoginSucess
	}else {
		// //登录成功
		// resultObject.setResult(userEntity);
		// resultObject.setCode(ResultCode.LoginSucess);
		// resultObject.setErrorMsg("认证通过");
		return global.LoginSucess
	}
}

// 登录逻辑
// loginID  手机号\用户名\邮箱\
// password    密码
// 返回 登录结果
func  Login( loginID string, password string, loginType int,  isEncryptPwd bool) (*adminModel.AdminUser,int) {
	
	// ResultObject<JoyConnAuthorizeAuthenticationInfoModel> resultObject = new ResultObject<>();
	if strtool.IsBlank(loginID) {
		// resultObject.setCode(ResultCode.LoginIdError);
		// resultObject.setErrorMsg("用户不存在");
		return nil,global.LoginIdError
	}
	if strtool.IsBlank(password) {
		// resultObject.setCode(ResultCode.LoginPassError);
		// resultObject.setErrorMsg("密码错误");
		return nil,global.LoginPassError
	}
	authenticationInfoModel := GetAdminUser(loginID,loginType);

	if (authenticationInfoModel != nil) {
		var pwd =password;
		if(!isEncryptPwd){
			pwd = encrypt.MakeMD5Str(strconv.Itoa(authenticationInfoModel.ID)  + "\f" + password)
		}
		code  := validationLogin(authenticationInfoModel,pwd);
		return authenticationInfoModel,code
	} else {
		// resultObject.setCode(ResultCode.LoginIdError);
		// resultObject.setErrorMsg("用户不存在");
		return nil,global.LoginIdError

	}

}


/**
* 用户ID、密文密码登录
*
* @param userID
* @param password 密文密码
* @return
*/
func  LoginByUserIDAndEncryptPwd(userID string, password string)  (*adminModel.AdminUser,int){
	return Login(userID,password,1,true);
}

/**
* 用户ID、密码登录
*
* @param userID
* @param password 明文密码
* @return
*/
func  LoginByUserID(userID string,  password string)  (*adminModel.AdminUser,int){
	return Login(userID,password,1,false);

}
/**
* 用户名、密码登录
*
* @param username
* @param password  明文密码
* @return
*/
func LoginByUserName(username string, password string)  (*adminModel.AdminUser,int){
	return Login(username,password,4,false);

}
/**
* 手机号、密码登录
*
* @param phone
* @param password  明文密码
* @return
*/
func LoginByPhone(phone string, password string) (*adminModel.AdminUser,int) {
	return Login(phone,password,2,false);

}
/**
* 邮箱、密码登录
*
* @param email
* @param password  明文密码
* @return
*/
func LoginByEmail(email string, password string)  (*adminModel.AdminUser,int){
	return Login(email,password,3,false);

}

/**
* 根据用户id获取用户认证信息
* @param userID
* @return
*/
func  SelectByUserID( userID string) *adminModel.AdminUser{
	return GetAdminUser(userID,1);
}

/**
* 根据用户id获取用户认证信息
* @param phone
* @return
*/
func  SelectByPhone( phone string) *adminModel.AdminUser{
	return GetAdminUser(phone,2);
}

/**
* 根据用户id获取用户认证信息
* @param userName
* @return
*/
func  SelectByUserName(userName string) *adminModel.AdminUser{
	return GetAdminUser(userName,4);
}

/**
* 根据用户id获取用户认证信息
* @param email
* @return
*/
func  SelectByEmail(email string) *adminModel.AdminUser{
	return GetAdminUser(email,3);
}
/**
* 根据一批用户id获取一批用户基本信息
* @param userIDs
* @return
*/
func  SelectByUserCDS(userIDs []string) []*adminModel.AdminUserBasic{
	if userIDs==nil{
		return nil
	}
	cacheKeyList :=make([]string, 0)
	notExisitIDs := make([]int, 0)
	var err error
	
	result := make([]*adminModel.AdminUserBasic, 0)
	userIDs = joyarray.RemoveDuplicateStr(userIDs)
	if userIDs!=nil{
		for _,userID := range userIDs{
			cacheKeyList = append(cacheKeyList, getCachekey(userID))
		}
		if len(cacheKeyList)>0{
			// var cachedModels []*adminModel.AdminUser
			// err =  cacheObj.Get(cacheKeyList,cachedModels);
			var cachedModel *adminModel.AdminUserBasic
			for _, key :=range cacheKeyList{
				err =cacheObj.Get(key,&cachedModel)
				if err==nil{
					result = append(result,cachedModel)
				}		
			}
		}
		for _, userID := range userIDs{
			uid,_:=strconv.Atoi(userID)
			exisit := false
			for _,user := range result{
				if user!=nil && uid == user.ID{
					exisit = true
					break
				}
			}
			if !exisit{
				notExisitIDs = append(notExisitIDs,uid)
			}
		}
		
		if notExisitIDs!=nil&&len(notExisitIDs)>0{
			userObjs := adminDao.GetPubInfoByUserIDs(notExisitIDs);
			if userObjs!=nil {
				for _,userObj :=range userObjs{
					if userObj!=nil{
						cacheKey:=getCachekey(strconv.Itoa(userObj.ID))
						cacheObj.Put(cacheKey, userObj,1000*60*60*24);
						result = append(result,userObj)
					}
							
				}
				
			}		
		}
	}
	return result;
}
/**
* 根据一批用户id获取一批用户基本信息
* @param userIDs
* @return
*/
func  SelectByAdminUserList(pageSize int,pageIndex int) ([]*adminModel.AdminUserBasic,int64){
	adminUser := &adminModel.AdminUser{}
	return adminUser.GetAdminUserIDs(pageSize,pageIndex)
}
/**
* 添加用户
* @param userID 用户ID
* @param userName 登录用户户名 不填则为null
* @param userPhone 登录手机号 不填则为null
* @param userEmail 登录邮箱 不填则为null
* @param pwd 密码
* @param state 状态
* @return
*/
func InsertAuthenticationUserModel(userID string,userName string, userPhone string,userEmail string,pwd string,state uint8) (int,*adminModel.AdminUser){
	adminUser := &adminModel.AdminUser{}
	// adminUser.ID=userObj.ID
	adminUser.Alias = userName
	adminUser.Username = userName
	adminUser.Phone = userPhone
	adminUser.Email = userEmail
	adminUser.Password=adminUser.GetSaltPwd(pwd)
	adminUser.Status=state
	uid := adminUser.Insert()
	if uid > 0 {
		adminUser.ID = uid
		return 1,adminUser
	}else{
		return -1,nil
	}
	
}
/**
* 添加用户
* @param userID 用户ID
* @param userName 登录用户户名 不填则为null
* @param userPhone 登录手机号 不填则为null
* @param userEmail 登录邮箱 不填则为null
* @param pwd 密码
* @param state 状态
* @return
*/
func InsertUserModel(adminUser *adminModel.AdminUser) (int,*adminModel.AdminUser){
	uid := adminUser.Insert()
	if uid > 0 {
		adminUser.ID = uid
		return 1,adminUser
	}else{
		return -1,nil
	}
	
}
func UpdateUserPubInfo(uid int, Alias string, Sex uint8, HeadPortrait string, CreatedAt  time.Time, UserCD  string) (int){

	obj :=new(adminModel.AdminUser)
	obj =obj.SelectByUserID(uid)
	if obj==nil{
		return -1
	}
	obj.Alias=Alias
	obj.Sex=Sex
	obj.HeadPortrait=HeadPortrait
	obj.CreatedAt=CreatedAt
	obj.UserCD=UserCD
	updateResult :=obj.UpdateInfo()
	
	if updateResult > 0 {
		cacheKey:=getCachekey(strconv.Itoa(uid))
		cacheObj.Delete(cacheKey);
		return 1
	}else{
		return -1
	}
	
}

func UpdateAuthenticationState(pUserID int,pState uint8) int64{
	obj :=new(adminModel.AdminUser)
	return obj.UpdateState(pUserID,pState)	
}

func UpdateAuthenticationUserName(pUserID int, val string) int64{
	obj :=new(adminModel.AdminUser)
	return obj.UpdateUserName(pUserID,val)	
}
func UpdateAuthenticationEmail(pUserID int, val string) int64{
	obj :=new(adminModel.AdminUser)
	return obj.UpdateEmail(pUserID,val)	
}
func UpdateAuthenticationPhone(pUserID int, val string) int64{
	obj :=new(adminModel.AdminUser)
	return obj.UpdatePhone(pUserID,val)	
}
func UpdateAuthenticationLoginValue(pUserID int,phone string, email string, username string) int64{
	obj :=new(adminModel.AdminUser)
	return obj.UpdateLoginValue(pUserID,phone,email,username)	
}
/**
* 修改用户的认证密码
* @param pUserID 用户ID
* @param pPassword 密码
* @return
*/
func UpdateAuthenticationPassword(pUserID int,pPassword string) int64{
	obj :=new(adminModel.AdminUser)
	return obj.UpdatePassword(pUserID,pPassword)	

}