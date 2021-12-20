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
	ScoreID string          `json:"score_id" gorm:"not null;unique;primaryKey;"`
	Score   decimal.Decimal `json:"score"`
	//Student和Score的has many外键
	StudentID string `json:"student_id"`
	//Course和Score的has many外键
	CourseID string `json:"course_id"`
	//Class和Socre的has many外键
	ClassID string `json:"class_id"`
	request.PageInfo
}
