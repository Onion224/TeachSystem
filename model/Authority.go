package model

type Authority struct {
	AuthorityId     string      `json:"type" gorm:"not null;unique;primary_key"`
	ParentId        string      `json:"parent_id"`
	AuthorityName   string      `json:"authority_name"`
	DataAuthorityId []Authority `json:"data_authority" gorm:"many2many:sys_data_authority_id"`
	Children        []Authority `json:"children" gorm:"-"`
}
