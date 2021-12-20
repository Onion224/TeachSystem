package model

import "time"

type Teacher struct {
	ID                string    `json:"teacher_id" gorm:"not null;unique;primaryKey;"`
	Name              string    `json:"teacher_name"`
	Account           string    `json:"account"`
	Password          string    `json:"password"`
	Sex               int       `json:"sex"`
	Birthday          time.Time `json:"birthday"`
	ProfessionalTitle string    `json:"professional_title"`
	Email             string    `json:"email"`
	Phone             string    `json:"phone"`
	//Teacher和Course的has many关系
	Courses []Course `json:"courses" gorm:"foreignKey:TeacherID"`
}
