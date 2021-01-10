# goadmin

#### 介绍
[goadmin] v1.0版本，基于gin框架和XAdmin前端框架，开发的go语言通用后台系统。
#### 软件架构
```
|--conf   配置相关  
|   |--app.yml          应用配置     
|   |--cache.yml        缓存配置  
|   |--db.yml           数据库配置
|   |--admin.yml        管理员信息（程序启动后自动初始化数据，数据存在则跳过）
|   |--log.yml          日志配置
|--controllers          控制器  
|   |--page             页面    
|   |--webapi           数据接口  
|--handle               相关处理函数  
|--lib                  类库  
|--models               数据模型及数据库访问  
|--routers              路由  
|--services             业务层 
|--static               静态文件 
|--view                 视图模板
|--main.go              主函数
```


#### 安装教程

##### 1、clone 项目到本地
```
GitHub:   git clone git@github.com:cn-joyconn/goadmin.git
```
或
```
码云:   git clone git@gitee.com:cn-joyconn/goadmin.git
```
##### 2、clone 项目到本地
```
go build main.go 将自动下载依赖包
```

#### 使用说明

1.  访问`/后台地址/index/index`，。
2.  超级管理员账号为conf/defuser.yml中superAdmin=true的账号
3.  默认超级管理员的密码为`super_admin`
4.  [日志配置](lib/log/readme.md)

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


## License ##
MIT License
