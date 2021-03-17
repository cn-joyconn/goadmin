package controllers

import "github.com/gin-gonic/gin"

// "fmt"
// "os"
// "runtime"
// "strconv"
// "time"

//后台首页管理
type IndexController struct {
	BaseController
}

func (controller *IndexController) Index(c *gin.Context) {
	data := gin.H{
		"pageTitle": "登录",
	}
	// username, err := c.Cookie(global.AppConf.Authorize.Cookie.LoginName)
	// if err == nil {
	// 	data["username"] = ""
	// } else {
	// 	data["username"] = username
	// }
	// data["joyconnVerifyCodeloginCodeenable"] = global.AppConf.Authorize.VerifyCode.Enable
	// data["ranPath"] = time.Now().Unix()

	controller.ResponseHtml(c, "index/index.html", data)
	//c.Abort()
}

// func (c *IndexController) Login() {

// 	adminUser := c.GetSession("admin_user")
// 	if adminUser != nil{
// 		// 当前已经登录
// 		c.RequestSuccess("已登录","/admin")
// 	}

// 	c.TplName = "index/login.html"
// }

// func (c *IndexController) DoLogin()  {
// 	username := c.GetString("username")
// 	password := c.GetString("password")

// 	sessionAdminUser := c.GetSession("admin_user")
// 	if sessionAdminUser != nil{
// 		//当前已登录
// 		c.ApiSuccess("用户已登录",nil)
// 	}

// 	if len(username) == 0 || len(password) == 0 {
// 		c.ApiError("用户名或者密码不能为空",nil)
// 	}

// 	//加密密码且加盐
// 	password = utils.Md5Encode(password + beego.AppConfig.String("pwd_salt"))

// 	adminUser,err := models.AdminUserGetUserOneByNameAndPwd(username,password)
// 	if err != nil{
// 		c.ApiError("用户名或者密码错误",nil)
// 	}
// 	if adminUser.Status == 2{
// 		c.ApiError("用户被禁用，请联系管理员",nil)
// 	}

// 	//保存信息到Session
// 	c.ApiSuccess("登陆成功",nil)
// }

// //退出登录
// func (c *IndexController) Logout(){
// 	c.DelSession("admin_user")
// 	c.RequestSuccess("退出成功","/admin/login")
// }

// func (c *IndexController) Welcome()  {
// 	//数据统计
// 	o := orm.NewOrm()
// 	//访问数
// 	var maps []orm.Params
// 	_, _ = o.Raw(`SELECT
// 			count(id) as sum_visit_count,
// 			count(IF(DATE_FORMAT(created_at, '%Y')=DATE_FORMAT(CURDATE(),'%Y'),true,null)) as year_visit_count,
// 			count(IF(DATE_FORMAT(created_at, '%Y%m')=DATE_FORMAT(CURDATE(),'%Y%m'),true,null)) as month_visit_count,
// 			count(IF(DATE_FORMAT(created_at, '%Y%m%d')=DATE_FORMAT(CURDATE(),'%Y%m%d'),true,null)) as day_visit_count
// 		FROM admin_log `).Values(&maps)

// 	c.Data["adminCount"],_ = o.QueryTable(new(models.AdminUser)).Count() //管理员数
// 	c.Data["visitCount"],_ = maps[0]["sum_visit_count"] //总访问数
// 	c.Data["yestVisitCount"],_ = maps[0]["year_visit_count"] //年访问数
// 	c.Data["monthVisitCount"],_ = maps[0]["month_visit_count"] //月访问数
// 	c.Data["dayVisitCount"],_ = maps[0]["day_visit_count"] //日访问数

// 	c.Data["adminUser"] = c.GetSession("admin_user")
// 	c.Data["nowTime"] = utils.GetDateByTime(time.Now().Unix())
// 	//获取系统信息
// 	c.Data["appName"] = beego.BConfig.AppName  //项目名称
// 	c.Data["sysName"] = runtime.GOOS  //操作系统
// 	c.Data["goArch"] = runtime.GOARCH //系统构架 386、amd64
// 	c.Data["siteUrl"] = c.Ctx.Request.Host //服务器地址
// 	c.Data["sysVersion"] = runtime.Version() //go版本
// 	c.Data["mysqlVersion"] = models.MysqlVersion() //mysql版本号
// 	c.Data["frameVersion"] = beego.VERSION //框架版本
// 	disk := utils.DiskUsage("/")
// 	diskAll,_ := strconv.ParseFloat(fmt.Sprintf("%.2f",float64(disk.All) / 1024 / 1024 / 1024),64)
// 	diskFree,_ := strconv.ParseFloat(fmt.Sprintf("%.2f",float64(disk.Free) / 1024 / 1024 / 1024),64)
// 	c.Data["diskAll"] = diskAll //总容量
// 	c.Data["diskFree"] = diskFree //可用的容量
// 	c.Data["cpuNumber"] = runtime.NumCPU() //CPU数量
// 	c.Data["goRoot"] = runtime.GOROOT() //GO语言路径
// 	c.Data["getwd"],_ = os.Getwd() //GO语言路径

// 	c.SetTpl("index/welcome.html")
// }

// //个人中心
// func (c *IndexController) Person(){

// 	if c.IsAjax(){
// 		//修改个人中心
// 		var adminUser models.AdminUser
// 		_ = c.Ctx.Input.Bind(&adminUser.Id, "id")

// 		o := orm.NewOrm()
// 		err := o.Read(&adminUser)
// 		if err != nil{
// 			c.ApiError("查询失败",nil)
// 		}
// 		_ = c.Ctx.Input.Bind(&adminUser.HeadPortrait,"head_portrait")
// 		_ = c.Ctx.Input.Bind(&adminUser.Sex,"sex")
// 		_ = c.Ctx.Input.Bind(&adminUser.Phone,"phone")
// 		_ = c.Ctx.Input.Bind(&adminUser.Email,"email")
// 		pass := c.GetString("pass","")
// 		if pass != ""{
// 			repass := c.GetString("repass")
// 			valid := validation.Validation{}
// 			valid.MinSize(pass,6,"pass").Message("密码长度不能小于6位")
// 			valid.MaxSize(pass,20,"pass").Message("密码长度不能大于20位")
// 			if valid.HasErrors(){
// 				for _,err := range valid.Errors{
// 					c.ApiError(err.Message,nil)
// 				}
// 			}
// 			if pass != repass{
// 				c.ApiError("两次密码不一致",nil)
// 			}
// 			adminUser.Password = utils.Md5Encode(pass + beego.AppConfig.String("pwd_salt"))
// 		}
// 		_, err = o.Update(&adminUser)
// 		if err != nil{
// 			c.ApiError(err.Error(),nil)
// 		}
// 		//更新Session
// 		c.ApiSuccess("更新成功",nil)
// 	}
// 	adminUser := c.GetSession("admin_user")
// 	if adminUser == nil {
// 		//未登录
// 		c.RequestError("用户未登录","/admin/login")
// 	}
// 	c.Data["adminUser"] = adminUser
// 	c.SetTpl("index/person.html")
// }
