package model

import uuid "github.com/satori/go.uuid"

type User struct {
	ID       uuid.UUID `json:"user_id" gorm:"not null;unique;primaryKey;"`
	Name     string    `json:"user_name"`
	Account  string    `json:"account"`
	Password string    `json:"password"`
	Type     int       `json:"type" gorm:"default:3"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
}
