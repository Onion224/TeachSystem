package system

import (
	"server/global"
	"server/model"
	systemReq "server/model/request"
)

type ManageService struct{}

func (m *ManageService) GetCoursePlan(id string) (err error, coursePlan model.CoursesPlan) {
	return
}

// 获取开课计划列表
func (m *ManageService) GetCoursesPlanlist(info systemReq.CoursePlanSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.CoursesPlan{})
	var coursePlans []model.CoursesPlan
	//如果有条件搜索 下方会自动创建搜索语句
	if info.CourseID != "" {
		db = db.Where("courseID = ?", info.CourseID)
	}
	if info.TeacherID != "" {
		db = db.Where("teacherID = ?", info.TeacherID)
	}
	if info.ClassID != "" {
		db = db.Where("classID = ?", info.ClassID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	//db.Order:指定从数据库检索记录时的排序方式
	err = db.Order("id desc").Limit(limit).Offset(offset).Preload("Class").Preload("Course").Preload("Teacher").Find(&coursePlans).Error
	return err, coursePlans, total
}

//获取班级花名册列表
func (m *ManageService) GetRosterlist(info systemReq.RosterSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.Student{})
	var rosters []model.Student
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&rosters).Error
	return err, rosters, total
}

//查看班级成绩单
func (m *ManageService) GetScorelist(info systemReq.ScoreSearch) (err error, list interface{}, total int64) {
	return
}
