package controllers

import (
	//"net/http"

	// global "github.com/cn-joyconn/goadmin/models/global"
	"time"

	"github.com/cn-joyconn/goadmin/models/global"
	"github.com/cn-joyconn/goadmin/utils/joyCaptcha"
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
	username, err := c.Cookie(global.AppConf.Authorize.Cookie.LoginName)
	if err == nil {
		data["username"] = ""
	} else {
		data["username"] = username
	}
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

		if !verifyResult {
			controller.ApiErrorCode(c, "验证码不正确", "", global.CehckCodeError)
			return
		}
	}

	// var L request.Login
	// _ = c.ShouldBindJSON(&L)
	// if err := utils.Verify(L, utils.LoginVerify); err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// }
	// if store.Verify(L.CaptchaId, L.Captcha, true) {
	// 	U := &model.SysUser{Username: L.Username, Password: L.Password}
	// 	if err, user := service.Login(U); err != nil {
	// 		global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误", zap.Any("err", err))
	// 		response.FailWithMessage("用户名不存在或者密码错误", c)
	// 	} else {
	// 		tokenNext(c, *user)
	// 	}
	// } else {
	// 	response.FailWithMessage("验证码错误", c)
	// }
}
