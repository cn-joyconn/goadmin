package auth
import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"

	"github.com/cn-joyconn/goadmin/models/global"
	"github.com/cn-joyconn/goutils/strtool"
	"github.com/gin-gonic/gin"
	adminServices "github.com/cn-joyconn/goadmin/services/admin"
	controllers "github.com/cn-joyconn/goadmin/controllers"
)



type ValidationResultModel struct {
	Pass int
	Msg  string
}

/**
* 处理验证结果
* @param isAjax
* @param pass
* @param request
* @param response
* @return
 */
func handelPass(isAjax bool, validationResultModel *ValidationResultModel, c *gin.Context) bool {
	bc := &controllers.BaseController{}
	if validationResultModel.Pass == 2 {
		return true
	} else if validationResultModel.Pass == 1 {
		//权限不足
		if isAjax {
			bc.ApiErrorCode(c, validationResultModel.Msg, "", global.NoRule)
		} else {
			c.Request.Response.StatusCode = 403
		}
		return false
	} else {
		//令牌已失效
		global.TokenHelper.DelMyAuthenticationToken(c)
		if isAjax {
			bc.ApiErrorCode(c, validationResultModel.Msg, "", global.NoLogin)
		} else {
			redirectURL := global.AppConf.ContextPath + "/" + global.AppConf.Authorize.LoginUrl
			redirectURL = strings.Replace(redirectURL, "//", "/", -1)
			if strings.Index(redirectURL, "?") < 0 {
				redirectURL += "?"
			}
			redirectURL += "&"
			redirectURL += global.AppConf.Authorize.LoginRefParam
			redirectURL += "="
			redirectURL += url.QueryEscape(c.Request.RequestURI)
			c.Redirect(302, redirectURL)
		}
		return false
	}
}

/**
* 扩展类 自定义取token
* @param request
* @param response
* @param resources
* @return 0:验证失败  1::权限不足  2:验证通过
 */
func getExtPass(c *gin.Context, resources []string, needPermission bool) int {
	return 0
}

/**
* 验证cookie 中的 token
* @param request
* @param response
* @param resources
* @return 0:验证失败  1::权限不足  2:验证通过
 */
func getAuthPass(c *gin.Context, resources []string, needPermission bool) *ValidationResultModel {
	pass := 0
	ltID,byCookie := global.TokenHelper.GetMyAuthToken(c)
	// , byCookie bool
	if ltID != nil && !strtool.IsBlank(ltID.Sign) {
		guSign := global.TokenHelper.ValidationSign(ltID.Uid, ltID.Sign)
		if guSign == global.SUCCESS {
			//缓存中存在该用户的sign
			pass = 1
		} else if guSign == global.TokenNotExist {
			//缓存中不存在该用户的sign,需要做模拟登录
			adminUserService := &adminServices.AdminUserService{}
			_, loginCode := adminUserService.LoginByUserIDAndEncryptPwd(ltID.Uid, ltID.Pwd)
			if loginCode == global.LoginSucess {
				global.TokenHelper.SetAuthenticationToken(ltID.Uid, ltID.Pwd, c, byCookie)
				pass = 1
			}
		}

	} else {
		pass = getExtPass(c, resources, needPermission)
	}
	//用户处于登录状态，判断是否有对当前路径访问的权限
	if pass == 1 {
		if needPermission {
			_,err := strconv.Atoi(ltID.Uid)
			if err==nil{
				adminUserPermissionService := &adminServices.AdminUserPermissionService{}
				if adminUserPermissionService.PathPermissin(ltID.Uid, resources) {
					pass = 2
				}
			}			
		} else {
			pass = 2
		}

	}
	validationResultModel := &ValidationResultModel{
		Pass: pass,
		Msg:  "",
	}
	if pass==1{
		adminResourceService := &adminServices.AdminResourceService{}
		resourceObjs := adminResourceService.GetPermissinsNames(resources)
		if resourceObjs != nil {
			msgBytes,err := json.Marshal(resourceObjs)
			 if err==nil{
				validationResultModel.Msg = string(msgBytes)
			 }
		}
	}
	
	return validationResultModel
}


