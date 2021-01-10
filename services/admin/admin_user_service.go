package admin

// import (	
// 	"regexp"
	
// 	adminModel "beego_xadmin/models/admin"
// 	encrypt "beego_xadmin/lib/encrypt"
// 	icache "beego_xadmin/lib/cache"
// 	config "github.com/beego/beego/v2/core/config"
// )

// var catalog string
// var admin_cacheName string
// var cachePreffix = "user_"
// var cacheUtil icache.CacheUtil

// func init(){
// 	catalog,_=config.String("cache::xadmin.cache.catalog");
// 	admin_cacheName,_=config.String("cache::xadmin.cache.admin.cacheName");
// 	cacheUtil = icache.CacheUtil{
// 		Catalog:catalog,
// 		CacheName:admin_cacheName, 
// 	}
// }

// func  getPassword( userId string, password string)string{
// 	return encrypt.MakeMD5Str(userId + "\f" + password);
// }
// /**
//  * 获取缓存用的键
//  * @param userId 用户id
//  * @return
//  */
// func getCachekey(  userId string)string {
// 	return cachePreffix + userId;
// }

// /**
//  * 删除缓存
//  * @param userId
//  */
// func removeCache( userId string){
// 	cacheUtil.Delete(getCachekey(userId));
// }


// // 是否是用户名
// // userName 用户名
// // 返回值 是/否
// func isUserName(userName string) bool {
// 	return (!isEmail(userName)) && (!isPhone(userName));
// }


// // 是否是邮箱
// // email 邮箱
// // 返回值 是/否
// func  isEmail( email string) bool{
// 	reg1 := regexp.MustCompile(`[a-zA-Z_]{1,}[0-9]{0,}@(([a-zA-z0-9]-*){1,}\.){1,3}[a-zA-z\-]{1,}`)
//     if reg1 == nil {
//         fmt.Println("regexp err")
//         return
//     }
//     //根据规则提取关键信息
//     result1 := reg1.FindStringSubmatch(email, -1)
	
// 	// 字符串是否与正则表达式相匹配
// 	return len(result1)>0;
// }


// // 是否是手机号
// // phone 手机号
// // 返回值 是/否
// func  isPhone(phone string) bool {
// 	reg1 := regexp.MustCompile(`^1[3456789]\d{9}$`)
//     if reg1 == nil {
//         fmt.Println("regexp err")
//         return
//     }
//     //根据规则提取关键信息
//     result1 := reg1.FindStringSubmatch(phone, -1)
//     // 字符串是否与正则表达式相匹配
// 	return len(result1)>0;

// }
// /**
//  * 查询用户信息(登录)
//  *
//  * @param searchID 查询ID
//  * @param type     查询类型 1用户id 2手机号 3邮箱 4用户名
//  * @return
//  */
// func  getAdminUser( searchID string,  searchType int) adminModel.AdminUser{
// 	var result adminModel.AdminUser;
// 	switch (searchType) {
// 		case 1:
// 			//userid
// 			cacheKey:= getCachekey(searchID);
// 			err:= cacheUtil.Get(cacheKey,&result);			
// 			if (result == null) {
// 				result = tJoyConnAuthenticateAuthenticationDao.selectByUserID(searchID);
// 				if (result != null) {
// 					try{
// 						userStr := objectMapper.writeValueAsString(result);
// 						cache.put(joyconnCacheCatalog,joyconnAuthenticateCahceCacheName,  cacheKey, userStr);
// 					}catch (Exception ex){

// 					}
// 				}
// 			}

// 			break;
// 		case 2:
// 			//phone
// 			result = tJoyConnAuthenticateAuthenticationDao.selectByPhone(searchID);
// 			break;
// 		case 3:
// 			//email
// 			result = tJoyConnAuthenticateAuthenticationDao.selectByEmail(searchID);
// 			break;
// 		case 4:
// 			//username
// 			result = tJoyConnAuthenticateAuthenticationDao.selectByUserName(searchID);
// 			break;
// 			default:
// 				result=null;break;
// 	}
// 	return result;
// }


// /**
//  * 登录验证
//  * @param userEntity
//  * @param password
//  * @return
//  */
// private ResultObject<JoyConnAuthorizeAuthenticationInfoModel> validationLogin(JoyConnAuthorizeAuthenticationInfoModel userEntity, String password){
// 	ResultObject<JoyConnAuthorizeAuthenticationInfoModel> resultObject = new ResultObject<>();
// 	if(userEntity==null){
// 		resultObject.setCode(ResultCode.LoginIdError);
// 		resultObject.setErrorMsg("用户不存在");
// 	}
// 	if (userEntity.getPState()<1) {
// 		resultObject.setCode(ResultCode.UserLocck);
// 		resultObject.setErrorMsg("用户已被锁定");
// 	}else if (!password.equals(userEntity.getPPassword())) {
// 		resultObject.setCode(ResultCode.LoginPassError);
// 		resultObject.setErrorMsg("密码错误");
// 	}else {
// 		//登录成功
// 		resultObject.setResult(userEntity);
// 		resultObject.setCode(ResultCode.LoginSucess);
// 		resultObject.setErrorMsg("认证通过");
// 	}
// 	return resultObject;
// }
// /**
//  * 登录逻辑
//  *
//  * @param loginID  手机号\用户名\邮箱\
//  * @param password
//  * @return
//  */
// private ResultObject<JoyConnAuthorizeAuthenticationInfoModel> login(String loginID, String password, int loginType, boolean isEncryptPwd) {
// 	ResultObject<JoyConnAuthorizeAuthenticationInfoModel> resultObject = new ResultObject<>();
// 	if (loginID == null || "".equals(loginID)) {
// 		resultObject.setCode(ResultCode.LoginIdError);
// 		resultObject.setErrorMsg("用户不存在");
// 	}
// 	if (password == null || "".equals(password)) {
// 		resultObject.setCode(ResultCode.LoginPassError);
// 		resultObject.setErrorMsg("密码错误");
// 	}
// 	JoyConnAuthorizeAuthenticationInfoModel authenticationInfoModel = getUserAuthenticationInfo(loginID,loginType);


// 	if (authenticationInfoModel != null) {
// 		try {
// 			String pwd =password;
// 			if(!isEncryptPwd){
// 				pwd = MD5Util.md5Encode(authenticationInfoModel.getPUserID() + "\f" + password);
// 			}
// 			resultObject = validationLogin(authenticationInfoModel,pwd);
// 		}catch (Exception ex){
// 			resultObject.setCode(ResultCode.LoginPassError);
// 			resultObject.setErrorMsg("密码错误");
// 		}
// 	} else {
// 		resultObject.setCode(ResultCode.LoginIdError);
// 		resultObject.setErrorMsg("用户不存在");

// 	}
// 	return resultObject;

// }