package global

// "errors"
// "net/http"
// "strings"

// handle "github.com/cn-joyconn/goadmin/handle"
// gin "github.com/gin-gonic/gin"

const (
	SUCCESS     = 1 //操作成功
	LoginSucess = 2 //登录成功

	ERROR          = 100101 //"操作失败"
	ParamsError    = 100102 //参数错误
	NoResult       = 100103 //没有结果
	NotFound       = 100104 //没有结果
	ServiceError   = 100105 //服务器错误
	NoRule         = 100106 //没有权限
	CehckCodeError = 100107 //验证码错误
	VerifyFail     = 100108 //参数签名验证失败
	DbError        = 100109 //数据访问错误
	FiledRepeat    = 100110 //字段重复

	NoLogin        = 100201 //用户未登录
	TokenNotExist  = 100202 //令牌不存在
	TokenFail      = 100203 //令牌错误
	UserLocck      = 100204 //用户锁定
	LoginIdError   = 100205 //登录id错误
	LoginPassError = 100206 //密码错误
	LoginFail      = 100207 //登录失败
	UserNameExisit = 100208 //用户名已存在
	EmailExisit    = 100209 //邮箱已存在
	PhoneExisit    = 100210 //手机号已存在

	CreditNotEnough = 100301 //余额不足
	GoodsInvalid    = 100302 //商品已失效
	PayError        = 100303 //支付失败

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
