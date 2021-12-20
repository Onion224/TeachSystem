package model

import "time"

type Student struct {
	ID          string    `json:"student_id" gorm:"not null;unique;primary_key;comment:角色ID;size:90"`
	Name        string    `json:"student_name"`
	Sex         int       `json:"sex"`
	Birthday    time.Time `json:"birthday"`
	NativePlace string    `json:"native_place"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	//Student和Score的has many关系
	Scores []Score `json:"scores" gorm:"foreignKey:StudentID"`
	//Student和Course的many2many关系
	Courses []Course `json:"courses" gorm:"many2many:student_courses;"`
	//Student和Teacher的many2many关系
	Teachers []Teacher `json:"teachers" gorm:"many2many:student_teachers"`
	//Class和Student的has many外键
	ClassID string `json:"class_id"`
}
