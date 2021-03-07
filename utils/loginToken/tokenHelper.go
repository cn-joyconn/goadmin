package loginToken

import (
	global "github.com/cn-joyconn/goadmin/models/global"
	gocache "github.com/cn-joyconn/gocache"
	log "github.com/cn-joyconn/gologs"
	"github.com/cn-joyconn/goutils/strtool"
	"github.com/gin-gonic/gin"
)

var tokenHelperCacheObj *gocache.Cache

type TokenHelper struct {
}

func init() {
	tokenHelperCacheObj = &gocache.Cache{
		Catalog:   global.AdminCatalog,
		CacheName: global.AdminCacheName,
	}
}

func (th *TokenHelper) getCacheKey(uid string) string {
	return "cas_token_" + uid
}
func (th *TokenHelper) getLoginTokenCookieName() string {
	val := global.AppConf.Authorize.Cookie.LoginToken // config.String("joyconn.loginToken.cookie")
	return val
}
func (th *TokenHelper) getLoginTokenKey() string {
	val := global.AppConf.Authorize.Cookie.LoginTokenAesKey // config.String("joyconn.loginToken.Key")
	return val
}

/**
* 清理缓存中的cas令牌（正常用户在下次访问时会自动重新把令牌写入缓存中，长时间不用的令牌便会清楚，防止内存无效的占用过多）
* @param uid
 */
func (th *TokenHelper) ClearAuthenticationToken(uid string) {
	tokenHelperCacheObj.Delete(th.getCacheKey(uid))
}

/**
* 验证用户和令牌是否有效
* @param uid
* @param sign
* @return 1 验证成功 0 令牌不存在 -1 验证失败
 */
func (th *TokenHelper) ValidationSign(uid string, sign string) int {
	loginUidCacacheKeyheKey := th.getCacheKey(uid)
	if global.AppConf.Authorize.Multilogin {
		loginUidCacacheKeyheKey = loginUidCacacheKeyheKey + "_" + sign
		var cachaObj int
		err := tokenHelperCacheObj.Get(loginUidCacacheKeyheKey, &cachaObj)
		if err != nil {
			return global.TokenNotExist
		} else {
			if cachaObj == global.LoginSucess {
				return global.SUCCESS
			} else {
				return global.TokenFail
			}
		}
	} else {
		//用户的登录唯一标示写入全局缓存，用作一个用户只能登录一次
		var cachaObj string
		err := tokenHelperCacheObj.Get(loginUidCacacheKeyheKey, &cachaObj)
		if err != nil {
			return global.TokenNotExist
		} else {
			if cachaObj == sign {
				return global.SUCCESS
			} else {
				return global.TokenFail
			}
		}
	}

}

/**
* 设置令牌
* @param uid
* @param pwd
* @param reponse
 */
func (th *TokenHelper) SetAuthenticationToken(uid string, pwd string, c *gin.Context, byCookie bool) string {
	return th.SetAuthenticationToken2(uid, pwd, "/", c, byCookie)

}

/**
* 设置令牌
* @param uid
* @param pwd
* @param reponse
 */
func (th *TokenHelper) SetAuthenticationToken2(uid string, pwd string, path string, c *gin.Context, byCookie bool) string {
	token := ""
	ltID := CreateLoginTokenID(uid, pwd, th.getLoginTokenKey())
	if ltID != nil {
		token = ltID.toString()
		if strtool.IsBlank(token) {
			log.GetLogger("").Error(uid + "生成令牌失败2！")
		} else {
			sign := ltID.Sign
			if global.AppConf.Authorize.Multilogin {
				cacheKey := th.getCacheKey(uid) + "_" + sign
				tokenHelperCacheObj.Put(cacheKey, global.LoginSucess, 1000*60*60*24)
			} else {
				//用户的登录唯一标示写入全局缓存，用作一个用户只能登录一次
				tokenHelperCacheObj.Put(th.getCacheKey(uid), sign, 1000*60*60*24)
			}
			if byCookie {
				cookieDomain := global.AppConf.Authorize.Cookie.Domain
				if "0.0.0.0" == cookieDomain {
					cookieDomain = ""
				}
				c.SetCookie(th.getLoginTokenCookieName(), token, 0, path, cookieDomain, false, false)
			} else {
				c.Request.Header.Set("Access-Control-Expose-Headers", global.AppConf.Authorize.Cookie.LoginToken)
				c.Request.Header.Set(global.AppConf.Authorize.Cookie.LoginToken, token)
			}

		}
	} else {
		log.GetLogger("").Error(uid + "生成令牌失败1！")
	}
	return token

}

/**
* 获取令牌
* @param request
* @return
 */
func (th *TokenHelper) GetMyAuthToken(c *gin.Context) *LoginTokenID {
	return th.GetMyAuthToken2(c, global.AppConf.Authorize.Cookie.LoginTokenAesKey)
}
func (th *TokenHelper) GetMyAuthToken2(c *gin.Context, tokenKey string) *LoginTokenID {

	token := c.Request.Header.Get(global.AppConf.Authorize.Cookie.LoginToken)
	if strtool.IsBlank(token) {
		cookie, err := c.Cookie(global.AppConf.Authorize.Cookie.LoginToken)
		if err != nil {
			token = ""
		} else {
			token = cookie
		}
	}
	var ltID *LoginTokenID
	if !strtool.IsBlank(token) {
		ltID = ParseLoginTokenID(token, tokenKey)
	}
	return ltID
}
func (th *TokenHelper) GetMyAuthenticationID(c *gin.Context) string {
	loginTokenID := th.GetMyAuthToken(c)
	if loginTokenID != nil {
		return loginTokenID.Uid
	} else {
		return ""
	}
}

/**
* 删除令牌
* @param request
 */
func (th *TokenHelper) DelMyAuthenticationToken(c *gin.Context) {
	c.SetCookie(th.getLoginTokenCookieName(), "", -1, "", "", false, false)
}
