package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model"
	"server/model/common/response"
	systemRes "server/model/response"
)

type AuthorityRouter struct{}

func (s *AuthorityRouter) InitAuthorityRouter(Router *gin.RouterGroup) {
	//返回JSON数据
	router := Router.Group("admin")
	adminrouter := RouterGroupApp.AuthorityRouter
	{
		router.POST("createAuthority", adminrouter.createAuthority)   // 创建角色
		router.POST("getAuthorityList", adminrouter.GetAuthorityList) //获取角色列表
	}
}

func (s *AuthorityRouter) createAuthority(c *gin.Context) {
	var authority model.User
	_ = c.ShouldBindJSON(&authority)
	if err, authBack := authorityService.CreateAuthority(authority); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(systemRes.AuthorityResponse{Authority: authBack}, "创建成功", c)
	}
}
