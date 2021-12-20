package model

import "github.com/shopspring/decimal"

type Score struct {
	ScoreID string          `json:"score_id" gorm:"not null;unique;primaryKey;"`
	Score   decimal.Decimal `json:"score"`
	//Student和Score的has many外键
	StudentID string `json:"student_id"`
	//Course和Score的has many外键
	CourseID string `json:"course_id"`
	//Class和Socre的has many外键
	ClassID string `json:"class_id"`
}
