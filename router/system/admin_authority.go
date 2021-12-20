package system

import "github.com/gin-gonic/gin"

type AuthorityRouter struct{}

func (s *AuthorityRouter) InitAuthorityRouter(Router *gin.RouterGroup) {
	//返回JSON数据
	router := Router.Group("admin")
	adminrouter := RouterGroupApp.AuthorityRouter
	{
		router.POST("createAuthority", adminrouter.createAuthority)
	}
}

func (s *AuthorityRouter) createAuthority(c *gin.Context) {

}
