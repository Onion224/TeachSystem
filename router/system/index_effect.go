package system

import "github.com/gin-gonic/gin"

type EffectRouter struct{}

func (s *EffectRouter) InitEffectRouter(Router *gin.RouterGroup) {
	//返回JSON数据
	router := Router.Group("effect")
	effectrouter := RouterGroupApp.EffectRouter
	{
		router.GET("expert", effectrouter.expert)
		router.GET("supervise", effectrouter.supervise)
		router.GET("self_esteem", effectrouter.selfesteem)
		router.GET("student", effectrouter.student)
		router.GET("social", effectrouter.social)
		router.GET("access", effectrouter.access)
	}
}

func (s *EffectRouter) expert(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 200,
	})
}
func (s *EffectRouter) supervise(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 200,
	})
}
func (s *EffectRouter) selfesteem(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 200,
	})
}
func (s *EffectRouter) student(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 200,
	})
}
func (s *EffectRouter) social(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 200,
	})
}
func (s *EffectRouter) access(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 200,
	})
}
