package admin

// import (
// 	"time"

// 	"github.com/beego/beego/v2/client/orm"
// 	"github.com/beego/beego/v2/server/web/context"
// )

// func (a *AdminLog) TableName() string {
// 	return "admin_log"
// }

// //AdminUser实体类
// type AdminLog struct {
// 	Id          uint 	`orm:"auto;column(f_id);"`
// 	AdminUserId uint32    `orm:"column(f_userid);"description:"管理员ID"`
// 	Username    string    `orm:"size(60);column(f_username);"description:"管理员名称"`
// 	Url         string    `orm:"size(1024);column(f_url);"description:"操作页面"`
// 	Title       string    `orm:"size(100);column(f_title);"description:"日志标题"`
// 	Content     string    `orm:"type(text);column(f_content);"description:"日志内容"`
// 	Ip          string    `orm:"size(100);column(f_ip);"description:"ip地址"`
// 	Useragent   string    `orm:"size(512);column(f_useragent);" description:"User-Agent"`
// 	CreatedAt   time.Time `orm:"auto_now_add;type(datetime);column(f_created_time);"description:"创建时间"`
// }

// //添加用户日志
// func AdminLogAddLog(ctx *context.Context, title string) {
// 	adminUserSession := ctx.Input.Session("admin_user")
// 	var adminUserId uint32
// 	var userName string
// 	if adminUserSession != nil {
// 		adminUser := adminUserSession.(*AdminUser)
// 		adminUserId = adminUser.ID
// 		userName = adminUser.Username
// 	}
// 	adminLog := AdminLog{
// 		AdminUserId: adminUserId,
// 		Username:    userName,
// 		Url:         ctx.Request.RequestURI,
// 		Title:       title, //请求标题
// 		Content:     "",    //请求数据D
// 		Ip:          ctx.Request.RemoteAddr,
// 		Useragent:   ctx.Request.UserAgent(),
// 	}
// 	m := orm.NewOrm()
// 	_, _ = m.Insert(&adminLog)
// }
