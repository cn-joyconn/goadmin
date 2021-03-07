package global

import (
	gologs "github.com/cn-joyconn/gologs"
	array "github.com/cn-joyconn/goutils/array"
	filetool "github.com/cn-joyconn/goutils/filetool"
	yaml "gopkg.in/yaml.v2"
)

var AppConf *AppConfig
var DBConf *DBConfig
var Admins *DefUsers
var appConfigPath string
var dbConfigPath string
var adminConfigPath string

//AppConfig 应用配置
type AppConfig struct {
	Name            string            `json:"name" yaml:"name"`                       // 应用名称
	WebPort         int               `json:"webport" yaml:"webport"`                 //web服务监听端口
	RunMode         string            `json:"runmode" yaml:"runmode"`                 //运行模式 dev prod test
	EnableGzip      bool              `json:"enablegzip" yaml:"enablegzip"`           //是否启用gzip
	ContextPath     string            `json:"contextpath" yaml:"contextpath"`         //虚拟路径
	JSPath          string            `json:"jspath" yaml:"jspath"`                   //js访问路径
	CSSPath         string            `json:"csspath" yaml:"csspath"`                 //css访问路径
	ImagePath       string            `json:"imagepath" yaml:"imagepath"`             //image访问路径
	FilePath        string            `json:"filepath" yaml:"filepath"`               //file访问路径
	Cache           map[string]string `json:"cache" yaml:"cache"`                     //缓存catalog 及 CacheName
	Authorize       AuthorizeCfg      `json:"authorize" yaml:"authorize"`             //登录认证相关配置
	SnowflakeWorkID int64             `json:"snowflakeWorkID" yaml:"snowflakeWorkID"` //全局唯一ID工作节点（雪花算法节点）
}
type appConfigs struct {
	App AppConfig `json:"app" yaml:"app"`
}

//DBConfig 数据库配置
type DBConfig struct {
	DBType                 string `json:"db_type" yaml:"db_type"`                                         // 数据库类型
	DBDtPrefix             string `json:"db_dt_prefix" yaml:"db_dt_prefix"`                               // 数据库表名前辍
	DBLog                  bool   `json:"db_log" yaml:"db_log"`                                           // 日志
	SkipDefaultTransaction bool   `json:"db_skip_default_transaction" yaml:"db_skip_default_transaction"` // 是否禁用事务（有助于提高性能）
	MaxIdleConns           int    `json:"db_max_idle_conns" yaml:"db_max_idle_conns"`                     // 设置空闲连接池中的最大连接数
	MaxOpenConns           int    `json:"db_max_open_conns" yaml:"db_max_open_conns"`                     // 数据库的最大打开连接数
	ConnMaxLifetime        int    `json:"db_conn_max_life_time" yaml:"db_conn_max_life_time"`             // 连接最长存活期，超过这个时间连接将不再被复用 单位秒
	SlowThreshold          int64  `json:"slow_threshold" yaml:"slow_threshold"`                           //慢查询阈值 0不记录   单位秒
	Postgres               DBlink `json:"postgres" yaml:"postgres"`
	Mysql                  DBlink `json:"mysql" yaml:"mysql"`
	Sqlite3                DBlink `json:"sqlite3" yaml:"sqlite3"`
}
type dbConfigs struct {
	Database DBConfig `json:"database" yaml:"database"`
}

//DBlink 数据库连接
type DBlink struct {
	DBName    string `json:"db_name" yaml:"db_name"`       // 连接数据库
	DBUser    string `json:"db_user" yaml:"db_user"`       // 连接用户
	DBPWD     string `json:"db_pwd" yaml:"db_pwd"`         // 连接密码
	DBHost    string `json:"db_host" yaml:"db_host"`       // 连接地址
	DBPort    int    `json:"db_port" yaml:"db_port"`       // 连接端口
	DBCharset string `json:"db_charset" yaml:"db_charset"` // 字符集
}

//DefUser 默认用户信息
type DefUser struct {
	ID         int    `json:"id" yaml:"id"`
	Alias      string `json:"alias" yaml:"alias"`
	SuperAdmin bool   `json:"superadmin" yaml:"superadmin"`
	UserName   string `json:"username" yaml:"username"`
	Phone      string `json:"phone" yaml:"phone"`
	Email      string `json:"email" yaml:"email"`
}
type DefUsers struct {
	Users []DefUser `json:"users" yaml:"users"`
}

//登录认证相关配置
type AuthorizeCfg struct {
	LoginUrl      string              `json:"loginUrl" yaml:"loginUrl"`           //登录页面的url
	LoginRefParam string              `json:"loginRefParam" yaml:"loginRefParam"` //跳转登录页面的携带源url的参数名
	Multilogin    bool                `json:"multilogin" yaml:"multilogin"`       //是否允许一个账号多人同时登录  是true 否false
	VerifyCode    LoginPageVerifyCode `json:"verifyCode" yaml:"verifyCode"`       //登录页验证码
	Cookie        AuthorizeCookie     `json:"cookie" yaml:"cookie"`
}

//登录页验证码
type LoginPageVerifyCode struct {
	Enable bool   `json:"enable" yaml:"enable"` //登录时是否启用验证码
	Method string `json:"method" yaml:"method"` //认证方式 1:数字,2:字母,3:算术,4:数字字母混合.
}

//登录认证cookie
type AuthorizeCookie struct {
	Domain           string `json:"domain" yaml:"domain"`                     //令牌作用的域名，设置为abc.com对a.abc.com、b.abc.com均有效，设置a.abc.com对b.abc.com无效。cookie的作用域参加w3c，如果设置0.0.0.0则代表默认域名
	LoginToken       string `json:"loginToken" yaml:"loginToken"`             //认证令牌在cookie中的名称
	LoginTokenAesKey string `json:"loginTokenAesKey" yaml:"loginTokenAesKey"` //认证令牌aes加密key
}

func InitAppConf(configPath string) {
	if filetool.IsExist(configPath) {
		configBytes, err := filetool.ReadFileToBytes(configPath)
		if err != nil {
			gologs.GetLogger("").Error("读取" + configPath + "文件错误。" + err.Error())
			return
		}
		gologs.GetLogger("").Info("成功读取" + configPath + "文件")
		var appconfigs appConfigs
		err = yaml.Unmarshal(configBytes, &appconfigs)
		if err != nil {
			AppConf = &AppConfig{Name: "joyconn-goadmin", WebPort: 8080, RunMode: "prod", EnableGzip: false, ContextPath: ""}
			gologs.GetLogger("").Error("解析" + configPath + "文件失败")
			return
		}
		gologs.GetLogger("").Info("成功解析" + configPath + "文件")
		appConfigPath = configPath
		AppConf = &appconfigs.App
	} else {
		gologs.GetLogger("").Error("未找到" + configPath)
		return
	}
}
func InitDBConf(configPath string) {
	if filetool.IsExist(configPath) {
		configBytes, err := filetool.ReadFileToBytes(configPath)
		if err != nil {
			gologs.GetLogger("").Error("读取" + configPath + "文件错误。" + err.Error())
			return
		}
		gologs.GetLogger("").Info("成功读取" + configPath + "文件")
		var dbconfs dbConfigs
		err = yaml.Unmarshal(configBytes, &dbconfs)
		if err != nil {
			DBConf = &DBConfig{
				DBType:                 "sqlite3",
				DBDtPrefix:             "t_",
				DBLog:                  true,
				SkipDefaultTransaction: false,
				MaxIdleConns:           10,
				MaxOpenConns:           30,
				ConnMaxLifetime:        600,
			}
			// DBConf.DBlinks=make(map[string]*DBlink)
			DBConf.Sqlite3 = DBlink{DBName: ".sqlite3.db"}
			gologs.GetLogger("").Error("解析" + configPath + "文件失败")
			return
		}
		gologs.GetLogger("").Info("成功解析" + configPath + "文件")
		dbConfigPath = configPath
		DBConf = &dbconfs.Database
	} else {
		gologs.GetLogger("").Error("未找到" + configPath)
		return
	}
}

func LoadAdmin(configPath string) {
	if filetool.IsExist(configPath) {
		configBytes, err := filetool.ReadFileToBytes(configPath)
		if err != nil {
			gologs.GetLogger("").Error("读取" + configPath + "文件错误。" + err.Error())
			return
		}
		gologs.GetLogger("").Info("成功读取" + configPath + "文件")
		err = yaml.Unmarshal(configBytes, &Admins)
		if err != nil {
			gologs.GetLogger("").Error("解析" + configPath + "文件失败")
			return
		}
		gologs.GetLogger("").Info("成功解析" + configPath + "文件")
		users := make([]int, 0)
		for _, val := range Admins.Users {
			if val.SuperAdmin && val.ID > 0 {
				users = append(users, val.ID)
			}
		}
		adminConfigPath = configPath
		SetSuperAdminUsers(users)
	} else {
		gologs.GetLogger("").Error("未找到" + configPath)
		return
	}
}

//SaveAdmin 保存管理员信息
func SaveAdminConfig() {
	saveConfig(&Admins, adminConfigPath)
}
func SaveDBConfig() {
	dbconfs := &dbConfigs{Database: *DBConf}
	saveConfig(dbconfs, dbConfigPath)
}
func SaveAppConfig() {
	appconfs := &appConfigs{App: *AppConf}
	saveConfig(appconfs, appConfigPath)
}
func saveConfig(in interface{}, configPath string) {
	configBytes, err := yaml.Marshal(in)
	if err != nil {
		gologs.GetLogger("").Error(err.Error())
		return
	}
	_, err = filetool.WriteBytesToFile(configPath, configBytes)
	if err != nil {
		gologs.GetLogger("").Error(err.Error())
		return
	}
}

var adminUsers []int

//SetSuperAdminUsers 设置超级管理员账号
func SetSuperAdminUsers(users []int) {
	adminUsers = make([]int, len(users))
	copy(users, adminUsers) //users 复制给 adminUsers
}

//GetSuperAdminUsers 获取所有超级管理员账号
func GetSuperAdminUsers() []int {
	var users = make([]int, len(adminUsers))
	copy(users, adminUsers) //adminUsers复制给users
	return users
}

//IsSuperAdmin 是否为超级管理员
func IsSuperAdmin(user int) bool {
	return array.InIntArray(user, adminUsers)
}
