package global

// "errors"
// "net/http"
// "strings"

// handle "github.com/cn-joyconn/goadmin/handle"
// gin "github.com/gin-gonic/gin"

const (
	SUCCESS     = 1 //操作成功
	LoginSucess = 2 //登录成功

	ERROR          = -1 //"操作失败"
	ParamsError    = -2 //参数错误
	ServiceError   = -4 //服务器错误
	NoResult       = -3 //没有结果
	NoRule         = -5 //没有权限
	CehckCodeError = -6 //验证码错误
	VerifyFail     = -7 //参数签名验证失败
	DbError        = -8 //数据访问错误
	FiledRepeat    = -9 //字段重复

	NoLogin        = -101 //用户未登录
	TokenNotExist  = -102 //令牌不存在
	TokenFail      = -103 //令牌错误
	UserLocck      = -104 //用户锁定
	LoginIdError   = -105 //登录id错误
	LoginPassError = -106 //密码错误
	LoginFail      = -107 //登录失败
	UserNameExisit = -108 //用户名已存在
	EmailExisit    = -109 //邮箱已存在
	PhoneExisit    = -110 //手机号已存在

	CreditNotEnough = -201 //余额不足
	GoodsInvalid    = -202 //商品已失效
	PayError        = -203 //支付失败

)

//响应参数结构体
type Response struct {
	Code     int         `json:"code"`
	Msg      string      `json:"msg"`
	Data     interface{} `json:"data"`
	Url      string      `json:"url"`
	Wait     int         `json:"wait"`
	AllCount int64       `json:"allcount"`
}

// //返回结果辅助函数
// func Result(code int, msg string, data interface{}, url string, wait int, header map[string]string, ctx *gin.Context) {
// 	if strings.ToLower(ctx.Request.Method) == "post" {
// 		result := Response{
// 			Code: code,
// 			Msg:  msg,
// 			Data: data,
// 			Url:  url,
// 			Wait: wait,
// 		}

// 		if len(header) > 0 {
// 			for k, v := range header {
// 				ctx.Header(k, v)
// 			}
// 		}

// 		ctx.JSON(200, result)

// 		//Controller中this.StopRun()用法
// 		panic(errors.New("user stop run"))
// 	}

// 	if url == "" {
// 		url = ctx.Request.Referer()
// 		if url == "" {
// 			url = "/admin/index/index"
// 		}
// 	}

// 	ctx.Redirect(http.StatusFound, url)
// }

// //Success 成功、普通返回
// func Success(ctx *gin.Context) {
// 	Result(SUCCESS, "操作成功", "", URL_BACK, 0, map[string]string{}, ctx)
// }

// //SuccessWithMessage 成功、返回自定义信息
// func SuccessWithMessage(msg string, ctx *gin.Context) {
// 	Result(SUCCESS, msg, "", URL_BACK, 0, map[string]string{}, ctx)
// }

// //SuccessWithMessageAndURL 成功、返回自定义信息和url
// func SuccessWithMessageAndURL(msg string, url string, ctx *gin.Context) {
// 	Result(SUCCESS, msg, "", url, 0, map[string]string{}, ctx)
// }

// //SuccessWithDetailed 成功、返回所有自定义信息
// func SuccessWithDetailed(msg string, url string, data interface{}, wait int, header map[string]string, ctx *gin.Context) {
// 	Result(SUCCESS, msg, data, url, wait, header, ctx)
// }

// //Error 失败、普通返回
// func Error(ctx *gin.Context) {
// 	Result(ERROR, "操作失败", "", URL_CURRENT, 0, map[string]string{}, ctx)
// }

// //ErrorWithMessage 失败、返回自定义信息
// func ErrorWithMessage(msg string, ctx *gin.Context) {
// 	Result(ERROR, msg, "", URL_CURRENT, 0, map[string]string{}, ctx)
// }

// //ErrorWithMessageAndURL 失败、返回自定义信息和url
// func ErrorWithMessageAndURL(msg string, url string, ctx *gin.Context) {
// 	Result(ERROR, msg, "", url, 0, map[string]string{}, ctx)
// }

// //ErrorWithDetailed 失败、返回所有自定义信息
// func ErrorWithDetailed(msg string, url string, data interface{}, wait int, header map[string]string, ctx *gin.Context) {
// 	Result(ERROR, msg, data, url, wait, header, ctx)
// }
