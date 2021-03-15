package controllers

import (
	"image/color"
	"strconv"

	"github.com/cn-joyconn/goadmin/models/global"
	"github.com/cn-joyconn/goadmin/utils/joyCaptcha"
	"github.com/cn-joyconn/goadmin/utils/saveFile"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

//通用管理
type CommonController struct {
	BaseController
}

//图片验证码
func (controller *CommonController) AuthImage(c *gin.Context) {
	HeightStr := c.DefaultQuery("height", "30")
	WidthStr := c.DefaultQuery("Width", "60")
	height, _ := strconv.Atoi(HeightStr)
	width, _ := strconv.Atoi(WidthStr)
	param := &joyCaptcha.CaptchaParam{
		CaptchaType: "string", //audio、string、math、chinese、digit

		Length: 4, //Length random string length.

		//audio
		Language: "en", // Language possible values for lang are "en", "ja", "ru", "zh".

		//string
		Height:          height,                                 // Height png height in pixel.
		Width:           width,                                  // Width Captcha png width in pixel.
		NoiseCount:      5,                                      //NoiseCount text noise count.
		ShowLineOptions: 2,                                      // OptionShowHollowLine=2 | OptionShowSlimeLine=4 | OptionShowSineLine=8 .
		Source:          base64Captcha.TxtSimpleCharaters,       //Source is a unicode which is the rand string from.
		BgColor:         &color.RGBA{R: 16, G: 16, B: 16, A: 1}, //BgColor captcha image background color (optional)
		// Fonts      :     []string    //Fonts loads by name see fonts.go's comment

		// Digit
		MaxSkew:  0.6, // MaxSkew max absolute skew factor of a single digit.
		DotCount: 8,   // DotCount Number of background circles.
	}
	id, b64s, err := joyCaptcha.GenerateCaptcha(param)
	if err != nil {
		controller.ApiError(c, "GenerateCaptchaHandler error", err)
		return
	}
	controller.ApiSuccess(c, id, b64s)
}

// 文件上传
func (controller *CommonController) Upload(c *gin.Context) {

	file, _ := c.FormFile("file")
	newFilePath, returnUrl := saveFile.GetSaveFilePath(file, global.AppConf.Upload)
	err := c.SaveUploadedFile(file, newFilePath)
	if err == nil {
		controller.ApiSuccess(c, "上传成功", returnUrl)
	} else {
		controller.ApiSuccess(c, "上传失败", "")
	}
}

// 文件上传
func (controller *CommonController) UploadFiles(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	if files != nil && len(files) > 0 {
		result := make([]string, len(files))
		for _, file := range files {
			newFilePath, returnUrl := saveFile.GetSaveFilePath(file, global.AppConf.Upload)
			c.SaveUploadedFile(file, newFilePath)
			result = append(result, returnUrl)
		}
		controller.ApiSuccess(c, "上传成功", result)
	} else {
		controller.ApiSuccess(c, "上传失败", "")
	}

}
