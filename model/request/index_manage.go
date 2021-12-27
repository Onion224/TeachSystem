package request

import (
	"github.com/shopspring/decimal"
	"server/model"
	"server/model/common/request"
	"time"
)

type CoursePlanSearch struct {
	model.CoursesPlan
	request.PageInfo
}

type RosterSearch struct {
	ID          string    `json:"student_id" gorm:"not null;unique;primary_key;comment:角色ID;size:90"`
	Name        string    `json:"student_name"`
	Sex         int       `json:"sex"`
	Birthday    time.Time `json:"birthday"`
	NativePlace string    `json:"native_place"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	request.PageInfo
}

type ScoreSearch struct {
	ScoreID string `json:"score_id" gorm:"not null;unique;primaryKey;column:ScoreID"`
	//学生姓名
	StudentName string `json:"student_name" gorm:"column:StudentName"`
	//Course和Score的has many外键
	CourseName string `json:"course_name" gorm:"column:CourseName"`
	//Class和Socre的has many外键
	ClassName string `json:"class_name" gorm:"column:ClassName"`
	//分数
	Score decimal.Decimal `json:"score" gorm:"column:Score"`
	request.PageInfo
}
