package controllers

import (
	//"net/http"

	// global "github.com/cn-joyconn/goadmin/models/global"
	"time"

	"github.com/cn-joyconn/goadmin/models/global"
	adminServices "github.com/cn-joyconn/goadmin/services/admin"
	adminModel "github.com/cn-joyconn/goadmin/models/admin"
	"github.com/cn-joyconn/goadmin/utils/joyCaptcha"
	"github.com/cn-joyconn/goutils/strtool"
	"github.com/gin-gonic/gin"
)

type AccountController struct {
	BaseController
}

//LoginPage 登录页面
// @Tags LoginPage
// @Summary 用户登录
func (controller *AccountController) LoginPage(c *gin.Context) {
	data := gin.H{
		"pageTitle": "登录",
	}
	// username, err := c.Cookie(global.AppConf.Authorize.Cookie.LoginName)
	// if err == nil {
	// 	data["username"] = ""
	// } else {
	// 	data["username"] = username
	// }
	data["joyconnVerifyCodeloginCodeenable"] = global.AppConf.Authorize.VerifyCode.Enable
	data["ranPath"] = time.Now().Unix()

	controller.ResponseHtml(c, "account/login.html", data)
}

//LoginApi 登录接口
// @Tags LoginApi
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
func (controller *AccountController) LoginApi(c *gin.Context) {
	//比对验证码
	if global.AppConf.Authorize.VerifyCode.Enable {
		captchaID := c.PostForm("CaptchaID")
		valcode := c.PostForm("valcode")
		verifyResult := joyCaptcha.CaptchaVerify(captchaID, valcode)

		if !verifyResult || strtool.IsBlank(captchaID) || strtool.IsBlank(valcode) {
			controller.ApiErrorCode(c, "验证码不正确", "", global.CehckCodeError)
			return
		}
	}
	loginID := c.PostForm("phone")
	pwd := c.PostForm("pwd")
	remember := c.PostForm("remember")
	adminUserService:=&adminServices.AdminUserService{}
	var adminUserModel *adminModel.AdminUser
	var code int
	if adminUserService.IsPhone(loginID){
		adminUserModel,code = adminUserService.LoginByPhone(loginID,pwd)
	}else if adminUserService.IsEmail(loginID){
		adminUserModel,code = adminUserService.LoginByEmail(loginID,pwd)
	}else if adminUserService.IsUserName(loginID){
		adminUserModel,code = adminUserService.LoginByUserName(loginID,pwd)
	}else{
		controller.ApiErrorCode(c, "用户不存在或密码不正确", "", global.LoginFail)
		return
	}
	if code==global.LoginSucess {
		ObjectMapper mapper=new ObjectMapper();
		result.setCode(loginResult.getCode());
		result.setResult(mapper.readValue(mapper.writeValueAsString(loginResult.getResult()), JoyConnAuthorizeAuthenticationResultModel.class));
		String token=  tokenHelper.setAuthenticationToken(String.valueOf( loginResult.getResult().getPUserID()),loginResult.getResult().getPPassword(),"/",reponse,false);
		result.getResult().setLoginToken(token);
		result.getResult().setHeaderTokenKey(tokenHelper.getLoginTokenName());
		if(rememberMe) {
            CookieUtils.addCookie(reponse, joyconnloginName, loginName, 0, "", "");
        }else {
            CookieUtils.delCookie(request,reponse, joyconnloginName);
        }
	}else{
		controller.ApiErrorCode(c, "用户不存在或密码不正确", "", global.LoginFail)
		return
	}
	
}
