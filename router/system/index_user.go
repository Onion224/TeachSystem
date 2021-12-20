package system

import (
	"github.com/gin-gonic/gin"
)

//用户相关路由
type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	r := Router.Group("user")
	baserouter := RouterGroupApp.BaseRouter
	{
		r.POST("register", baserouter.register) //用户注册
		r.POST("login", baserouter.Login)       //用户登录
	}
}
