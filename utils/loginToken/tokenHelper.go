package loginToken

import (
	gocache "github.com/cn-joyconn/gocache"
	log "github.com/cn-joyconn/gologs"
	"github.com/cn-joyconn/goutils/strtool"
	"github.com/gin-gonic/gin"
	// "strconv"
)

type TokenUUid interface {
	CreatID() string
}

type TokenHelper struct {
	CacheObj       *gocache.Cache
	LoginTokenName string //loginToken在cookie或header种存储的名称
	LoginTokenKey  string //loginToken加密key(aes加密)
	Multilogin     bool   //是否运行一个账号同时登录多次
	CookieDomain   string //存储在cookie中的domain

	SUCCESS       int //成功状态码
	LoginSucess   int //登陆成功状态码
	TokenFail     int //token认证失败状态码
	TokenNotExist int //token不存在状态码

	MkUUid TokenUUid //

}

// func Crea(cacheCatlog string,cacheName string,tokenConfig *TokenConfig) {
// 	tokenHelperCacheObj = &gocache.Cache{
// 		Catalog:   cacheCatlog,
// 		CacheName: cacheName,
// 	}
// 	tokenHelperConfig = tokenConfig
// }

func (th *TokenHelper) getCacheKey(uid string) string {
	return "cas_token_" + uid
}
func (th *TokenHelper) getLoginTokenCookieName() string {
	val := th.LoginTokenName // config.String("joyconn.loginToken.cookie")
	return val
}
func (th *TokenHelper) getLoginTokenKey() string {
	val := th.LoginTokenKey // config.String("joyconn.loginToken.Key")
	return val
}

/**
* 清理缓存中的cas令牌（正常用户在下次访问时会自动重新把令牌写入缓存中，长时间不用的令牌便会清楚，防止内存无效的占用过多）
* @param uid
 */
func (th *TokenHelper) ClearAuthenticationToken(uid string) {
	th.CacheObj.Delete(th.getCacheKey(uid))
}

/**
* 验证用户和令牌是否有效
* @param uid
* @param sign
* @return 1 验证成功 0 令牌不存在 -1 验证失败
 */
func (th *TokenHelper) ValidationSign(uid string, sign string) int {
	loginUidCacacheKeyheKey := th.getCacheKey(uid)
	if th.Multilogin {
		loginUidCacacheKeyheKey = loginUidCacacheKeyheKey + "_" + sign
		var cachaObj int
		err := th.CacheObj.Get(loginUidCacacheKeyheKey, &cachaObj)
		if err != nil {
			return th.TokenNotExist
		} else {
			if cachaObj == th.LoginSucess {
				return th.SUCCESS
			} else {
				return th.TokenFail
			}
		}
	} else {
		//用户的登录唯一标示写入全局缓存，用作一个用户只能登录一次
		var cachaObj string
		err := th.CacheObj.Get(loginUidCacacheKeyheKey, &cachaObj)
		if err != nil {
			return th.TokenNotExist
		} else {
			if cachaObj == sign {
				return th.SUCCESS
			} else {
				return th.TokenFail
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
	sign := th.MkUUid.CreatID()
	ltID := CreateLoginTokenID(sign, uid, pwd, th.getLoginTokenKey())
	if ltID != nil {
		token = ltID.toString()
		if strtool.IsBlank(token) {
			log.GetLogger("").Error(uid + "生成令牌失败2！")
		} else {
			if th.Multilogin {
				cacheKey := th.getCacheKey(uid) + "_" + sign
				th.CacheObj.Put(cacheKey, th.LoginSucess, 1000*60*60*24)
			} else {
				//用户的登录唯一标示写入全局缓存，用作一个用户只能登录一次
				th.CacheObj.Put(th.getCacheKey(uid), sign, 1000*60*60*24)
			}
			if byCookie {
				cookieDomain := th.CookieDomain
				if "0.0.0.0" == cookieDomain {
					cookieDomain = ""
				}
				c.SetCookie(th.getLoginTokenCookieName(), token, 0, path, cookieDomain, false, false)
			} else {
				c.Request.Header.Set("Access-Control-Expose-Headers", th.LoginTokenName)
				c.Request.Header.Set(th.LoginTokenName, token)
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
	return th.GetMyAuthToken2(c, th.LoginTokenKey)
}
func (th *TokenHelper) GetMyAuthToken2(c *gin.Context, tokenKey string) *LoginTokenID {

	token := c.Request.Header.Get(th.LoginTokenName)
	if strtool.IsBlank(token) {
		cookie, err := c.Cookie(th.LoginTokenName)
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
