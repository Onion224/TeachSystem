package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model"
	"server/model/common/response"
	systemReq "server/model/request"
)

//定义的结构体要在enter.go中注册
type ManageRouter struct{}

func (m *ManageRouter) InitManageRouter(Router *gin.RouterGroup) {
	//返回JSON数据
	router := Router.Group("manage")
	managerouter := RouterGroupApp.ManageRouter
	{
		router.GET("findCoursePlanById", managerouter.findCoursePlanById)
		router.GET("findCoursePlanList", managerouter.findCoursePlanList)
		router.GET("findClass", managerouter.findClass)
		router.GET("findRoster", managerouter.findRoster)
		router.GET("findGrade", managerouter.findGrade)
		router.GET("findHomework", managerouter.findHomework)
	}
}

// @Tags CoursesPlan
// @Summary 用id查询CoursesPlan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CoursesPlan true "ID"
func (m *ManageRouter) findCoursePlanById(c *gin.Context) {
	var coursePlan model.CoursesPlan
	_ = c.ShouldBindQuery(&coursePlan)
	if err, recoursesPlan := manageService.GetCoursePlan(coursePlan.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(gin.H{"coursePlan": recoursesPlan}, "查询成功", c)
	}
}

// @Summary 分页获取CoursesPlan列表
func (m *ManageRouter) findCoursePlanList(c *gin.Context) {
	var pageInfo systemReq.CoursePlanSearch
	_ = c.ShouldBindJSON(&pageInfo)
	if err, list, total := manageService.GetCoursesPlanlist(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

//查看班级花名册
func (m *ManageRouter) findRoster(c *gin.Context) {
	var pageInfo systemReq.RosterSearch
	_ = c.ShouldBindJSON(&pageInfo)
	if err, list, total := manageService.GetRosterlist(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
func (m *ManageRouter) findClass(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 200,
	})
}
func (m *ManageRouter) findGrade(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 200,
	})
}
func (m *ManageRouter) findHomework(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 200,
	})
}
