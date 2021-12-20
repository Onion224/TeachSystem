package model

type Class struct {
	ID        string `json:"class_id" gorm:"not null;unique;primaryKey;"`
	Name      string `json:"class_name"`
	Grade     string `json:"grade"`
	Specialty string `json:"specialty"`
	//has many关系
	Students []Student `json:"students"  gorm:"foreignKey:ClassID"`
	//has many关系
	Scores []Score `json:"scores" gorm:"foreignKey:ClassID"`
}
