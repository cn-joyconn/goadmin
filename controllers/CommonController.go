package controllers

import (
	"image/color"
	"strconv"

	"github.com/cn-joyconn/goadmin/utils/joyCaptcha"
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

//文件上传
// func (c *CommonController) Upload(){
// 	var files models.Files
// 	adminUser := c.GetSession("admin_user")
// 	if adminUser != nil{
// 		//后台用户
// 		files.AdminUserId = adminUser.(*models.AdminUser).Id
// 	}
// 	// TODO 这里预留未设置前台用户ID

// 	_ = c.Ctx.Input.Bind(&files.Type, "type")
// 	_ = c.Ctx.Input.Bind(&files.Remark, "remark")
// 	//保存文件
// 	f, h, err := c.GetFile("file")
// 	defer f.Close()
// 	if err != nil {
// 		c.ApiError("文件上传失败",nil)
// 	}
// 	if logic.CheckFileExt(h.Filename) == false{
// 		c.ApiError("后缀名不符合上传要求",nil)
// 	}
// 	//自动创建日期文件夹
// 	newDir := logic.AutoCreateUploadDateDir("head_images")
// 	//生成新的文件路径
// 	newFilePath := newDir + "/" + utils.GetRandomString(32) + path.Ext(h.Filename)
// 	_ = c.SaveToFile("file", newFilePath)
// 	//将图片保存到数据库
// 	files.FilePath = newFilePath
// 	o := orm.NewOrm()
// 	_, err = o.Insert(&files)
// 	if err != nil{
// 		c.ApiError("文件数据上传失败",nil)
// 	}
// 	//得到图片的完整路径
// 	fullNewFilePath := logic.GetFullPath(c.Ctx,newFilePath)
// 	fmt.Println(fullNewFilePath)
// 	c.ApiSuccess("上传成功",map[string]interface{}{"path":fullNewFilePath})
// }
