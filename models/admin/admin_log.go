package admin

import "time"



// func (a *AdminLog) TableName() string {
// 	return "admin_log"
// }

//AdminUser实体类
type AdminLog struct {
	Id          uint 	`gorm:"primaryKey;autoIncrement;column:f_id;";json:"pId"`
	AdminUserId uint32    `gorm:"column:f_userid;description:管理员ID";json:"pAdminUserId"`
	Url         string    `gorm:"type:varchar(1024);column:f_url;description:操作页面";json:"pUrl"`
	Title       string    `gorm:"type:varchar(100);column:f_title;description:日志标题";json:"pTitle"`
	Content     string    `gorm:"type:text;column:f_content;description:日志内容";json:"pContent"`
	Ip          string    `gorm:"type:varchar(100);column:f_ip;description:ip地址";json:"pIp"`
	Useragent   string    `gorm:"type:varchar(512);column:f_useragent;description:User-Agent";json:"pUseragent"`
	CreatedAt   time.Time `gorm:"index;auto_now_add;type:datetime;column:f_created_time;description:创建时间";json:"pCreatedAt"`
}

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
