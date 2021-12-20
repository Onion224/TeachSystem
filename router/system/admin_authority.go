package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model"
	"server/model/common/request"
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
	var authority model.Authority
	_ = c.ShouldBindJSON(&authority)
	if err, authBack := authorityService.CreateAuthority(authority); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(systemRes.AuthorityResponse{Authority: authBack}, "创建成功", c)
	}
}

func (s *AuthorityRouter) GetAuthorityList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)

	if err, list, total := authorityService.GetAuthorityInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败！", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}

}
