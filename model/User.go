package model

import uuid "github.com/satori/go.uuid"

type User struct {
	ID          uuid.UUID   `json:"user_id" gorm:"not null;unique;primaryKey;"`
	Name        string      `json:"user_name"`
	Account     string      `json:"account"`
	Password    string      `json:"password"`
	AuthorityId string      `json:"authority_id" gorm:"default:888"`
	Email       string      `json:"email"`
	Phone       string      `json:"phone"`
	Authority   Authority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Authorities []Authority `json:"authorities" gorm:"many2many:sys_user_authority;"`
}
