package handle
// import (
// 	cacheLib "beego-xadmin/lib/cache"
// 	"github.com/beego/beego/v2/server/web"
// 	"github.com/beego/beego/v2/core/config"
// )

// type TokenHelper struct {
// }
// const(
// 	cacheCatalog = "go-svr-mg"
// 	cacheName = "loginToken"
// )

// func (c *TokenHelper)  getCacheKey(uid string){
// 	return "cas_token_"+uid
// }
 

// /**
// * 清理缓存中的cas令牌（正常用户在下次访问时会自动重新把令牌写入缓存中，长时间不用的令牌便会清楚，防止内存无效的占用过多）
// * @param uid
// */
// func (c *TokenHelper)  ClearAuthenticationToken(uid string){
// 	cache := cacheLib.ICache{catalog:cacheCatalog,cacheName:cacheName}
// 	cache.remove(c.getCacheKey(uid));
// }
 
 
// /**
// * 验证用户和令牌是否有效
// * @param uid
// * @param sign
// * @return 1 验证成功 0 令牌不存在 -1 验证失败
// */
// func (c *TokenHelper)   validationSign( uid string,  sign string) ValidationSignResultEnum{
// 	if (joyConnAuthenticationMultiLogin) {
// 		cacheKey :=cachePreffix+uid+"_"+sign;
// 		String cacheStr = cache.get(joyConnCacheCatalog,joyConnAuthenticationCacheCacheName,cacheKey);
// 		if(ResultCode.LoginSucess.toString().equals(cacheStr)){
// 			return ValidationSignResultEnum.Success;
// 		}else{
// 			return ValidationSignResultEnum.NonExistent;
// 		}
// 	} else {
// 		//用户的登录唯一标示写入全局缓存，用作一个用户只能登录一次
// 		String tk = cache.get(joyConnCacheCatalog,joyConnAuthenticationCacheCacheName, cachePreffix+uid);
// 		if (tk == null) {
// 			return ValidationSignResultEnum.NonExistent;
// 		} else {
// 			if(tk.equals(sign)){
// 				return ValidationSignResultEnum.Success;
// 			}else{
// 				return ValidationSignResultEnum.Failed;
// 			}
// 		}
// 	}
// }
 
 
// /**
// * 设置令牌
// * @param uid
// * @param pwd
// * @param reponse
// */
// func (c *TokenHelper)   setAuthenticationToken( uid string, pwd string,ctx *context.Context , byCookie bool) string{
// 	return   setAuthenticationToken(uid,pwd,"/",reponse,byCookie);

// }
 
// /**
// * 设置令牌
// * @param uid
// * @param pwd
// * @param reponse
// */
// func  setAuthenticationToken(  uid string, pwd string, path string,ctx *context.Context , byCookie bool) string{
// 	token:="";
// 	LoginTokenID ltID = new LoginTokenID(uid, pwd,tokenKey);
// 	if (ltID != null) {
// 		token = ltID.toString();
// 		if (token == null || "".equals(token)) {
// 			LogHelper.logger().error(uid + "生成令牌失败2！");
// 		} else {
// 			String  sign = ltID.getUniqueObjectID().toString();
// 			if (joyConnAuthenticationMultiLogin) {
// 				String cacheKey=cachePreffix+uid+"_"+sign;
// 				cache.put(joyConnCacheCatalog, joyConnAuthenticationCacheCacheName, cacheKey, ResultCode.LoginSucess.toString());
// 			} else {
// 				//用户的登录唯一标示写入全局缓存，用作一个用户只能登录一次
// 				cache.put(joyConnCacheCatalog,joyConnAuthenticationCacheCacheName, cachePreffix+uid, sign);
// 			}
// 			if(byCookie){
// 				ctx.SetCookie(getLoginTokenCookieName(),token,-1)
// 			}else{
// 				reponse.setHeader("Access-Control-Expose-Headers",joyConnAuthenticationCookieLoginToken);
// 				reponse.setHeader(joyConnAuthenticationCookieLoginToken,token);
// 			}

// 		}
// 	} else {
// 		LogHelper.logger().error(uid + "生成令牌失败1！");
// 	}
// 	return  token;

// }

// /**
// * 获取令牌
// * @param request
// * @return
// */
// @Override
// public LoginTokenID getMyAuthenticationToken(ctx *context.Context){
// 	return getMyAuthenticationToken(request,tokenKey);
// }
// @Override
// public LoginTokenID getMyAuthenticationToken(ctx *context.Context,String tokenKey){
// 	String token =  request.getHeader(joyConnAuthenticationCookieLoginToken);
// 	if(token==null){
// 		Cookie cookie = CookieUtils.getCookieByName(request,joyConnAuthenticationCookieLoginToken);
// 		token =cookie==null?"":cookie.getValue();
// 	}
// 	LoginTokenID ltID =null;
// 	if (!Strings.isNullOrEmpty(token)) {
// 		try {
// 			ltID = new LoginTokenID(token,tokenKey);
// 		} catch (Exception ex) {

// 		}
// 	}
// 	return ltID;
// }
// func getMyAuthenticationID(ctx *context.Context) string{
// 	LoginTokenID  loginTokenID= getMyAuthenticationToken(request);
// 	if(loginTokenID!=null){
// 		return loginTokenID.getUid();
// 	}else{
// 		return null;
// 	}
// }
// /**
// * 删除令牌
// * @param request
// */
// func delMyAuthenticationToken(ctx *context.Context){
// 	ctx.SetCookie(getLoginTokenCookieName(),"",0)
// }


// func getLoginTokenCookieName() string {
// 	val, _ := config.String("joyconn.loginToken.cookie")
// 	return val;
// }
// func getLoginTokenKey() string{
// 	val, _ := config.String("joyconn.loginToken.Key")
// 	return val;
// }