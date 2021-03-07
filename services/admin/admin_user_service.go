package admin

import (
	// "encoding/json"
	// "crypto/md5"
	"regexp"
	"strconv"

	// "time"

	adminModel "github.com/cn-joyconn/goadmin/models/admin"
	defaultOrm "github.com/cn-joyconn/goadmin/models/defaultOrm"
	global "github.com/cn-joyconn/goadmin/models/global"
	gocache "github.com/cn-joyconn/gocache"
	gologs "github.com/cn-joyconn/gologs"
	joyarray "github.com/cn-joyconn/goutils/array"
	encrypt "github.com/cn-joyconn/goutils/encrypt"
	strtool "github.com/cn-joyconn/goutils/strtool"
)

var pwdsalt = "\fjoyadmin"
var userCacheObj *gocache.Cache

type AdminUserService struct {
}

func init() {
	userCacheObj = &gocache.Cache{
		Catalog:   global.AdminCatalog,
		CacheName: global.AdminCacheName,
	}
}

//GetSaltPwd 密码加盐
func (service *AdminUserService) GetSaltPwd(password string) string {
	return strtool.Md5(password + pwdsalt)
}

// 获取缓存用的键
// userId 用户id
// 返回值 缓存key
func (service *AdminUserService) getUserCachekey(userId string) string {
	return "user_" + userId
}

// 删除缓存
// userId 用户id
func (service *AdminUserService) removeUserCache(userId string) {
	userCacheObj.Delete(service.getUserCachekey(userId))
}

// 是否是用户名
// userName 用户名
// 返回值 是/否
func (service *AdminUserService) IsUserName(userName string) bool {
	return (!service.IsEmail(userName)) && (!service.IsPhone(userName))
}

// 是否是邮箱
// email 邮箱
// 返回值 是/否
func (service *AdminUserService) IsEmail(email string) bool {
	reg1 := regexp.MustCompile(`[a-zA-Z_]{1,}[0-9]{0,}@(([a-zA-z0-9]-*){1,}\.){1,3}[a-zA-z\-]{1,}`)
	if reg1 == nil {
		//fmt.Println("regexp err")
		return false
	}
	//根据规则提取关键信息
	result1 := reg1.FindStringSubmatch(email)

	// 字符串是否与正则表达式相匹配
	return len(result1) > 0
}

// 是否是手机号
// phone 手机号
// 返回值 是/否
func (service *AdminUserService) IsPhone(phone string) bool {
	reg1 := regexp.MustCompile(`^1[3456789]\d{9}$`)
	if reg1 == nil {
		// fmt.Println("regexp err")
		return false
	}
	//根据规则提取关键信息
	result1 := reg1.FindStringSubmatch(phone)
	// 字符串是否与正则表达式相匹配
	return len(result1) > 0

}

// 查询用户信息(登录)
// searchID 查询ID
// type     查询类型 1用户id 2手机号 3邮箱 4用户名
// 返回 用户信息
func (service *AdminUserService) GetAdminUser(searchID string, searchType int) *adminModel.AdminUser {
	var result *adminModel.AdminUser
	switch searchType {
	case 1:
		//userid
		uid, _ := strconv.Atoi(searchID)
		defaultOrm.DB.First(&result, uid)
		break
	case 2:
		//phone
		defaultOrm.DB.Where("f_phone_md5 = ?", strtool.Md5(searchID)).First(&result)
		break
	case 3:
		//email
		defaultOrm.DB.Where("f_email_md5 = ?", strtool.Md5(searchID)).First(&result)
		break
	case 4:
		//username
		defaultOrm.DB.Where("f_user_name_md5 = ?", strtool.Md5(searchID)).First(&result)
		break
	default:
		result = nil
		break
	}
	return result
}

// 登录验证
// userEntity 用户信息
// password    密码
// 返回 验证结果
func (service *AdminUserService) validationLogin(userEntity *adminModel.AdminUser, password string) int {
	if userEntity == nil {

		// resultObject.setCode(ResultCode.LoginIdError);
		// resultObject.setErrorMsg("用户不存在");
		return global.LoginIdError
	}
	if userEntity.Status < 1 {
		// resultObject.setCode(ResultCode.UserLocck);
		// resultObject.setErrorMsg("用户已被锁定");
		return global.UserLocck
	} else if password != userEntity.Password {
		// resultObject.setCode(ResultCode.LoginPassError);
		// resultObject.setErrorMsg("密码错误");
		return global.LoginPassError
	} else {
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
func (service *AdminUserService) Login(loginID string, password string, loginType int, isEncryptPwd bool) (*adminModel.AdminUser, int) {

	// ResultObject<JoyConnAuthorizeAuthenticationInfoModel> resultObject = new ResultObject<>();
	if strtool.IsBlank(loginID) {
		// resultObject.setCode(ResultCode.LoginIdError);
		// resultObject.setErrorMsg("用户不存在");
		return nil, global.LoginIdError
	}
	if strtool.IsBlank(password) {
		// resultObject.setCode(ResultCode.LoginPassError);
		// resultObject.setErrorMsg("密码错误");
		return nil, global.LoginPassError
	}
	authenticationInfoModel := service.GetAdminUser(loginID, loginType)

	if authenticationInfoModel != nil {
		var pwd = password
		if !isEncryptPwd {
			pwd = encrypt.MakeMD5Str(strconv.Itoa(authenticationInfoModel.ID) + "\f" + password)
		}
		code := service.validationLogin(authenticationInfoModel, pwd)
		return authenticationInfoModel, code
	} else {
		// resultObject.setCode(ResultCode.LoginIdError);
		// resultObject.setErrorMsg("用户不存在");
		return nil, global.LoginIdError

	}

}

/**
* 用户ID、密文密码登录
*
* @param userID
* @param password 密文密码
* @return
 */
func (service *AdminUserService) LoginByUserIDAndEncryptPwd(userID string, password string) (*adminModel.AdminUser, int) {
	return service.Login(userID, password, 1, true)
}

/**
* 用户ID、密码登录
*
* @param userID
* @param password 明文密码
* @return
 */
func (service *AdminUserService) LoginByUserID(userID string, password string) (*adminModel.AdminUser, int) {
	return service.Login(userID, password, 1, false)

}

/**
* 用户名、密码登录
*
* @param username
* @param password  明文密码
* @return
 */
func (service *AdminUserService) LoginByUserName(username string, password string) (*adminModel.AdminUser, int) {
	return service.Login(username, password, 4, false)

}

/**
* 手机号、密码登录
*
* @param phone
* @param password  明文密码
* @return
 */
func (service *AdminUserService) LoginByPhone(phone string, password string) (*adminModel.AdminUser, int) {
	return service.Login(phone, password, 2, false)

}

/**
* 邮箱、密码登录
*
* @param email
* @param password  明文密码
* @return
 */
func (service *AdminUserService) LoginByEmail(email string, password string) (*adminModel.AdminUser, int) {
	return service.Login(email, password, 3, false)

}

/**
* 根据用户id获取用户认证信息
* @param userID
* @return
 */
func (service *AdminUserService) GetUserByUserID(userID string) *adminModel.AdminUser {
	return service.GetAdminUser(userID, 1)
}

/**
* 根据用户id获取用户认证信息
* @param phone
* @return
 */
func (service *AdminUserService) GetUserByPhone(phone string) *adminModel.AdminUser {
	return service.GetAdminUser(phone, 2)
}

/**
* 根据用户id获取用户认证信息
* @param userName
* @return
 */
func (service *AdminUserService) GetUserByUserName(userName string) *adminModel.AdminUser {
	return service.GetAdminUser(userName, 4)
}

/**
* 根据用户id获取用户认证信息
* @param email
* @return
 */
func (service *AdminUserService) GetUserByEmail(email string) *adminModel.AdminUser {
	return service.GetAdminUser(email, 3)
}

/**
* 根据一批用户id获取一批用户基本信息
* @param userIDs
* @return
 */
func (service *AdminUserService) GetUserInfoByUserCDS(userIDs []string) []*adminModel.AdminUserBasic {
	if userIDs == nil {
		return nil
	}
	cacheKeyList := make([]string, 0)
	notExisitIDs := make([]int, 0)
	var err error

	result := make([]*adminModel.AdminUserBasic, 0)
	userIDs = joyarray.RemoveDuplicateStr(userIDs)
	if userIDs != nil {
		for _, userID := range userIDs {
			cacheKeyList = append(cacheKeyList, service.getUserCachekey(userID))
		}
		if len(cacheKeyList) > 0 {
			// var cachedModels []*adminModel.AdminUser
			// err =  userCacheObj.Get(cacheKeyList,cachedModels);
			var cachedModel *adminModel.AdminUserBasic
			for _, key := range cacheKeyList {
				err = userCacheObj.Get(key, &cachedModel)
				if err == nil {
					result = append(result, cachedModel)
				}
			}
		}
		for _, userID := range userIDs {
			uid, _ := strconv.Atoi(userID)
			exisit := false
			for _, user := range result {
				if user != nil && uid == user.ID {
					exisit = true
					break
				}
			}
			if !exisit {
				notExisitIDs = append(notExisitIDs, uid)
			}
		}

		if notExisitIDs != nil && len(notExisitIDs) > 0 {
			var userObjs []*adminModel.AdminUserBasic
			defaultOrm.DB.Select("ID", "Alias", "Sex", "HeadPortrait", "CreatedAt", "UserCD").Find(&userObjs, notExisitIDs)
			if userObjs != nil {
				for _, userObj := range userObjs {
					if userObj != nil {
						cacheKey := service.getUserCachekey(strconv.Itoa(userObj.ID))
						userCacheObj.Put(cacheKey, userObj, 1000*60*60*24)
						result = append(result, userObj)
					}

				}

			}
		}
	}
	return result
}

/**
* 根据一批用户id获取一批用户基本信息
* @param userIDs
* @return
 */
func (service *AdminUserService) SelectUserList(pageSize int, pageIndex int) ([]*adminModel.AdminUser, int64) {
	var result []*adminModel.AdminUser
	var count int64
	defaultOrm.DB.Model(&adminModel.AdminUser{}).Count(&count)
	defaultOrm.DB.Order("ID desc").Limit(pageIndex).Offset((pageIndex - 1) * pageSize).Find(&result)
	return result, count
}

//InsertUser 添加用户
//userID 用户ID
//userName 登录用户户名 不填则为null
//userPhone 登录手机号 不填则为null
//userEmail 登录邮箱 不填则为null
//pwd 密码
//state 状态
func (service *AdminUserService) InsertUser(userID string, userName string, userPhone string, userEmail string, pwd string, state uint8) (int, *adminModel.AdminUser) {
	adminUser := &adminModel.AdminUser{}
	// adminUser.ID=userObj.ID
	adminUser.Alias = userName
	adminUser.Username = userName
	adminUser.Phone = userPhone
	adminUser.Email = userEmail
	adminUser.Password = service.GetSaltPwd(pwd)
	adminUser.Status = state
	return service.InsertUserModel(adminUser)

}

//InsertUserModel 添加用户
func (service *AdminUserService) InsertUserModel(adminUser *adminModel.AdminUser) (int, *adminModel.AdminUser) {
	toMd5 := adminUser.Phone + adminUser.Email + adminUser.Username
	if !strtool.IsBlank(adminUser.Phone) {
		adminUser.PhoneMd5 = strtool.Md5(adminUser.Phone)
	} else {
		adminUser.PhoneMd5 = strtool.Md5(toMd5)
	}
	if !strtool.IsBlank(adminUser.Email) {
		adminUser.EmailMD5 = strtool.Md5(adminUser.Email)
	} else {
		adminUser.EmailMD5 = strtool.Md5(toMd5)
	}
	if !strtool.IsBlank(adminUser.Username) {
		adminUser.UsernameMd5 = strtool.Md5(adminUser.Username)
	} else {
		adminUser.UsernameMd5 = strtool.Md5(toMd5)
	}
	result := defaultOrm.DB.Model(&adminUser).Create(adminUser)
	// err := defaultOrm.DB.Model(&AdminUser{}).Where("f_id = @id",  sql.Named("id", adminUser.ID)).FirstOrCreate(&adminUser)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
	}
	if adminUser.ID > 0 {
		return adminUser.ID, adminUser
	} else {
		return -1, nil
	}
}
func (service *AdminUserService) UpdateUserPubInfo(uid int, Alias string, Sex uint8, HeadPortrait string, UserCD string) int {

	var obj *adminModel.AdminUser
	defaultOrm.DB.First(&obj, uid)
	if obj == nil {
		return -1
	}
	obj.Alias = Alias
	obj.Sex = Sex
	obj.HeadPortrait = HeadPortrait
	obj.UserCD = UserCD
	result := defaultOrm.DB.Model(&obj).Select("Alias", "Sex", "HeadPortrait", "UserCD", "RealName", "Description", "Remarks").Updates(obj)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
		return -1
	}
	if result.RowsAffected > 0 {
		service.removeUserCache(strconv.Itoa(uid))
		return 1
	} else {
		return -1
	}

}

//UpdateState 修改一个用户的认证状态
//  userID 用户ID
//  state 状态
func (service *AdminUserService) UpdateUserState(pUserID int, pState uint8) int64 {
	updateObj := &adminModel.AdminUser{}
	updateObj.ID = pUserID
	updateObj.Status = pState
	result := defaultOrm.DB.Model(&updateObj).Select("Status").Updates(updateObj)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
		return 0
	}
	return result.RowsAffected
}

//UpdateUserName 修改用户名
//  userID  用户ID
//  username 新用户名
//返回修改结果
func (service *AdminUserService) UpdateUserName(pUserID int, username string) int64 {
	if strtool.IsBlank(username) {
		return 0
	}
	updateObj := &adminModel.AdminUser{}
	updateObj.ID = pUserID
	updateObj.Username = username
	updateObj.UsernameMd5 = strtool.Md5(username)
	result := defaultOrm.DB.Model(&updateObj).Select("Username", "UsernameMd5").Updates(updateObj)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
		return 0
	}
	return result.RowsAffected
}

//UpdateEmail 修改邮箱
//  userID  用户ID
//  email 新邮箱
//返回修改结果
func (service *AdminUserService) UpdateUserEmail(pUserID int, email string) int64 {
	if strtool.IsBlank(email) {
		return 0
	}
	updateObj := &adminModel.AdminUser{}
	updateObj.ID = pUserID
	updateObj.Email = email
	updateObj.EmailMD5 = strtool.Md5(email)
	result := defaultOrm.DB.Model(&updateObj).Select("Email", "EmailMD5").Updates(updateObj)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
		return 0
	}
	return result.RowsAffected
}

//UpdatePhone  修改手机号
//  userID  用户ID
//  phone 新手机号
//返回修改结果
func (service *AdminUserService) UpdateUserPhone(pUserID int, phone string) int64 {
	if strtool.IsBlank(phone) {
		return 0
	}
	updateObj := &adminModel.AdminUser{}
	updateObj.ID = pUserID
	updateObj.Phone = phone
	updateObj.PhoneMd5 = strtool.Md5(phone)
	result := defaultOrm.DB.Model(&updateObj).Select("Phone", "PhoneMd5").Updates(updateObj)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
		return 0
	}
	return result.RowsAffected
}

//UpdateLoginValue  修改登录账号
//  userID  用户ID
//  phone 手机号
//  email 邮箱
//  username 用户名
//返回修改结果
func (service *AdminUserService) UpdateUserLoginValue(pUserID int, phone string, email string, username string) int64 {

	updateObj := &adminModel.AdminUser{}
	updateObj.ID = pUserID
	updateObj.Username = username
	updateObj.Phone = phone
	updateObj.Email = email
	toMd5 := updateObj.Phone + updateObj.Email + updateObj.Username
	if !strtool.IsBlank(updateObj.Phone) {
		updateObj.PhoneMd5 = strtool.Md5(updateObj.Phone)
	} else {
		updateObj.PhoneMd5 = strtool.Md5(toMd5)
	}
	if !strtool.IsBlank(updateObj.Email) {
		updateObj.EmailMD5 = strtool.Md5(updateObj.Email)
	} else {
		updateObj.EmailMD5 = strtool.Md5(toMd5)
	}
	if !strtool.IsBlank(updateObj.Username) {
		updateObj.UsernameMd5 = strtool.Md5(updateObj.Username)
	} else {
		updateObj.UsernameMd5 = strtool.Md5(toMd5)
	}
	result := defaultOrm.DB.Model(&updateObj).Select("UserName", "UsernameMd5", "Phone", "PhoneMd5", "Email", "EmailMD5").Updates(updateObj)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
		return 0
	}
	return result.RowsAffected

}

//UpdatePassword 修改用户的认证密码
//  userID 用户ID
//  password 密码
func (service *AdminUserService) UpdateUserPassword(pUserID int, pPassword string) int64 {
	updateObj := &adminModel.AdminUser{}
	updateObj.ID = pUserID
	updateObj.Password = pPassword
	result := defaultOrm.DB.Model(&updateObj).Select("Password").Updates(updateObj)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
		return 0
	}
	return result.RowsAffected
}

//DeleteByUserID 删除一个用户的认证信息
//  userID 用户ID
func (service *AdminUserService) DeleteUser(userID int) int64 {
	updateObj := &adminModel.AdminUser{}
	updateObj.ID = userID
	result := defaultOrm.DB.Delete(&updateObj)
	if result.Error != nil {
		gologs.GetLogger("orm").Error(result.Error.Error())
		return 0
	}
	return result.RowsAffected
}
