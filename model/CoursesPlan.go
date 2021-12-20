package model

type CoursesPlan struct {
	ID        string  `json:"id"`
	CourseID  string  `json:"course_id"`
	TeacherID string  `json:"teacher_id"`
	ClassID   string  `json:"class_id"`
	Class     Class   `json:"class" gorm:"foreignKey:ClassID"`
	Course    Course  `json:"course" gorm:"foreignKey:CourseID"`
	Teacher   Teacher `json:"teacher" gorm:"foreignKey:TeacherID"`
}
