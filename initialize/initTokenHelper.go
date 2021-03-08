package initialize

import (
	"strconv"

	"github.com/cn-joyconn/goadmin/models/global"
	"github.com/cn-joyconn/goadmin/utils/loginToken"
	gocache "github.com/cn-joyconn/gocache"
)


func initTokenHelper(){
	cacheObj := &gocache.Cache{
		Catalog:   global.AdminCatalog,
		CacheName: global.AdminCacheName,
	}
	mmu := &myMakeUUid{}
	global.TokenHelper = &loginToken.TokenHelper {
		CacheObj : cacheObj,
		LoginTokenName : global.AppConf.Authorize.Cookie.LoginToken, //loginToken在cookie或header种存储的名称
		LoginTokenKey  : global.AppConf.Authorize.Cookie.LoginTokenAesKey, //loginToken加密key(aes加密)
		Multilogin     : global.AppConf.Authorize.Multilogin ,  //是否运行一个账号同时登录多次
		CookieDomain   : global.AppConf.Authorize.Cookie.Domain, //存储在cookie中的domain

		SUCCESS       : global.SUCCESS, //成功状态码
		LoginSucess   : global.LoginSucess, //登陆成功状态码
		TokenFail     : global.TokenFail, //token认证失败状态码
		TokenNotExist : global.TokenNotExist, //token不存在状态码

		MkUUid : mmu ,
	}
}
type myMakeUUid struct{

} 
func ( mmu *myMakeUUid)CreatID() string {
	id:=global.SnowflakeWorker.GetId()
	return strconv.FormatInt(id, 16)
}