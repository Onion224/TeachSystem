package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/global"
	"server/middleware"
	"server/router/system"
)

// 初始化总路由
func Routers() *gin.Engine {
	Router := gin.Default()

	Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址

	global.GVA_LOG.Info("use middleware logger")

	// 跨域
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	global.GVA_LOG.Info("use middleware cors")

	//获取路由组实例
	manageRouter := system.RouterGroupApp.ManageRouter
	EffectRouter := system.RouterGroupApp.EffectRouter
	UserRouter := system.RouterGroupApp.UserRouter
	AuthorityRouter := system.RouterGroupApp.AuthorityRouter
	//下面的代码实现的是一个嵌套路由组
	PrivateGroup := Router.Group("")
	//这里我们不使用中间件
	//路由注册与初始化
	manageRouter.InitManageRouter(PrivateGroup)       //教学管理相关路由
	EffectRouter.InitEffectRouter(PrivateGroup)       //教学效果相关路由
	UserRouter.InitUserRouter(PrivateGroup)           //用户相关路由
	AuthorityRouter.InitAuthorityRouter(PrivateGroup) //管理员相关路由
	global.GVA_LOG.Info("router register success")

	return Router
}
