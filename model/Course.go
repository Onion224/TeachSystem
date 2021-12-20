package model

import "github.com/shopspring/decimal"

type Course struct {
	ID     string          `json:"course_id" gorm:"not null;unique;primaryKey;"`
	Name   string          `json:"name"`
	Credit decimal.Decimal `json:"credit"`
	Period decimal.Decimal `json:"period"`
	//Teacher和Course的has many关系
	TeacherID uint `json:"teacher_id"`
	//Course和Score的has many关系
	Scores []Score `json:"scores" gorm:"foreignKey:CourseID"`
}
