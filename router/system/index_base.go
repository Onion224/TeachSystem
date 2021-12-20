package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model"
	"server/model/common/response"
	systemReq "server/model/request"
	systemRes "server/model/response"
)

type BaseRouter struct{}

func (b *BaseRouter) register(c *gin.Context) {
	var register systemReq.Register
	//将request的body中的数据，自动按照json格式解析到结构体
	_ = c.ShouldBindJSON(&register)
	//if err := utils.Verify(register, utils.RegisterVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	//注册user对象
	user := &model.User{Account: register.Account, Password: register.Password}
	err, userReturn := userService.Register(*user)
	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册成功", c)
	}
}

func (b *BaseRouter) Login(c *gin.Context) {
	var l systemReq.Login
	_ = c.ShouldBindJSON(&l)
	u := &model.User{Account: l.Account, Password: l.Password}
	if err, user := userService.Login(u); err != nil {
		global.GVA_LOG.Error("登录失败！用户名不存在或者密码错误", zap.Error(err))
		response.FailWithMessage("用户名不存在或者密码错误", c)
	} else {
		response.OkWithDetailed(systemRes.LoginResponse{
			User: *user,
		}, "登录成功", c)
	}
}
