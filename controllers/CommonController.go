package controllers

import (
	// "beego_admin/logic"
	// "beego_xadmin/models"
	// "beego_admin/utils"
	// "fmt"
	// "github.com/astaxie/beego/orm"
	// "path"
)

//通用管理
type CommonController struct {
	BaseController
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
